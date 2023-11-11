package main

import (
	"fmt"
	"net/http"

	"github.com/IainMcl/go-behind-the-scenes/internal/logging"
	"github.com/IainMcl/go-behind-the-scenes/internal/settings"
	"github.com/IainMcl/go-behind-the-scenes/models"
	"github.com/IainMcl/go-behind-the-scenes/router"
)

func init() {
	// log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	settings.Setup()
	models.Setup()
	logging.Setup()
}

func main() {
	logging.Info("Starting the application...")

	routersInit := router.InitRouter()
	readTimeout := settings.ServerSettings.ReadTimeout
	writeTimeout := settings.ServerSettings.WriteTimeout
	endPoint := fmt.Sprintf(":%d", settings.ServerSettings.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	server.ListenAndServe()
	logging.Info("Closing the application...")
}
