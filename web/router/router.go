// package router

// import (
// 	// "go-server/middleware"
// 	"github.com/gorilla/mux"
// )

// // Router is exported and used in main.go
// func Router() *mux.Router {

// 	router := mux.NewRouter()

// 	router.HandleFunc("/login", middleware.GetAllTask).Methods("GET", "OPTIONS")
// 	router.HandleFunc("/signup", middleware.CreateTask).Methods("POST", "OPTIONS")
// 	router.HandleFunc("/profile/{id}", middleware.TaskComplete).Methods("PUT", "OPTIONS")
// 	router.HandleFunc("/compose/tweet/{id}", middleware.UndoTask).Methods("PUT", "OPTIONS")
// 	router.HandleFunc("/connect/people/{id}", middleware.DeleteTask).Methods("DELETE", "OPTIONS")
// 	// router.HandleFunc("/api/deleteAllTask", middleware.DeleteAllTask).Methods("DELETE", "OPTIONS")
// 	return router
// }

package routes

import (
	"github.com/gin-gonic/gin"

	controllers "web/controllers"
)

func PublicRoutes(g *gin.RouterGroup) {

	g.GET("/login", controllers.LoginGetHandler())
	g.POST("/login", controllers.LoginPostHandler())
	g.GET("/", controllers.IndexGetHandler())

}

func PrivateRoutes(g *gin.RouterGroup) {

	g.GET("/dashboard", controllers.DashboardGetHandler())
	g.GET("/logout", controllers.LogoutGetHandler())

}