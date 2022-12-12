package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"net/http"

	homepage "backend/homepage"
)

func IndexGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"user": user,
		})
	}
}

func DashboardGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")

		feed := homepage.GetTweetsForHomepage(user.(string))

		c.HTML(http.StatusOK, "dashboard.html", gin.H{
			"content": feed,
			"user":    user,
		})
	}
}

func HomepageGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")

		feed := homepage.GetTweetsForHomepage(user.(string))

		c.HTML(http.StatusOK, "home.html", gin.H{
			"content": feed,
			"user":    user,
		})
	}
}
