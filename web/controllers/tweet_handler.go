package controllers

import (
	"github.com/gin-contrib/sessions"

	"net/http"

	"github.com/gin-gonic/gin"

	globals "proj/web/globals"
	tweet "proj/web/tweet"

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

		tweet.PostTweet(user.(string), tweetMsg)

		c.Redirect(http.StatusMovedPermanently, "/dashboard")
	}
}
