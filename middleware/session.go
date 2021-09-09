package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/yjymh/songlist-go/model"
	"github.com/yjymh/songlist-go/pkg/e"
	"net/http"
)

// Session 判断是否登录了
func Session() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		auth := session.Get("auth")
		a, ok := (auth).(model.Auth)
		if !ok || a.ID == 0 {
			c.Abort()
			c.JSON(http.StatusOK, model.R{}.Result(e.NotLogin))
		}
	}
}
