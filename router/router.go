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

	base := r.Group("/api")
	base.Use(jwt.JWT())
	{
		r.GET("/ping", api.Ping)

	}

	return r
}
