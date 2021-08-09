package auth_service

import (
	"errors"
	"github.com/yjymh/songlist-go/model"
)

// CreateAuthUser 创建一个用户
// 需要判断一下username email是否存在重复项
func CreateAuthUser(username, password, email string) (bool, error) {
	auth := new(model.Auth)
	auth.Username = username
	auth.Password = password
	auth.Email = email

	if !IsExistUsername(username) {
		return false, errors.New("用户名重复")
	}
	if !IsExistEmail(email) {
		return false, errors.New("邮箱重复")
	}

	err := model.DB().Create(&auth).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

// IsExistUsername 判断是否存在用户名
func IsExistUsername(username string) bool {
	auth := model.Auth{}
	model.DB().Where("username = ?", username).First(&auth)
	if auth.ID != 0 {
		return true
	}
	return false
}

// IsExistEmail 判断是否存在改邮箱
func IsExistEmail(email string) bool {
	auth := model.Auth{}
	model.DB().Where("email = ?", email).First(&auth)
	if auth.ID != 0 {
		return true
	}
	return false
}

// Check 登录验证
func Check(username, password string) (bool, error) {
	auth := new(model.Auth)

	err := model.DB().Where("username = ? and password = ?", username, password).First(&auth).Error
	if err != nil {
		return false, err
	}
	if auth.ID != 0 {
		return true, nil
	}
	return false, nil
}
