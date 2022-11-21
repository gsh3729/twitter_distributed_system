package controllers

import (
	"github.com/gin-contrib/sessions"

	"net/http"
	"log"

	"github.com/gin-gonic/gin"

	globals "proj/web/globals"
	"sort"
)

func HomepageGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)

		if user != nil {
			c.Redirect(http.StatusAccepted, "/dashboard")
			return
		}
		log.Println("user : ", user)
		var feed []globals.Tweet
		for _, element := range globals.Following[user.(string)] { //check once
			// feed = append(feed, globals.Tweets[element])
			for _, tweet := range globals.Tweets[element] {
				feed = append(feed, tweet)
			}

		}

		sort.SliceStable(feed[:], func(i, j int) bool {
			return feed[i].Time < feed[j].Time
		})

		c.HTML(http.StatusOK, "home.html", gin.H{
			"content": "Posted tweet successfully",
			"user":    user,
		})
	}
} 

