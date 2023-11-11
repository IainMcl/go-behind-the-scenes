package router

import (
	"github.com/gin-gonic/gin"

	"github.com/IainMcl/go-behind-the-scenes/internal/logging"
	"github.com/IainMcl/go-behind-the-scenes/middleware/jwt"
	"github.com/IainMcl/go-behind-the-scenes/router/api"
)

func InitRouter() *gin.Engine {
	logging.Info("Initializing the router...")
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Anonymous routes
	r.POST("api/auth/login", api.Login)
	r.POST("api/auth/register", api.Register)

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
