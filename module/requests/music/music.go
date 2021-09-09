package music

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/tidwall/gjson"
	"github.com/yjymh/songlist-go/model"
	"github.com/yjymh/songlist-go/module/requests"
	"github.com/yjymh/songlist-go/util"
	"strings"
)

var (
	baseMusicUrlQQ    = "https://y.qq.com/n/ryqq/songDetail/"
	baseSearchUrlQQ   = "https://c.y.qq.com/soso/fcgi-bin/client_search_cp?w="
	baseMusicUrlKugou = "https://songsearch.kugou.com/song_search_v2?keyword=%s&page=1&pagesize=1&platform=WebFilter"
)

func GetMusicInfoByKugou(title string) (model.Song, error) {
	songInfo := model.Song{}
	url := fmt.Sprintf(baseMusicUrlKugou, title)
	jsonData := requests.GetApiData(url)

	list := jsonData.Get("data.lists")
	if list.Raw == "[]" {
		log.Error("没有找到该歌曲")
		return model.Song{}, error(fmt.Errorf("没有错误"))
	}
	if list.IsArray() {
		for i := range list.Array() {
			songData := list.Array()[i]

			songName := songData.Get("SongName").String()
			if songName != title {
				log.Error("无法获取到相同的歌曲：", songName)
				continue
			}
			songInfo.Title = songName
			songInfo.Album = songData.Get("AlbumName").String()
			songInfo.Time = int(songData.Get("Duration").Int())
			singers := songData.Get("Singers").Array()
			//num := len(singers)
			artists := make([]model.Artist, len(singers))
			for i := 0; i < len(singers); i++ {
				//artist := model.Artists{Name: singers[i].Get("name").String()}
				artists[i].Name = singers[i].Get("name").String()

			}
			songInfo.Artists = artists
		}
	}
	//fmt.Println(jsonData)
	return songInfo, nil
}

// GetMusicInfoByQQ 通过歌曲名，返回歌手信息
func GetMusicInfoByQQ(title string) (model.Song, error) {
	url := baseSearchUrlQQ + strings.Replace(title, " ", "%20", -1)
	jsonData := requests.GetApiData(url)
	list := jsonData.Get("data.song.list")
	//songInfo := new(model.Song)
	songInfo := model.Song{}

	if list.Raw == "[]" {
		log.Error("没有找到该歌曲")
		return model.Song{}, error(fmt.Errorf("没有错误"))
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
			setSongINfo(resp, &songInfo)

			songInfo.SourceUrl = baseMusicUrlQQ + mid

			return songInfo, nil
		}
	}
	return songInfo, error(fmt.Errorf("没有找到歌曲"))
}

// 把该链接下的数据提取出来
func setSongINfo(s string, songInfo *model.Song) {
	data := gjson.Parse(util.SearchJson(s))

	infoData := data.Get("detail")
	singer := infoData.Get("singer")
	if singer.IsArray() {
		var s string
		singers := singer.Array()
		for i := range singers {
			s += singers[i].Get("name").String()
			if i != len(singers)-1 {
				s += ","
			}
		}
		//songInfo.Artists = s // 歌手信息，是个list，可以有多个歌手
	}

	songInfo.Album = infoData.Get("albumName").String() // 专辑名

	pubTime := infoData.Get("info.pub_time.content")
	if pubTime.IsArray() {
		songInfo.PubTime = pubTime.Array()[0].Get("value").String() // 歌曲发行时间
	}
}
