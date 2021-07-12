package main

import (
	"bufio"
	"github.com/yjymh/songlist-go/conf"
	"github.com/yjymh/songlist-go/service/song_service"
	"os"
)

func init() {
	conf.InitConfig("")
}

func main() {
	var title string
	// 需要导入的文本
	f, err := os.Open("song.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		title = scanner.Text()
		song_service.AddSongInfo(title)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
