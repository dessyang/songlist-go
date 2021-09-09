package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxPage(t *testing.T) {
	var sum int
	a := assert.New(t)
	page := 10

	sum = 4
	a.EqualValues(1, MaxPage(sum, page))
	sum = 0
	a.EqualValues(1, MaxPage(sum, page))
	sum = -7
	a.EqualValues(1, MaxPage(sum, page))
	sum = 10
	a.EqualValues(1, MaxPage(sum, page))
	sum = 11
	a.EqualValues(2, MaxPage(sum, page))
}

func TestIsEmail(t *testing.T) {
	var s string
	a := assert.New(t)

	s = ""
	a.EqualValues(false, IsEmail(s))
	s = "814123650@qq.com"
	a.EqualValues(true, IsEmail(s))
	s = " "
	a.EqualValues(false, IsEmail(s))
	s = "8@1.4"
	a.EqualValues(true, IsEmail(s))
}

func TestIsUsername(t *testing.T) {
	var s string
	a := assert.New(t)

	s = ""
	a.EqualValues(false, IsUsername(s))
	s = "1234"
	a.EqualValues(true, IsUsername(s))
	s = "qwer"
	a.EqualValues(true, IsUsername(s))
	s = "a 1 "
	a.EqualValues(false, IsUsername(s))
	s = "1234567890qwerty"
	a.EqualValues(true, IsUsername(s))
	s = "1234567890qwertyu"
	a.EqualValues(false, IsUsername(s))
	s = "a-1_"
	a.EqualValues(true, IsUsername(s))
}

func TestIsPassword(t *testing.T) {
	var s string
	a := assert.New(t)

	s = ""
	a.EqualValues(false, IsPassword(s))
	s = "12345678"
	a.EqualValues(true, IsPassword(s))
	s = "123456@"
	a.EqualValues(false, IsPassword(s))
	s = "12345"
	a.EqualValues(false, IsPassword(s))
	s = "AsdQwe"
	a.EqualValues(true, IsPassword(s))
}
