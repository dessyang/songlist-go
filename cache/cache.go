package cache

import (
	"github.com/yjymh/songlist-go/model/song"
	"github.com/yjymh/songlist-go/service/song_service"
)

var (
	cache *[]song.Info
)

func Cache() *[]song.Info {
	var err error
	if cache == nil {
		cache, err = NewCache()
		if err != nil {
			return cache
		}
	}
	return cache
}

func NewCache() (*[]song.Info, error) {
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
