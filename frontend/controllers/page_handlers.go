package controllers

import (
	context "context"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"net/http"

	"backend/follow"
	"backend/tweet"
)

type TweetStruct struct {
	Time string
	Text string
	User string
}

func IndexGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"user": user,
		})
	}
}

func DashboardGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")

		var conn *grpc.ClientConn
		conn, err2 := grpc.Dial(":9000", grpc.WithInsecure())
		if err2 != nil {
			log.Fatalf("Couldn't connect: %s", err2)
		}

		tweet_server := tweet.NewTweetServiceClient(conn)
		response, err := tweet_server.GetTweets(context.Background(), &tweet.GetTweetsRequest{
			Username: user.(string),
		})

		if err != nil {
			log.Fatalf("Error when calling Follow: %s", err)
		}

		if response.Success {
			var result []TweetStruct
			for i := 0; i < len(response.Text); i++ {
				tweet := TweetStruct{
					Text: response.Text[i],
					User: response.User[i],
					Time: response.Time[i],
				}
				result = append(result, tweet)
			}
			c.HTML(http.StatusOK, "dashboard.html", gin.H{
				"content": result,
				"user":    user,
			})
		} else {
			c.HTML(http.StatusInternalServerError, "dashboard.html", gin.H{"content": "Something went wrong, try again later."})
		}
	}
}

func TweetGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")

		c.HTML(http.StatusOK, "composeTweet.html", gin.H{
			"user": user,
		})
	}
}

func TweetPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")

		tweetMsg := c.PostForm("tweetMsg")

		var conn *grpc.ClientConn
		conn, err2 := grpc.Dial(":9000", grpc.WithInsecure())
		if err2 != nil {
			log.Fatalf("Couldn't connect: %s", err2)
		}

		tweet_server := tweet.NewTweetServiceClient(conn)
		response, err := tweet_server.PostTweet(context.Background(), &tweet.PostTweetRequest{
			Username: user.(string),
			Text:     tweetMsg,
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

func FindPeopleGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")

		var conn *grpc.ClientConn
		conn, err2 := grpc.Dial(":9000", grpc.WithInsecure())
		if err2 != nil {
			log.Fatalf("Couldn't connect: %s", err2)
		}

		follow_server := follow.NewFollowServiceClient(conn)
		response, err := follow_server.GetUsers(context.Background(), &follow.GetUsersRequest{
			Username: user.(string),
		})

		if err != nil {
			log.Fatalf("Error when calling Follow: %s", err)
		}

		if response.Success {
			c.HTML(http.StatusOK, "findPeople.html", gin.H{
				"content": response.Users,
				"user":    user,
			})
		} else {
			c.HTML(http.StatusInternalServerError, "dashboard.html", gin.H{"content": "Something went wrong, try again later."})
		}

	}
}
