package model

import "gorm.io/gorm"

// Song 歌曲信息
type Song struct {
	gorm.Model
	Title     string   // 歌名
	Album     string   // 专辑
	Artists   []Artist `gorm:"many2many:song_artists;"`
	Time      int      // 歌曲长度
	PubTime   string   // 发行时间
	SourceUrl string   // 歌曲链接
}
