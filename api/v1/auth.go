package v1

import (
	"encoding/gob"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/yjymh/songlist-go/conf"
	"github.com/yjymh/songlist-go/middleware"
	"github.com/yjymh/songlist-go/model"
	"github.com/yjymh/songlist-go/pkg/e"
	service "github.com/yjymh/songlist-go/service"
	"github.com/yjymh/songlist-go/util"
	"net/http"
)

func init() {
	// 注册结构体,这样才能在session使用该结构体
	// https://github.com/gin-contrib/sessions/issues/134
	gob.Register(model.Auth{})
}

// Register 注册
// Router /api/v1/login/register [post]
func Register(c *gin.Context) {
	R := model.R{}
	if !conf.Conf.App.Register {
		c.JSON(http.StatusOK, R.Result(e.RegistrationNotAllowed))
		return
	}
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")

	if util.IsEmpty(username) || util.IsEmpty(password) || util.IsEmpty(email) {
		c.JSON(http.StatusOK, R.Result(e.MissParam))
		return
	}

	if !util.IsUsername(username) {
		c.JSON(http.StatusOK, R.Result(e.UsernameFormatError))
		return
	}
	if !util.IsEmail(email) {
		c.JSON(http.StatusOK, R.Result(e.EmailFormatError))
		return
	}
	if !util.IsPassword(password) {
		c.JSON(http.StatusOK, R.Result(e.PasswordFormatError))
		return
	}

	if service.IsExistUsername(username) {
		c.JSON(http.StatusOK, R.Result(e.RepeatUser))
		return
	}
	if service.IsExistEmail(email) {
		c.JSON(http.StatusOK, R.Result(e.RepeatEmail))
		return
	}
	flag, err := service.CreateAuthUser(username, password, email)
	if err != nil {
		middleware.Logger().Error(err)
	}
	if !flag {
		c.JSON(http.StatusOK, R.Result(e.Fail))
		return
	}
	c.JSON(http.StatusOK, R.Success("注册成功"))
}

// Logout 注销
func Logout(c *gin.Context) {
	R := model.R{}
	session := sessions.Default(c)
	auth := session.Get("auth")
	if auth != nil {
		session.Clear()
		session.Save()
		c.JSON(http.StatusOK, R.Success("退出成功"))
		return
	}
	c.JSON(http.StatusOK, R.Success("已经退出,无需重复操作"))
}

// Login 登录
func Login(c *gin.Context) {
	session := sessions.Default(c)
	R := model.R{}

	username := c.PostForm("username")
	password := c.PostForm("password")

	if util.IsEmpty(username) || util.IsEmpty(password) {
		c.JSON(http.StatusOK, R.Result(e.MissParam))
		return
	}

	auth, flag := service.CheckAuth(username, password)

	if flag {
		session.Set("auth", auth)
		session.Save()
		c.JSON(http.StatusOK, R.Success("登录成功"))
		return
	}

	c.JSON(http.StatusOK, R.Result(e.AuthFail))
}
