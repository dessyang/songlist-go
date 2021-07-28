package model

import "github.com/yjymh/songlist-go/pkg/e"

type R struct {
	Code int
	Data interface{}
	Msg  string
}

func (r R) Success(data interface{}) R {
	r.Code = e.Success
	r.Data = data
	return r
}

func (r R) Fail(code int) R {
	r.Code = code
	r.Data = ""
	r.Msg = e.GetMsg(code)
	return r
}
