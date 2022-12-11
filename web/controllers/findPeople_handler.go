package controllers

import (
	"github.com/gin-contrib/sessions"

	"net/http"

	"github.com/gin-gonic/gin"

	globals "proj/web/globals"
	findPeople "proj/web/findPeople"
)

func FindPeopleGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)

		people := findPeople.GetPeopleForUser(user.(string))

		c.HTML(http.StatusOK, "findPeople.html", gin.H{
			"content": people,
			"user":    user,
		})
	}
}
