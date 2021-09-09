package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/yjymh/songlist-go/conf"
	"github.com/yjymh/songlist-go/model"
	"github.com/yjymh/songlist-go/pkg/e"
	"github.com/yjymh/songlist-go/service"
	"github.com/yjymh/songlist-go/util"
	"math/rand"
	"net/http"
	"strconv"
)

type song struct {
	SongId  uint
	Title   string
	Album   string
	Artists []artist
}

type artist struct {
	ArtistId uint
	Name     string
}

// GetSongList 展示歌曲列表
// @Router /api/v1/songs/:name?page=${page}&random=${bool} [get]
func GetSongList(c *gin.Context) {
	R := model.R{}
	maxNum := conf.Conf.App.PageMaxNum // 单次请求的最大数量

	name := c.Param("name")
	random, _ := strconv.ParseBool(c.DefaultQuery("random", "false"))
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))

	if err != nil {
		// 页数不为数字
		c.JSON(http.StatusOK, R.Result(e.PageNotNum))
		return
	}

	if page < 0 {
		page = 1
	}

	songs, _ := service.QuerySongByUser(name)
	songsNum := len(songs)
	resp := make([]song, songsNum)

	if songs == nil {
		c.JSON(http.StatusOK, R.Result(e.SongNotFound))
		return
	}

	for i := 0; i < songsNum; i++ {
		resp[i].SongId = songs[i].ID
		resp[i].Title = songs[i].Title
		resp[i].Album = songs[i].Album
		if songs[i].Artists == nil {
			songs[i].Artists = []model.Artist{}
			continue
		} else {
			resp[i].Artists = make([]artist, len(songs[i].Artists))
			for j := 0; j < len(songs[i].Artists); j++ {
				resp[i].Artists[j].ArtistId = songs[i].Artists[j].ID
				resp[i].Artists[j].Name = songs[i].Artists[j].Name
			}
		}
	}

	if random {
		index := rand.Intn(songsNum)
		c.JSON(http.StatusOK, resp[index])
		return
	}
	maxPage := util.MaxPage(len(songs), maxNum)
	firstNum := (page - 1) * maxNum
	lastNum := page * maxNum

	// 判断访问的页数是否在总页数里面，根据不同情况输出不同数据
	if page < 0 {

	} else if page > maxPage {
		c.JSON(http.StatusOK, R.Result(e.PageOutBound))
	} else if page == maxPage {
		c.JSON(http.StatusOK, R.Success("", resp[firstNum:]))
	} else {
		c.JSON(http.StatusOK, R.Success("", resp[firstNum:lastNum]))
	}
}
