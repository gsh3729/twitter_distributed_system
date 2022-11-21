package controllers

import (
	"github.com/gin-contrib/sessions"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	globals "proj/web/globals"
	// helpers "proj/web/helpers"
)

func FindPeopleGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		log.Println("User from followers get handler: ", user)
		// if user != nil {
		// 	log.Println("Inside if")
		// 	c.Redirect(http.StatusAccepted, "/dashboard")
		// 	return
		// }
		

		people := []string{} 
		for key, element := range globals.UserPass {
			// fmt.Println("Key:", key, "=>", "Element:", element)
			if ( key != user) {
				people = append(people, key)
			}
		}

		c.HTML(http.StatusOK, "followers.html", gin.H{
			"content": "harsha",
			"user":    user,
		})
	}
} 

