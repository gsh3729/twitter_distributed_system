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
		// if user != nil {
		// 	c.Redirect(http.StatusAccepted, "/dashboard")
		// 	return
		// }

		

		c.HTML(http.StatusOK, "composeTweet.html", gin.H{
			"user":    user,
		})
	}
} 

func TweetPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		// if user != nil {
		// 	c.Redirect(http.StatusAccepted, "/dashboard")
		// 	return
		// }
		

		

		c.HTML(http.StatusOK, "home.html", gin.H{
			"user":    user,
		})
	}
} 

