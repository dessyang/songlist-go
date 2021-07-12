package song

// Info 歌曲信息
type Info struct {
	SongId    int    `gorm:"primarykey"` // 歌曲ID
	Title     string // 歌名
	Album     string // 专辑
	Artist    string // 歌手
	AlbumPic  string // 专辑图片
	Company   string // 唱片公司
	Genre     string // 歌曲流派
	Lang      string // 歌曲语种
	Time      int    // 歌曲长度
	PubTime   string // 发行时间
	SourceUrl string // 歌曲链接
}
