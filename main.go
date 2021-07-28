package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yjymh/songlist-go/conf"
	"github.com/yjymh/songlist-go/routers"
	"github.com/yjymh/songlist-go/util"
	"net/http"
)

func init() {
	conf.Setup("")
	util.Setup()
}

func main() {
	gin.SetMode(conf.Conf.App.Mode)
	routersInit := routers.InitRouter()

	endPoint := fmt.Sprintf(":%d", conf.Conf.Server.Port)
	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
