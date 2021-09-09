package e

var MsgFlags = map[int]string{
	Success:          "ok",
	Fail:             "未知错误",
	NotLogin:         "未登录或非法访问",
	NotFound:         "404 Not Found",
	InvalidOperation: "无效操作",

	MissParam:           "缺少参数",
	UsernameFormatError: "用户名格式错误",
	PasswordFormatError: "密码格式错误",
	EmailFormatError:    "邮箱格式错误",

	RepeatUser:  "用户名重复",
	RepeatEmail: "邮箱重复",
	AuthFail:    "账号认证失败",

	PageOutBound: "超过最大页数",
	SongNotFound: "歌曲未找到",
	PageNotNum:   "页码不是数字",
	ParamNotNul:  "参数不能为空",

	RegistrationNotAllowed: "不允许注册",
	MethodError:            "方法错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[Fail]
}
