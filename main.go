package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yjymh/songlist-go/conf"
	"github.com/yjymh/songlist-go/routers"
	"net/http"
)

func init() {
	conf.InitConfig("")
}

func main() {
	gin.SetMode("release")
	routersInit := routers.InitRouter()

	endPoint := fmt.Sprintf(":%d", conf.Conf.Server.Port)
	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
