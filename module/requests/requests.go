package requests

import (
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"strings"
)

func Fetch(url string) string {
	resp, err := http.Get(url)
	if err != nil {
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	}
	resp.Body.Close()
	return string(data)
}

func GetApiData(url string) gjson.Result {
	data := Fetch(url)
	// 去除头尾
	data = strings.Trim(string(data), "callback(")
	data = strings.Trim(data, ")")

	return gjson.Parse(data)
}
