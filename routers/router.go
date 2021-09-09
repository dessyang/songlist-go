package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/yjymh/songlist-go/api/v1"
	"github.com/yjymh/songlist-go/middleware"
	"github.com/yjymh/songlist-go/model"
	"github.com/yjymh/songlist-go/pkg/e"
	"net/http"
)

var store = cookie.NewStore([]byte("123"))

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	r.Use(sessions.Sessions("session-name", store))

	r.Use(middleware.LoggerToFile())

	// 需要认证使用得api
	authApiV1 := r.Group("/api/v1").Use(middleware.Session())
	{
		authApiV1.POST("/import/song", v1.ImportMusic)
	}

	// 公开api
	pubApiV1 := r.Group("/api/v1")
	{
		pubApiV1.POST("/auth/register", v1.Register)
		pubApiV1.POST("/auth/login", v1.Login)
		pubApiV1.DELETE("/auth/logout", v1.Logout)
		pubApiV1.GET("/song/:name", v1.GetSongList)
	}
	// 捕获404
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, model.R{}.Result(e.NotFound))
	})

	return r
}
