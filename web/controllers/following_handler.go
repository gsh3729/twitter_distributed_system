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

		unfollowPerson := c.PostForm("unfollowPerson")

		i := helpers.IndexOf(unfollowPerson, globals.Following[username])
		globals.Following[username] = helpers.RemoveFromSlice(globals.Following[username], i)

		j := helpers.IndexOf(username, globals.Followers[unfollowPerson])
		globals.Followers[unfollowPerson] = helpers.RemoveFromSlice(globals.Followers[unfollowPerson], j)

		c.Redirect(http.StatusMovedPermanently, "/dashboard")
	}
}
