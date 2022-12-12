package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"log"
	"net/http"

	auth "backend/auth"
	helpers "frontend/helpers"
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

		auth.SignUp(username, password)

		c.HTML(http.StatusCreated, "index.html", gin.H{"content": "Created user successfully"})
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

		auth.SignIn(username, password)

		session.Set("user", username)
		if err := session.Save(); err != nil {
			c.HTML(http.StatusInternalServerError, "login.html", gin.H{"content": "Failed to save session"})
			return
		}

		c.Redirect(http.StatusMovedPermanently, "/dashboard")
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
		session.Delete("user")
		if err := session.Save(); err != nil {
			log.Println("Failed to save session:", err)
			return
		}

		c.Redirect(http.StatusMovedPermanently, "/")
	}
}
