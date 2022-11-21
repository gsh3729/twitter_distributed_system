package controllers

import (
	"github.com/gin-contrib/sessions"

	"net/http"

	"github.com/gin-gonic/gin"

	globals "proj/web/globals"
)

func HomepageGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		if user != nil {
			c.Redirect(http.StatusAccepted, "/dashboard")
			return
		}

		var feed []globals.Tweet
		for _, element := range globals.Following[user] {
			feed = append(feed, globals.Tweets[element])
		}

		sort.Slice(feed, func(i, j int) bool {
			return feed[i].time < feed[j].time
		})


		c.HTML(http.StatusOK, "home.html", gin.H{
			"content": feed,
			"user":    user,
		})
	}
} 

