package v1

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/yjymh/songlist-go/middleware"
	"github.com/yjymh/songlist-go/model"
	"github.com/yjymh/songlist-go/module/requests/music"
	"github.com/yjymh/songlist-go/pkg/e"
	"github.com/yjymh/songlist-go/service"
	"net/http"
	"strconv"
	"time"
)

// Upload 上传需要导入的文件
func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(500, "上传图片出错")
	}
	c.SaveUploadedFile(file, file.Filename)
}

//ImportMusicByTxt 导入歌曲
//@Router /api/v1/import [get]
//TODO 文件的读写
func ImportMusicByTxt(c *gin.Context) {
	R := model.R{}
	session := sessions.Default(c)
	auth := (session.Get("auth")).(model.Auth)

	fileName := auth.Username + "-" + strconv.Itoa(int(time.Now().Unix())) + ".txt"

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, R.Fail("上传文件错误"))
	}
	c.SaveUploadedFile(file, fileName)
}

// ImportMusicByName 通过歌曲名导入歌曲
func ImportMusicByName(c *gin.Context) {
	title := c.PostForm("title")
	song, err := music.GetMusicInfoByKugou(title)
	if err != nil {
		middleware.Logger().Errorln(err)
	} else {
		session := sessions.Default(c)
		auth := (session.Get("auth")).(model.Auth)

		flag, err := service.AddSongInfo(&song, &auth)
		if err != nil {
			middleware.Logger().Errorln(err)
		}
		if flag {
			c.JSON(http.StatusOK, model.R{}.Success("添加成功"))
			return
		}
	}
	c.JSON(http.StatusOK, model.R{}.Result(e.Fail))
}

func ImportMusicByExcel(c *gin.Context) {

}

// ImportMusic api/v1/import/song?method="file"
func ImportMusic(c *gin.Context) {
	method := c.Query("method")
	switch method {
	case "txt":
		ImportMusicByTxt(c)
	case "name":
		ImportMusicByName(c)
	case "excel":
		ImportMusicByExcel(c)
	default:
		c.JSON(http.StatusOK, model.R{}.Result(e.MethodError))
	}
	return
}
