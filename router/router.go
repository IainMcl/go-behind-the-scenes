package router

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/IainMcl/go-behind-the-scenes/middleware/jwt"
	"github.com/IainMcl/go-behind-the-scenes/router/api"
)

func InitRouter() *gin.Engine {
	log.Println("Initializing the router...")
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Anonymous routes
	r.POST("api/auth/login", api.Login)

	base := r.Group("/api")
	base.Use(jwt.JWT())
	{
		base.GET("/ping", api.Ping)

		authRoutes := base.Group("/auth")
		{
			authRoutes.POST("/logout", api.Logout)
		}
	}
	return r
}
