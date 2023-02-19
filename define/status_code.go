package define

const (
	OK = 0 // success

	REQUEST_OFTEN    = 101 // 请求频繁
	PARAMETER_FAILED = 102 // 参数解析失败
	PARAMETER_WRONG  = 103 // 参数有误

	CODE_SEND_FAILED    = 20001 // 验证码发送失败
	CODE_WRONG          = 20002 // 验证码错误
	PHONE_NUMBERS_EMPTY = 20003 // 请先获取验证码
)

var message = map[int]string{
	OK: "success",

	REQUEST_OFTEN:    "请求频繁!",
	PARAMETER_FAILED: "参数解析失败!",
	PARAMETER_WRONG:  "参数有误!",

	CODE_SEND_FAILED:    "验证码发送失败!",
	CODE_WRONG:          "验证码错误!",
	PHONE_NUMBERS_EMPTY: "请先获取验证码!",
}

// GetMessage 获取message
func Message(code int) string {
	if msg, ok := message[code]; ok {
		return msg
	} else {
		return "服务器发生未知错误~"
	}
}
