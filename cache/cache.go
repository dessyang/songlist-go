package cache

import (
	"github.com/yjymh/songlist-go/model"
	"github.com/yjymh/songlist-go/service/song_service"
)

var (
	cache *[]model.SongInfo
)

func Cache() *[]model.SongInfo {
	var err error
	if cache == nil {
		cache, err = NewCache()
		if err != nil {
			return cache
		}
	}
	return cache
}

func NewCache() (*[]model.SongInfo, error) {
	songs, err := song_service.QuerySongInfo()
	if err != nil {
		return nil, err
	}
	return songs, nil
}

func UpdateCache() bool {
	var err error
	cache, err = NewCache()
	if err != nil {
		return false
	}
	return true
}
