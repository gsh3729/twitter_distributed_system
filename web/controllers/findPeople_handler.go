package controllers

import (
	"github.com/gin-contrib/sessions"

	"net/http"

	"github.com/gin-gonic/gin"

	globals "proj/web/globals"
	// helpers "proj/web/helpers"
)

func FindPeopleGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)

		people := []string{}
		for key := range globals.UserPass {

			if key != user {
				people = append(people, key)
			}
		}

		c.HTML(http.StatusOK, "findPeople.html", gin.H{
			"content": people,
			"user":    user,
		})
	}
}
