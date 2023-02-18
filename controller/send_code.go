package controller

import (
	"sms-code/db"
	"sms-code/define"
	"sms-code/help"
	"time"

	"github.com/gin-gonic/gin"
)

type SendCodeRequest struct {
	Number string `json:"number"`
}

func SendCode(c *gin.Context) {
	// 参数解析
	req := new(SendCodeRequest)
	err := c.ShouldBind(&req)
	if err != nil {
		Res(c, define.PARAMETER_FAILED, "")
		return
	}
	number := req.Number

	// 参数校验
	if number == "" {
		Res(c, define.PARAMETER_WRONG, "")
		return
	} else {
		check := help.CheckNumber(number)
		if !check {
			Res(c, define.PARAMETER_WRONG, "手机号码格式错误！")
			return
		}
	}
	// 限频
	down, err := db.Redis.TTL(c, number).Result()
	if err != nil {
		help.VczLog("get number ttl from redis error", err)
		return
	}
	if down > 0 {
		Res(c, define.REQUEST_OFTEN, "")
		return
	}
	// 生成验证码
	code := help.GenerateCode()
	// 发送验证码

	// 存储验证码
	err = db.Redis.Set(c, number, code, time.Second*time.Duration(define.CodeExpire)).Err()
	if err != nil {
		help.VczLog("set code to redis error", err)
		return
	}

	Res(c, define.OK, "发送成功")
}
