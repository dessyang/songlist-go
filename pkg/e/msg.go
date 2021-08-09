package e

var MsgFlags = map[int]string{
	Success:         "ok",
	Fail:            "未知错误",
	NotLogin:        "未登录或非法访问",
	JwtTokenTimeout: "token已过期",
	JwtTokenFail:    "token错误",
	PageOutBound:    "超过最大页数",
	AuthFail:        "账号或密码错误",
	SongNotFound:    "歌曲未找到",
	PageNotNum:      "页码不是数字",
	ParamNotNul:     "参数不能为空",
	RepeatUser:      "用户名重复",
	RepeatEmail:     "邮箱重复",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[Fail]
}
