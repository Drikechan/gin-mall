package e

var MsgFlags = map[int]string{
	SUCCESS:            "请求成功",
	ERROR:              "请求失败",
	InvalidParams:      "错误请求",
	ErrorAuthCheckFail: "Token验证错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
