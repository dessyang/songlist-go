package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/yjymh/songlist-go/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	pubApiV1 := r.Group("/api/v1")
	{
		pubApiV1.GET("/songs/:name", v1.GetSongList)
	}

	return r
}
