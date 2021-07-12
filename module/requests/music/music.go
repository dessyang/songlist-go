package music

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/tidwall/gjson"
	"github.com/yjymh/songlist-go/model/song"
	"github.com/yjymh/songlist-go/module/requests"
	"github.com/yjymh/songlist-go/util"
	"strings"
)

var (
	baseMusicUrlQQ  = "https://y.qq.com/n/ryqq/songDetail/"
	baseSearchUrlQQ = "https://c.y.qq.com/soso/fcgi-bin/client_search_cp?w="
)

// GetMusicInfoByQQ 通过歌曲名，返回歌手信息
func GetMusicInfoByQQ(title string) (*song.Info, error) {
	url := baseSearchUrlQQ + strings.Replace(title, " ", "%20", -1)
	jsonData := requests.GetApiData(url)
	list := jsonData.Get("data.song.list")
	songInfo := new(song.Info)

	if list.Raw == "[]" {
		log.Error("没有找到该歌曲")
		return nil, error(fmt.Errorf("没有错误"))
	}
	if list.IsArray() {
		for i := range list.Array() {
			songData := list.Array()[i]

			mid := songData.Get("songmid").String()
			songName := songData.Get("songname").String() // 歌名
			songTime := songData.Get("interval").Int()    // 时长
			if title != songName {
				log.Error("无法获取到相同的歌曲：", songName)
				continue
			}

			songInfo.Title = title
			songInfo.Time = int(songTime)

			resp := requests.Fetch(baseMusicUrlQQ + mid)
			setSongINfo(resp, songInfo)

			songInfo.SourceUrl = baseMusicUrlQQ + mid

			return songInfo, nil
		}
	}
	return nil, error(fmt.Errorf("没有找到歌曲"))
}

// 把该链接下的数据提取出来
func setSongINfo(s string, songInfo *song.Info) {
	data := gjson.Parse(util.SearchJson(s))

	infoData := data.Get("detail")

	songInfo.AlbumPic = "https:" + infoData.Get("picurl").String() // 专辑或者歌曲图片

	singer := infoData.Get("singer")
	if singer.IsArray() {
		var s string
		singers := singer.Array()
		for i := range singers {
			singers[i].Get("name").String()
			if i != len(singers)-1 {
				s += ","
			}
		}
		songInfo.Artist = s // 歌手信息，是个list，可以有多个歌手
	}

	songInfo.Album = infoData.Get("albumName").String() // 专辑名

	company := infoData.Get("info.company.content")
	if company.IsArray() {
		songInfo.Company = company.Array()[0].Get("value").String() // 唱片公司 有时会空
	}

	genre := infoData.Get("info.genre.content")
	if genre.IsArray() {
		songInfo.Genre = genre.Array()[0].Get("value").String() // 歌曲流派 也可能会空
	}

	songInfo.Lang = infoData.Get("info.lan.content").Array()[0].Get("value").String() // 语种

	pubTime := infoData.Get("info.pub_time.content")
	if pubTime.IsArray() {
		songInfo.PubTime = pubTime.Array()[0].Get("value").String() // 歌曲发行时间
	}
}
