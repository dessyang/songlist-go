package model

import "gorm.io/gorm"

type UserSong struct {
	gorm.Model
	UserId int
	SongId int
}
