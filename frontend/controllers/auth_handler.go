package controllers

import (
	context "context"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"log"
	"net/http"

	helpers "frontend/helpers"

	"backend/authbackend"

	"google.golang.org/grpc"
)

func SignupGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")
		if user != nil {
			c.Redirect(http.StatusAccepted, "/dashboard")
			return
		}
		c.HTML(http.StatusOK, "signup.html", gin.H{})
	}
}

func LoginGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")
		if user != nil {
			c.Redirect(http.StatusAccepted, "/dashboard")
			return
		}
		c.HTML(http.StatusOK, "login.html", gin.H{})
	}
}

func SignupPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")

		if user != nil {
			c.Redirect(http.StatusAccepted, "/dashboard")
			return
		}

		username := c.PostForm("username")
		password := c.PostForm("password")

		if helpers.EmptyUserPass(username, password) {
			c.HTML(http.StatusBadRequest, "signup.html", gin.H{"content": "Parameters can't be empty"})
			return
		}

		var conn *grpc.ClientConn
		conn, err2 := grpc.Dial(":9000", grpc.WithInsecure())
		if err2 != nil {
			log.Fatalf("Couldn't connect: %s", err2)
		}
		auth_server := authbackend.NewAuthServiceClient(conn)
		response, err := auth_server.SignUp(context.Background(), &authbackend.UserSignUpRequest{
			Username: username,
			Password: password,
		})

		if err != nil {
			log.Fatalf("Error when calling UserSignUp: %s", err)
		}

		if response.Success {
			c.HTML(http.StatusCreated, "index.html", gin.H{"content": "Created user successfully"})
		} else {
			c.HTML(http.StatusInternalServerError, "index.html", gin.H{"content": "User exists."})
		}
	}
}

func LoginPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")

		if user != nil {
			c.Redirect(http.StatusAccepted, "/dashboard")
			return
		}

		username := c.PostForm("username")
		password := c.PostForm("password")

		if helpers.EmptyUserPass(username, password) {
			c.HTML(http.StatusBadRequest, "login.html", gin.H{"content": "Parameters can't be empty"})
			return
		}

		var conn *grpc.ClientConn
		conn, err2 := grpc.Dial(":9000", grpc.WithInsecure())
		if err2 != nil {
			log.Fatalf("Couldn't connect: %s", err2)
		}
		auth_server := authbackend.NewAuthServiceClient(conn)
		response, err := auth_server.SignIn(context.Background(), &authbackend.UserSignInRequest{
			Username: username,
			Password: password,
		})

		if err != nil {
			log.Fatalf("Error when calling UserSignIn: %s", err)
		}

		if response.Success {
			session.Set("user", username)
			if err := session.Save(); err != nil {
				c.HTML(http.StatusInternalServerError, "login.html", gin.H{"content": "Failed to save session"})
				return
			}
			c.Redirect(http.StatusMovedPermanently, "/dashboard")
		} else {
			c.HTML(http.StatusCreated, "index.html", gin.H{"content": "Invalid Username or password. Please check again."})
		}
	}
}

func LogoutGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")

		if user == nil {
			log.Println("Invalid session token")
			return
		}

		session.Set("user", "")
		session.Clear()
		session.Options(sessions.Options{Path: "/", MaxAge: -1})
		session.Save()

		c.Redirect(http.StatusTemporaryRedirect, "/")
	}
}
