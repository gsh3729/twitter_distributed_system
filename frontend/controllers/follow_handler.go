package controllers

import (
	context "context"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"net/http"

	"backend/follow"
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

		follow_server := follow.NewFollowServiceClient(conn)
		response, err := follow_server.Follow(context.Background(), &follow.FollowRequest{
			User1: user.(string),
			User2: connectTo,
		})

		if err != nil {
			log.Fatalf("Error when calling Follow: %s", err)
		}

		if response.Success {
			c.Redirect(http.StatusMovedPermanently, "/dashboard")
		} else {
			c.HTML(http.StatusInternalServerError, "dashboard.html", gin.H{"content": "Something went wrong, try again later."})
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

		follow_server := follow.NewFollowServiceClient(conn)
		response, err := follow_server.Unfollow(context.Background(), &follow.UnfollowRequest{
			User1: user.(string),
			User2: unfollowPerson,
		})

		if err != nil {
			log.Fatalf("Error when calling Unfollow: %s", err)
		}

		if response.Success {
			c.Redirect(http.StatusMovedPermanently, "/dashboard")
		} else {
			c.HTML(http.StatusInternalServerError, "dashboard.html", gin.H{"content": "Something went wrong, try again later."})
		}
	}
}

func FollowersGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")

		var conn *grpc.ClientConn
		conn, err2 := grpc.Dial(":9000", grpc.WithInsecure())
		if err2 != nil {
			log.Fatalf("Couldn't connect: %s", err2)
		}

		follow_server := follow.NewFollowServiceClient(conn)
		response, err := follow_server.GetUserFollowers(context.Background(), &follow.GetFollowersRequest{
			Username: user.(string),
		})

		if err != nil {
			log.Fatalf("Error when calling GetFollowers: %s", err)
		}

		if response.Success {
			c.HTML(http.StatusOK, "followers.html", gin.H{
				"content": response.Users,
				"user":    user,
			})
		} else {
			c.HTML(http.StatusInternalServerError, "followers.html", gin.H{"content": "Something went wrong, try again later."})
		}
	}
}

func FollowingGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")

		var conn *grpc.ClientConn
		conn, err2 := grpc.Dial(":9000", grpc.WithInsecure())
		if err2 != nil {
			log.Fatalf("Couldn't connect: %s", err2)
		}

		follow_server := follow.NewFollowServiceClient(conn)
		response, err := follow_server.GetUserFollowing(context.Background(), &follow.GetFollowingRequest{
			Username: user.(string),
		})

		if err != nil {
			log.Fatalf("Error when calling GetFollowing: %s", err)
		}

		if response.Success {
			c.HTML(http.StatusOK, "following.html", gin.H{
				"content": response.Users,
				"user":    user,
			})
		} else {
			c.HTML(http.StatusInternalServerError, "following.html", gin.H{"content": "Something went wrong, try again later."})
		}
	}
}
