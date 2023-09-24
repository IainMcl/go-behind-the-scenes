package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/IainMcl/go-behind-the-scenes/internal/logging"
	"github.com/IainMcl/go-behind-the-scenes/internal/settings"
	"github.com/IainMcl/go-behind-the-scenes/router"
)

func init() {
	// log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	settings.Setup()
	logging.Setup()
}

func main() {
	log.Println("Starting the application...")

	routersInit := router.InitRouter()
	readTimeout := settings.ServerSetting.ReadTimeout
	writeTimeout := settings.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", settings.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	server.ListenAndServe()

	log.Printf("[info] start http server listening %s", endPoint)
}
