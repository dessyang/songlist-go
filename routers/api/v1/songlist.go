package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/yjymh/songlist-go/cache"
	"github.com/yjymh/songlist-go/conf"
	"github.com/yjymh/songlist-go/model"
	"github.com/yjymh/songlist-go/pkg/e"
	"github.com/yjymh/songlist-go/service/song_service"
	"github.com/yjymh/songlist-go/util"
	"math/rand"
	"net/http"
	"strconv"
)

// GetSongList 展示歌曲列表
// @Router /api/v1/songs/:name?page=${page}&random=${bool} [get]
func GetSongList(c *gin.Context) {
	R := model.R{}
	maxNum := conf.Conf.App.PageMaxNum // 单次请求的最大数量

	name := c.Param("name")
	random, _ := strconv.ParseBool(c.DefaultQuery("random", "false"))
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))

	songs := cache.GetSongsCache(name)

	if err != nil {
		// 页数不为数字
		c.JSON(http.StatusOK, R.Fail(e.PageNotNum))
		return
	}

	if songs == nil {
		//
		c.JSON(http.StatusOK, R.Fail(e.SongNotFound))
		return
	}

	if random {
		index := rand.Intn(len(songs))
		song := songs[index]
		c.JSON(http.StatusOK, song)
		return
	}
	maxPage := util.MaxPage(len(songs), maxNum)
	firstNum := (page - 1) * maxNum
	lastNum := page * maxNum

	// 判断访问的页数是否在总页数里面，根据不同情况输出不同数据
	if page > maxPage {
		c.JSON(http.StatusOK, R.Fail(e.PageOutBound))
	} else if page == maxPage {
		c.JSON(http.StatusOK, R.Success(songs[firstNum:]))
	} else {
		c.JSON(http.StatusOK, R.Success(songs[firstNum:lastNum]))
	}
}

// ImportMusic 导入歌曲
// @Router /api/v1/import [get]
// TODO 歌单导入需要权限，并且各个用户导入自己的歌单
func ImportMusic(c *gin.Context) {
	title := c.Query("title")
	flag := song_service.AddSongInfo(title)
	// cache.UpdateCache()
	c.JSON(http.StatusOK, flag)
}
