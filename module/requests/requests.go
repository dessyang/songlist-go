package requests

import (
	"github.com/imroc/req"
	"github.com/tidwall/gjson"
	"strings"
)

func Fetch(url string) string {
	r, _ := req.Get(url)
	return r.String()
}

func GetApiData(url string) gjson.Result {

	data := Fetch(url)
	// 去除头尾
	data = strings.Trim(string(data), "callback(")
	data = strings.Trim(data, ")")

	return gjson.Parse(data)
}
