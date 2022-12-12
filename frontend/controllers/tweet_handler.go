package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"net/http"

	tweet "backend/tweet"
)

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

		tweet.PostTweet(user.(string), tweetMsg)

		c.Redirect(http.StatusMovedPermanently, "/dashboard")
	}
}
