package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/yjymh/songlist-go/middleware"
	"github.com/yjymh/songlist-go/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 需要认证使用得api
	authApiV1 := r.Group("/api/v1").Use(middleware.JWTAuth())
	{
		authApiV1.POST("/import", v1.ImportMusic)
	}

	// 公开api
	pubApiV1 := r.Group("/api/v1")
	{
		pubApiV1.POST("/auth/:method", v1.Auth)
		pubApiV1.GET("/songs/:name", v1.GetSongList)
	}

	return r
}
