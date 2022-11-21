package controllers

import (
	"github.com/gin-contrib/sessions"

	// "log"
	"net/http"

	"github.com/gin-gonic/gin"

	globals "proj/web/globals"
)

func FollowingGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		if user != nil {
			c.Redirect(http.StatusAccepted, "/dashboard")
			return
		}
		
		username := c.Param("userId")
		userFollowers := globals.Following[username]


		c.HTML(http.StatusOK, "following.html", gin.H{
			"content": userFollowers,
			"user":    user,
		})
	}
} 

