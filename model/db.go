package model

import (
	"github.com/labstack/gommon/log"
	"github.com/yjymh/songlist-go/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)

var db *gorm.DB

func DB() *gorm.DB {
	if db == nil {
		newDb, err := newDB()
		if err != nil {
			panic(err)
		}
		if newDb == nil {
			panic("db为空")
		}
		db = newDb
	}
	return db
}

func newDB() (*gorm.DB, error) {
	db, err := newMysql()
	if err != nil {
		log.Error("无法链接到数据，请检查数据库配置是否正确")
		return nil, err
	}
	return db, nil
}

func newMysql() (*gorm.DB, error) {
	dsn := conf.Conf.DB.UserName + ":" + conf.Conf.DB.PassWord + "@tcp(" + conf.Conf.DB.Host + ":" + strconv.Itoa(conf.Conf.DB.Port) + ")/" + conf.Conf.DB.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	log.Debug(dsn)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
