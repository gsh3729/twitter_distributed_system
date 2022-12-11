package controllers

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	connect "proj/web/connect"
	findPeople "proj/web/findPeople"
	followers "proj/web/followers"
	following "proj/web/following"
	globals "proj/web/globals"
	helpers "proj/web/helpers"
	homepage "proj/web/homepage"
	tweet "proj/web/tweet"
)

func SignupGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		if user != nil {
			c.Redirect(http.StatusAccepted, "/dashboard")
			return
		}
		c.HTML(http.StatusOK, "signup.html", gin.H{})
	}
}

func LoginGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		if user != nil {
			c.HTML(http.StatusBadRequest, "login.html",
				gin.H{
					"content": "Please logout first",
					"user":    user,
				})
			return
		}
		c.HTML(http.StatusOK, "login.html", gin.H{})
	}
}

func SignupPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		if user != nil {
			c.Redirect(http.StatusAccepted, "/dashboard")
			return
		}

		username := c.PostForm("username")
		password := c.PostForm("password")

		if helpers.EmptyUserPass(username, password) {
			c.HTML(http.StatusBadRequest, "signup.html", gin.H{"content": "Parameters can't be empty"})
			return
		}

		globals.UserPass[username] = password

		c.HTML(http.StatusCreated, "index.html", gin.H{"content": "Created user successfully"})
	}
}

func LoginPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		if user != nil {
			c.Redirect(http.StatusAccepted, "/dashboard")
			return
		}

		username := c.PostForm("username")
		password := c.PostForm("password")

		if helpers.EmptyUserPass(username, password) {
			c.HTML(http.StatusBadRequest, "login.html", gin.H{"content": "Parameters can't be empty"})
			return
		}

		if !helpers.CheckUserPass(username, password) {
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{"content": "Incorrect username or password"})
			return
		}

		session.Set(globals.Userkey, username)
		if err := session.Save(); err != nil {
			c.HTML(http.StatusInternalServerError, "login.html", gin.H{"content": "Failed to save session"})
			return
		}

		c.Redirect(http.StatusMovedPermanently, "/dashboard")
	}
}

func LogoutGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)

		if user == nil {
			log.Println("Invalid session token")
			return
		}
		session.Delete(globals.Userkey)
		if err := session.Save(); err != nil {
			log.Println("Failed to save session:", err)
			return
		}

		c.Redirect(http.StatusMovedPermanently, "/")
	}
}

func IndexGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"user": user,
		})
	}
}

func DashboardGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
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
		user := session.Get(globals.Userkey)

		feed := homepage.GetTweetsForHomepage(user.(string))

		c.HTML(http.StatusOK, "home.html", gin.H{
			"content": feed,
			"user":    user,
		})
	}
}

func TweetGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)

		c.HTML(http.StatusOK, "composeTweet.html", gin.H{
			"user": user,
		})
	}
}

func TweetPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)

		tweetMsg := c.PostForm("tweetMsg")

		tweet.PostTweet(user.(string), tweetMsg)

		c.Redirect(http.StatusMovedPermanently, "/dashboard")
	}
}

func FollowingGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)

		username := user.(string)
		userFollowing := following.GetUserFollowing(username)

		c.HTML(http.StatusOK, "following.html", gin.H{
			"content": userFollowing,
			"user":    user,
		})
	}
}

func FollowersGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)

		username := user.(string)
		userFollowers := followers.GetUserFollowers(username)

		c.HTML(http.StatusOK, "followers.html", gin.H{
			"content": userFollowers,
			"user":    user,
		})
	}
}

func FindPeopleGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)

		people := findPeople.GetPeopleForUser(user.(string))

		c.HTML(http.StatusOK, "findPeople.html", gin.H{
			"content": people,
			"user":    user,
		})
	}
}

func FollowPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)

		connectTo := c.PostForm("connectTo")

		connect.Follow(user.(string), connectTo)

		c.Redirect(http.StatusMovedPermanently, "/dashboard")
	}
}

func UnfollowPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		username := user.(string)

		unfollowPerson := c.PostForm("unfollowPerson")

		connect.Unfollow(username, unfollowPerson)

		c.Redirect(http.StatusMovedPermanently, "/dashboard")
	}
}
