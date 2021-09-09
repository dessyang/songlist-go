package model

import (
	"fmt"
	"github.com/yjymh/songlist-go/pkg/e"
)

type R struct {
	Code int
	Data interface{}
	Msg  string
}

func (r R) Success(data ...interface{}) R {
	r.Code = e.Success
	r.Msg = fmt.Sprint(data[0])
	if len(data) == 1 {
		r.Data = ""
	} else {
		r.Data = data[1]
	}
	return r
}

func (r R) Fail(msg string) R {
	r.Code = e.Fail
	r.Data = ""
	r.Msg = msg
	return r
}

func (r R) Result(code int) R {
	r.Code = code
	r.Data = ""
	r.Msg = e.GetMsg(code)
	return r
}
