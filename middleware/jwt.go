package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/yjymh/songlist-go/model"
	"github.com/yjymh/songlist-go/pkg/e"
	"github.com/yjymh/songlist-go/util"
	"net/http"
)

// JWTAuth 权限认证
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int

		code = e.Success

		R := model.R{}
		token := c.Request.Header.Get("token")
		if token == "" {
			code = e.NotLogin
		} else {
			_, err := util.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = e.JwtTokenTimeout
				default:
					code = e.JwtTokenFail
				}
			}
		}

		if code != e.Success {
			c.JSON(http.StatusOK, R.Fail(code))
			c.Abort()
			return
		}

		c.Next()
	}
}
