package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"net/http"

	userservice "backend/user"
)

func FindPeopleGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")

		people := userservice.GetPeopleForUser(user.(string))

		c.HTML(http.StatusOK, "findPeople.html", gin.H{
			"content": people,
			"user":    user,
		})
	}
}
