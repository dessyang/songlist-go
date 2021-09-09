package model

// SetUp 初始化数据库表
func SetUp() {
	DB().AutoMigrate(Auth{})
	DB().AutoMigrate(Artist{})
	DB().AutoMigrate(Song{})
}
