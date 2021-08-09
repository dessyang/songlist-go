package main

import (
	"bufio"
	"fmt"
	"github.com/yjymh/songlist-go/conf"
	"github.com/yjymh/songlist-go/service/song_service"
	"os"
)

func init() {
	conf.Setup("")
}

func main() {
	var title string
	var faile []string
	var flag bool
	// 需要导入的文本
	f, err := os.Open("song.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		title = scanner.Text()
		flag = song_service.AddSongInfo(title)
		if !flag {
			faile = append(faile, title)
		}
	}
	fmt.Println(faile)
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
