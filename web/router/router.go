package routes

import (
	"github.com/gin-gonic/gin"

	controllers "proj/web/controllers"
)

func PublicRoutes(g *gin.RouterGroup) {

	g.GET("/login", controllers.LoginGetHandler())
	g.POST("/login", controllers.LoginPostHandler())
	g.GET("/signup", controllers.SignupGetHandler())
	g.POST("/signup", controllers.SignupPostHandler())
	g.GET("/", controllers.IndexGetHandler())

}

func PrivateRoutes(g *gin.RouterGroup) {
	g.GET("/followers", controllers.FollowersGetHandler())
	g.GET("/following", controllers.FollowingGetHandler())

	g.GET("/find", controllers.FindPeopleGetHandler())
	g.POST("/connect", controllers.FollowPostHandler())
	g.POST("/unfollow", controllers.UnfollowPostHandler())

	g.GET("/compose", controllers.TweetGetHandler())
	g.POST("/compose", controllers.TweetPostHandler())

	g.GET("/home", controllers.HomepageGetHandler())
	g.GET("/dashboard", controllers.DashboardGetHandler())
	g.GET("/logout", controllers.LogoutGetHandler())

}
