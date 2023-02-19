package controller

import (
	"sms-code/db"
	"sms-code/define"
	"sms-code/help"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type CheckCodeRequest struct {
	PhoneNumbers string `json:"phoneNumbers"`
	Code         string `json:"code"`
}

func CheckCode(c *gin.Context) {
	// 参数解析
	req := new(CheckCodeRequest)
	err := c.ShouldBind(&req)
	if err != nil {
		help.VczLog("parameter analysis failed", err)
		Res(c, define.PARAMETER_FAILED, "")
		return
	}

	// 参数非空检查
	phoneNumbers := req.PhoneNumbers
	code := req.Code
	if phoneNumbers == "" || code == "" {
		Res(c, define.PARAMETER_WRONG, "")
		return
	}

	// 获取验证码
	val, err := db.Redis.Get(c, phoneNumbers).Result()
	if err != nil {
		if err == redis.Nil {
			Res(c, define.PHONE_NUMBERS_EMPTY, "")
			return
		} else {
			help.VczLog("get code from redis error", err)
			Res(c, -1, err.Error())
			return
		}
	}

	// 检查验证码
	if code != val {
		Res(c, define.CODE_WRONG, "")
		return
	}

	Res(c, define.OK, "验证码正确!")
}
