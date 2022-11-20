package controllers

import (
	"github.com/gin-contrib/sessions"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	globals "proj/web/globals"
	helpers "proj/web/helpers"
)

func FollowersGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		if user != nil {
			c.Redirect(http.StatusAccepted, "/dashboard")
			return
		}
		c.Request.URL.Path[len("/followers/"):]

		username := c.PostForm("username")

		for key, element := range followers {
			// fmt.Println("Key:", key, "=>", "Element:", element)
			if key == : {

			}

		}
		c.HTML(http.StatusOK, "followers.html", gin.H{
			"content": "",
			"user":    user,
		})
	}
} 

