package frontend

import (
	routes "proj/frontend/router"
	middleware "proj/web/auth"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("frontend/templates/*.html")

	router.Use(sessions.Sessions("session", cookie.NewStore([]byte("secret"))))

	public := router.Group("/")
	routes.PublicRoutes(public)

	private := router.Group("/")
	private.Use(middleware.AuthRequired)
	routes.PrivateRoutes(private)

	router.Run("0.0.0.0:8000")
}
