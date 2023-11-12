package router

import (
	"github.com/gin-gonic/gin"

	"github.com/IainMcl/go-behind-the-scenes/internal/logging"
	"github.com/IainMcl/go-behind-the-scenes/internal/settings"
	middleware "github.com/IainMcl/go-behind-the-scenes/middleware"
	"github.com/IainMcl/go-behind-the-scenes/router/api"
)

func InitRouter() *gin.Engine {
	gin.DefaultWriter = logging.F
	// Set the gin mode to release mode
	runMode := settings.AppSettings.RunMode
	switch runMode {
	case "debug":
		gin.SetMode(gin.DebugMode)
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}
	logging.Info("Initializing the router...")
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Anonymous routes
	r.POST("api/auth/login", api.Login)
	r.POST("api/auth/register", api.Register)

	base := r.Group("/api")
	base.Use(middleware.JWT())
	{
		base.GET("/ping", api.Ping)

		authRoutes := base.Group("/auth")
		{
			authRoutes.POST("/logout", api.Logout)
		}
	}
	return r
}
