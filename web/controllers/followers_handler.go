package controllers

import (
	"github.com/gin-contrib/sessions"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	globals "proj/web/globals"
	helpers "proj/web/helpers"
)

func FollowersGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		if user != nil {
			c.Redirect(http.StatusAccepted, "/dashboard")
			return
		}
		
		c.HTML(http.StatusOK, "followers.html", gin.H{
			"content": "",
			"user":    user,
		})
	}
} 

