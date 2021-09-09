package util

import (
	"regexp"
	"strings"
)

func SearchJson(s string) string {
	s1 := strings.Index(s, "{")
	s2 := strings.LastIndex(s, "}")
	return s[s1 : s2+1]
}

func IsEmpty(s string) bool {
	if s == "" {
		return true
	}
	return false
}

func IsNotEmpty(s string) bool {
	return !IsEmpty(s)
}

func IsEmail(s string) bool {
	s1 := "^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$"
	return customRegexp(s, s1)
}

func IsPassword(s string) bool {
	s1 := `^[a-zA-Z0-9]{6,16}$`
	return customRegexp(s, s1)
}

func IsUsername(s string) bool {
	s1 := "^[a-zA-Z0-9_-]{4,16}$"
	return customRegexp(s, s1)
}

func customRegexp(s string, s1 string) bool {
	re := regexp.MustCompile(s1)
	a := re.FindString(s)
	if a == s && IsNotEmpty(a) {
		return true
	}
	return false
}
