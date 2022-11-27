package controllers

import (
	"github.com/gin-contrib/sessions"

	"net/http"

	"github.com/gin-gonic/gin"

	globals "proj/web/globals"
	helpers "proj/web/helpers"
)

func FollowingGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)

		username := user.(string)
		userFollowers := globals.Following[username]

		c.HTML(http.StatusOK, "following.html", gin.H{
			"content": userFollowers,
			"user":    user,
		})
	}
}

func UnfollowPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		username := user.(string)



		c.Redirect(http.StatusMovedPermanently, "/dashboard")
	}
}
