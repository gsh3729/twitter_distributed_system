package controllers

import (
	"github.com/gin-contrib/sessions"

	"net/http"

	"github.com/gin-gonic/gin"

	connect "proj/web/connect"
	globals "proj/web/globals"
)

func FollowPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)

		connectTo := c.PostForm("connectTo")

		connect.Follow(user.(string), connectTo)

		c.Redirect(http.StatusMovedPermanently, "/dashboard")
	}
}

func UnfollowPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		username := user.(string)

		unfollowPerson := c.PostForm("unfollowPerson")

		connect.Unfollow(username, unfollowPerson)

		c.Redirect(http.StatusMovedPermanently, "/dashboard")
	}
}
