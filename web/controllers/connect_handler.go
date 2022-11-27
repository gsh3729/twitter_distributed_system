package controllers

import (
	"github.com/gin-contrib/sessions"

	"net/http"

	"github.com/gin-gonic/gin"

	"log"
	globals "proj/web/globals"
)

func ConnectPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)

		connectTo := c.PostForm("connectTo")

		if globals.
		globals.Following[user.(string)] = append(globals.Following[user.(string)], connectTo)
		globals.Followers[connectTo] = append(globals.Followers[connectTo], user.(string))

		c.Redirect(http.StatusMovedPermanently, "/dashboard")
	}
}
