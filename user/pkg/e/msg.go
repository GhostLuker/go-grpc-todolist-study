package e

var MsgFlags = map[uint]string{
	Sucess:        "ok",
	Error:         "failed",
	InvaildParams: "请求参数错误",
}

//GetMsg 获取状态码的信息
func GetMsg(code uint) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	} else {
		return MsgFlags[Error]
	}
}
