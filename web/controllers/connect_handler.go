package controllers

import (
	"github.com/gin-contrib/sessions"

	"net/http"

	"github.com/gin-gonic/gin"

	globals "proj/web/globals"
)

// func ConnectPostHandler() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		session := sessions.Default(c)
// 		user := session.Get(globals.Userkey)
// 		if user != nil {
// 			c.Redirect(http.StatusAccepted, "/dashboard")
// 			return
// 		}

// 		username := c.Request.URL.Path[len("/connect/"):]
// 		connectTo := c.PostForm("connectTo")
// 		userFollowing := globals.Following[username]
// 		userFollowing = append(userFollowing, connectTo)

// 		userFollowers := globals.Followers[connectTo]
// 		userFollowers = append(userFollowers, username)

// 		// c.HTML(http.StatusOK, "followers.html", gin.H{
// 		// 	"content": userFollowers,
// 		// 	"user":    user,
// 		// })
// 	}
// } 

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
		userFollowing := globals.Following[user.(string)]
		userFollowing = append(userFollowing, connectTo)

		log.Println("Following: ", userFollowing)

		userFollowers := globals.Followers[connectTo]
		userFollowers = append(userFollowers, user.(string))


		// c.HTML(http.StatusOK, "followers.html", gin.H{
		// 	"content": userFollowers,
		// 	"user":    user,
		// })
		c.HTML(http.StatusCreated, "index.html", gin.H{"content": "Connected to user successfully"})
	}
}

