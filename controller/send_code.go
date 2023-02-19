package controller

import (
	"sms-code/db"
	"sms-code/define"
	"sms-code/help"
	"time"

	"github.com/gin-gonic/gin"
)

type SendCodeRequest struct {
	PhoneNumbers string `json:"phoneNumbers"`
}

func SendCode(c *gin.Context) {
	// 参数解析
	req := new(SendCodeRequest)
	err := c.ShouldBind(&req)
	if err != nil {
		help.VczLog("parameter analysis failed", err)
		Res(c, define.PARAMETER_FAILED, "")
		return
	}

	// 参数非空检查
	phoneNumbers := req.PhoneNumbers
	if phoneNumbers == "" {
		Res(c, define.PARAMETER_WRONG, "")
		return
	}

	// 参数校验
	if phoneNumbers == "" {
		Res(c, define.PARAMETER_WRONG, "")
		return
	} else {
		check := help.CheckNumber(phoneNumbers)
		if !check {
			Res(c, define.PARAMETER_WRONG, "手机号码格式错误！")
			return
		}
	}
	// 限频
	down, err := db.Redis.TTL(c, phoneNumbers).Result()
	if err != nil {
		help.VczLog("get number ttl from redis error", err)
		Res(c, -1, err.Error())
		return
	}
	if down > 0 {
		Res(c, define.REQUEST_OFTEN, "")
		return
	}
	// 生成验证码
	code := help.GenerateCode()
	// 发送验证码
	err = help.AlibabaSendSmsCode(phoneNumbers, code)
	if err != nil {
		help.VczLog("get number ttl from redis error", err)
		Res(c, define.CODE_SEND_FAILED, "")
		return
	}
	// 存储验证码
	err = db.Redis.Set(c, phoneNumbers, code, time.Second*time.Duration(define.CodeExpire)).Err()
	if err != nil {
		help.VczLog("set code to redis error", err)
		Res(c, -1, err.Error())
		return
	}

	Res(c, define.OK, "发送成功")
}
