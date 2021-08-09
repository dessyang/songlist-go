package song_service

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

var cache = make(chan string, 10)
var file = new([]string)

func SET() {
	var title string
	// 需要导入的文本
	fmt.Println(filepath.Abs("../../song.txt"))
	f, err := os.Open("../../song.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		title = scanner.Text()
		cache <- title
	}
	close(cache)
}

func ImportMusic() {
	for {
		title, ok := <-cache
		if !ok {
			return
		}
		fmt.Println(title)
	}
}
