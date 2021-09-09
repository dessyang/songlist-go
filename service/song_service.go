package service

import (
	"github.com/yjymh/songlist-go/model"
)

func AddSongInfo(song *model.Song, auth *model.Auth) (bool, error) {
	if IsExistSongInfo(song.Title) {
		return false, nil
	}

	auth.Songs = append(auth.Songs, *song)

	for i := 0; i < len(auth.Songs); i++ {
		for j := 0; j < len(auth.Songs[i].Artists); j++ {
			artist, _ := QueryArtist(auth.Songs[i].Artists[j].Name)
			auth.Songs[i].Artists[j] = artist
		}
	}

	model.DB().Create(auth)
	model.DB().Save(auth)

	return true, nil
}

// QuerySongByUser 通过用户名查询所有歌曲
func QuerySongByUser(user string) ([]model.Song, bool) {
	auth, b := QueryUserByUsername(user)
	if b == false {
		return nil, false
	}
	err := model.DB().Preload("Songs").Preload("Songs.Artists").Find(&auth).Error
	if err != nil {
		return nil, false
	}
	return auth.Songs, false
}

// IsExistSongInfo 是否存在同名歌曲
func IsExistSongInfo(title string) bool {
	song := model.Song{Title: title}
	model.DB().Where(&song).First(&song)
	if song.ID != 0 {
		return true
	}
	return false
}

func QuerySongInfo(title string) (model.Song, error) {
	song := model.Song{Title: title}
	err := model.DB().Model(&song).First(&song).Error
	if err != nil {
		return song, err
	}
	return song, nil
}

func QueryArtist(name string) (model.Artist, error) {
	artist := model.Artist{Name: name}
	err := model.DB().Model(&artist).Where(&artist).First(&artist).Error
	if err != nil {
		return artist, err
	}
	return artist, nil
}
