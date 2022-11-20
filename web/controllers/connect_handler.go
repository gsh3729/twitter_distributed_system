package controllers

import (
	"github.com/gin-contrib/sessions"

	"net/http"

	"github.com/gin-gonic/gin"

	globals "proj/web/globals"
)

func ConnectPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		if user != nil {
			c.Redirect(http.StatusAccepted, "/dashboard")
			return
		}

		username := c.Request.URL.Path[len("/connect/"):]
		connectTo := c.PostForm("connectTo")
		userFollowers := globals.Followers[username]
		userFollowers = append(userFollowers, connectTo)

		// c.HTML(http.StatusOK, "followers.html", gin.H{
		// 	"content": userFollowers,
		// 	"user":    user,
		// })
	}
} 

