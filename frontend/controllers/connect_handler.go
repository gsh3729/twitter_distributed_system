package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"net/http"

	connect "backend/connect"
)

func FollowPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")

		connectTo := c.PostForm("connectTo")

		connect.Follow(user.(string), connectTo)

		c.Redirect(http.StatusMovedPermanently, "/dashboard")
	}
}

func UnfollowPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")
		username := user.(string)

		unfollowPerson := c.PostForm("unfollowPerson")

		connect.Unfollow(username, unfollowPerson)

		c.Redirect(http.StatusMovedPermanently, "/dashboard")
	}
}
