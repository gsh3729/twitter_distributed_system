package controllers

import (
	"github.com/gin-contrib/sessions"

	"net/http"

	"github.com/gin-gonic/gin"

	globals "proj/web/globals"

	"time"
)

func TweetGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)

		c.HTML(http.StatusOK, "composeTweet.html", gin.H{
			"user": user,
		})
	}
}

func TweetPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)

		tweetMsg := c.PostForm("tweetMsg")

		tweet := globals.Tweet{
			Time: time.Now(),
			Text: tweetMsg,
			User: user.(string),
		}

		globals.Tweets[user.(string)] = append(globals.Tweets[user.(string)], tweet)

		c.Redirect(http.StatusMovedPermanently, "/dashboard")
	}
}
