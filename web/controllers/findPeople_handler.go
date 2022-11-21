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
		for key := range globals.UserPass {
			log.Println("key: ", key)
			if (key != user) {
				log.Println("suc")
				people = append(people, key)
			}
		}

		c.HTML(http.StatusOK, "findPeople.html", gin.H{
			"content": people,
			"user":    user,
		})
	}
} 

