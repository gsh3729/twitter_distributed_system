package controllers

import (
	context "context"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"log"
	"net/http"

	"backend/connect"

	"google.golang.org/grpc"
)

func FollowPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")

		connectTo := c.PostForm("connectTo")

		var conn *grpc.ClientConn
		conn, err2 := grpc.Dial(":9000", grpc.WithInsecure())
		if err2 != nil {
			log.Fatalf("Couldn't connect: %s", err2)
		}

		auth_server := connect.NewConnectServiceClient(conn)
		response, err := auth_server.Follow(context.Background(), &connect.FollowRequest{
			User1: user.(string),
			User2: connectTo,
		})

		if err != nil {
			log.Fatalf("Error when calling Follow: %s", err)
		}

		if response.Success {
			c.Redirect(http.StatusMovedPermanently, "/dashboard")
		} else {
			c.HTML(http.StatusInternalServerError, "index.html", gin.H{"content": "Something went wrong, try again later."})
		}
	}
}

func UnfollowPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")

		unfollowPerson := c.PostForm("unfollowPerson")

		var conn *grpc.ClientConn
		conn, err2 := grpc.Dial(":9000", grpc.WithInsecure())
		if err2 != nil {
			log.Fatalf("Couldn't connect: %s", err2)
		}

		auth_server := connect.NewConnectServiceClient(conn)
		response, err := auth_server.Follow(context.Background(), &connect.FollowRequest{
			User1: user.(string),
			User2: unfollowPerson,
		})

		if err != nil {
			log.Fatalf("Error when calling Follow: %s", err)
		}

		if response.Success {
			c.Redirect(http.StatusMovedPermanently, "/dashboard")
		} else {
			c.HTML(http.StatusInternalServerError, "index.html", gin.H{"content": "Something went wrong, try again later."})
		}
	}
}
