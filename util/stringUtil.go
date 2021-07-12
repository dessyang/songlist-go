package util

import "strings"

func SearchJson(s string) string {
	s1 := strings.Index(s, "{")
	s2 := strings.LastIndex(s, "}")
	return s[s1 : s2+1]
}
