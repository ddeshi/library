package main

import (
	"fmt"
	"github.com/ddeshi/library/pkg/config"
	"github.com/ddeshi/library/pkg/database"
	"github.com/ddeshi/library/pkg/logs"
	"github.com/ddeshi/library/pkg/routers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	config.Init()
	database.DBInit()
	logs.LogInit()

	gin.SetMode(config.ServerCfg.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := config.ServerCfg.ReadTimeout
	writeTimeout := config.ServerCfg.WriteTimeout
	endPoint := fmt.Sprintf(":%d", config.ServerCfg.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
