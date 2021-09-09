package service

import (
	"github.com/yjymh/songlist-go/middleware"
	"github.com/yjymh/songlist-go/model"
)

// CreateAuthUser 创建一个用户
func CreateAuthUser(username, password, email string) (bool, error) {
	auth := model.Auth{
		Username: username,
		Password: password,
		Email:    email,
	}
	err := model.DB().Create(&auth).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

// IsExistUsername 判断是否存在用户名
func IsExistUsername(username string) bool {
	auth := model.Auth{Username: username}
	model.DB().Where(&auth).First(&auth)
	if auth.ID != 0 {
		return true
	}
	return false
}

// IsExistEmail 判断是否存在改邮箱
func IsExistEmail(email string) bool {
	auth := model.Auth{Email: email}
	model.DB().Where(&auth).First(&auth)
	if auth.ID != 0 {
		return true
	}
	return false
}

// CheckAuth 登录验证
func CheckAuth(username, password string) (model.Auth, bool) {
	flag := false
	auth := model.Auth{Username: username, Password: password}
	err := model.DB().Where(&auth).First(&auth).Error
	if err != nil {
		middleware.Logger().Errorln(err)
	}
	if auth.ID != 0 {
		flag = true
	}
	return auth, flag
}

// QueryUserByUsername 查询用户
func QueryUserByUsername(username string) (model.Auth, bool) {
	var auth model.Auth
	var flag bool

	auth.Username = username
	err := model.DB().Where(&auth).First(&auth).Error

	if err != nil {
		flag = false
	} else {
		if auth.ID == 0 {
			flag = false
		} else {
			flag = true
		}
	}
	return auth, flag
}
