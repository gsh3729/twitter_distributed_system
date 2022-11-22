package controllers

import (
	"github.com/gin-contrib/sessions"

	"net/http"

	"github.com/gin-gonic/gin"

	globals "proj/web/globals"
	"log"
)



func ConnectPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		// if user != nil {
		// 	c.Redirect(http.StatusAccepted, "/dashboard")
		// 	return
		// }

		// username := c.Request.URL.Path[len("/connect/"):]
		connectTo := c.PostForm("connectTo")
		globals.Following[user.(string)] = append(globals.Following[user.(string)], connectTo)

		log.Println("Following: ", globals.Following[user.(string)])

		globals.Followers[connectTo] = append(globals.Followers[connectTo], user.(string))
		log.Println("Followers: ", globals.Followers[connectTo])

		// c.HTML(http.StatusOK, "followers.html", gin.H{
		// 	"content": userFollowers,
		// 	"user":    user,
		// })
		// c.HTML(http.StatusCreated, "index.html", gin.H{"content": "Connected to user successfully"})
		c.Redirect(http.StatusMovedPermanently, "/dashboard")
	}
}

