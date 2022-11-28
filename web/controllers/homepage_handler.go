package controllers

import (
	"github.com/gin-contrib/sessions"

	"net/http"

	"github.com/gin-gonic/gin"

	globals "proj/web/globals"
	"sort"
)

func HomepageGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)



		c.HTML(http.StatusOK, "home.html", gin.H{
			"content": feed,
			"user":    user,
		})
	}
}
