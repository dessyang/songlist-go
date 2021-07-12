package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/yjymh/songlist-go/cache"
	"net/http"
)

// GetSongList 展示歌曲列表
// @Router /api/v1/songs?singer=%s&album=%s [get]
func GetSongList(c *gin.Context) {
	songs := cache.Cache()
	if songs == nil {
		c.JSON(http.StatusOK, "[]")
		return
	}
	c.JSON(http.StatusOK, songs)
}
