package model

// SongInfo 歌曲信息
type SongInfo struct {
	SongId    int    `gorm:"primarykey"` // 歌曲ID
	Title     string // 歌名
	Album     string // 专辑
	Artist    string // 歌手
	Time      int    // 歌曲长度
	PubTime   string // 发行时间
	SourceUrl string // 歌曲链接
}
