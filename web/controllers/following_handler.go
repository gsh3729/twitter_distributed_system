package controllers

import (
	"github.com/gin-contrib/sessions"

	// "log"
	"net/http"

	"github.com/gin-gonic/gin"

	globals "proj/web/globals"
	// helpers "proj/web/helpers"
)

func FollowingGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		if user != nil {
			c.Redirect(http.StatusAccepted, "/dashboard")
			return
		}
		
		username := c.Request.URL.Path[len("/followers/"):]
		userFollowers := globals.Following[username]


		c.HTML(http.StatusOK, "following.html", gin.H{
			"content": userFollowers,
			"user":    user,
		})
	}
} 

