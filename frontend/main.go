package main

import (
	middleware "frontend/auth"
	routes "frontend/router"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("./templates/*.html")

	router.Use(sessions.Sessions("session", cookie.NewStore([]byte("secret"))))

	public := router.Group("/")
	routes.PublicRoutes(public)

	private := router.Group("/")
	private.Use(middleware.AuthRequired)
	routes.PrivateRoutes(private)

	router.Run("0.0.0.0:8000")
}
