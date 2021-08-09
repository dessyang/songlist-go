package model

func SetUp() {
	DB().AutoMigrate(UserSong{})
	DB().AutoMigrate(Auth{})
	DB().AutoMigrate(SongInfo{})
}
