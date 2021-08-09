package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/yjymh/songlist-go/model"
	"github.com/yjymh/songlist-go/pkg/e"
	"github.com/yjymh/songlist-go/service/auth_service"
	"github.com/yjymh/songlist-go/util"
	"net/http"
)

// Auth 权限管理
// Router /api/v1/login/:method [post]
func Auth(c *gin.Context) {
	method := c.Param("method")
	switch method {
	case "login":
		Login(c)
	case "logout":
		Logout(c)
	case "register":
		Register(c)
	}
	return
}

// Register 注册
// Router /api/v1/login/register [post]
func Register(c *gin.Context) {
	R := model.R{}
	// 用户名为 3 ~ n字母
	username := c.PostForm("username")
	// 密码要求：
	password := c.PostForm("password")
	// 邮箱：56789@qq.com
	email := c.PostForm("email")
	// TODO 还需要验证用户名，密码，邮箱的格式，其余全部完工

	// 参数不能为空
	if username == "" || password == "" || email == "" {
		c.JSON(http.StatusOK, R.Fail(e.ParamNotNul))
		return
	}
	// 防止用户名重复
	if auth_service.IsExistUsername(username) {
		c.JSON(http.StatusOK, R.Fail(e.RepeatUser))
		return
	}
	// 防止邮箱重复
	if auth_service.IsExistEmail(email) {
		c.JSON(http.StatusOK, R.Fail(e.RepeatEmail))
		return
	}

	flag, _ := auth_service.CreateAuthUser(username, password, email)
	if flag {
		c.JSON(http.StatusOK, R.Success("注册成功"))
	} else {
		c.JSON(http.StatusOK, R.Success("注册失败"))
	}
}

// Logout 注销
// TODO JWT的注销还不知道怎么做
func Logout(c *gin.Context) {
	c.Request.Header.Get("x-token")
}

// Login 登录
func Login(c *gin.Context) {
	R := model.R{}

	username := c.Param("username")
	password := c.Param("password")

	flag, _ := auth_service.Check(username, password)
	if !flag {
		c.JSON(http.StatusOK, R.Fail(e.AuthFail))
		return
	}
	token, err := util.GenerateToken(username, password)
	if err != nil {
		c.JSON(http.StatusOK, R.Fail(e.Fail))
	}
	// 登录成功，进行下一步操作
	c.JSON(http.StatusOK, R.Success(token))
}
