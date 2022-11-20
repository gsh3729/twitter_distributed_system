package controllers

import (
	"github.com/gin-contrib/sessions"

	"net/http"

	"github.com/gin-gonic/gin"

	globals "proj/web/globals"
)

func HomepageGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		if user != nil {
			c.Redirect(http.StatusAccepted, "/dashboard")
			return
		}

		feed 

		username := c.Request.URL.Path[len("//"):]
		connectTo := c.PostForm("connectTo")
		userFollowing := globals.Following[username]
		userFollowing = append(userFollowing, connectTo)

		userFollowers := globals.Followers[connectTo]
		userFollowers = append(userFollowers, username)

		c.HTML(http.StatusOK, "home.html", gin.H{
			"content": userFollowers,
			"user":    user,
		})
	}
} 

