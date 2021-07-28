package cache

import (
	"errors"
	"github.com/yjymh/songlist-go/model"
	"github.com/yjymh/songlist-go/service/song_service"
)

var (
	songsCache = make(map[string][]model.SongInfo)
)

func IsExistSongsCache(key string) bool {
	if GetSongsCache(key) != nil {
		return true
	} else {
		return false
	}
}

func SetSongsCache(key string, value []model.SongInfo) error {
	if !IsExistSongsCache(key) {
		songsCache[key] = value
		return nil
	} else {
		return errors.New("里面已经存在该key，请使用Update方法")
	}
}

func GetSongsCache(key string) []model.SongInfo {
	if songs, ok := songsCache[key]; !ok {
		// 从数据库取出该用户的歌曲
		songsCache[key], _ = song_service.QuerySongInfo(key)
		return songsCache[key]
	} else {
		return songs
	}
}

func UpdateSongsCache(key string, value []model.SongInfo) error {
	if IsExistSongsCache(key) {
		songsCache[key] = value
		return nil
	} else {
		return errors.New("该key值为空，请使用set方法")
	}
}

func DeleteSongsCache(key string) error {
	if IsExistSongsCache(key) {
		delete(songsCache, key)
		return nil
	} else {
		return errors.New("该key值不存在")
	}
}
