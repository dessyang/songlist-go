package util

import "github.com/yjymh/songlist-go/conf"

func Setup() {
	jwtSecret = []byte(conf.Conf.App.JwtSecret)
}
