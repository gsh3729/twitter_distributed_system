package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"net/http"

	follow "backend/follow"
)

func FollowersGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")

		username := user.(string)
		userFollowers := follow.GetUserFollowers(username)

		c.HTML(http.StatusOK, "followers.html", gin.H{
			"content": userFollowers,
			"user":    user,
		})
	}
}

func FollowingGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")

		username := user.(string)
		userFollowing := follow.GetUserFollowing(username)

		c.HTML(http.StatusOK, "following.html", gin.H{
			"content": userFollowing,
			"user":    user,
		})
	}
}
