package controller

import (
	"fmt"
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
	count := 0
	countDown, err := db.Redis.TTL(c, define.RedisCodeCountTitle+phoneNumbers).Result()
	if err != nil {
		help.VczLog("get number code count ttl from redis error", err)
		Res(c, -1, err.Error())
		return
	}
	if countDown > 0 {
		count, err = db.Redis.Get(c, define.RedisCodeCountTitle+phoneNumbers).Int()
		if err != nil {
			help.VczLog("get number code count from redis error", err)
			Res(c, -1, err.Error())
			return
		}
		if count >= define.MaxCodeCount {
			Res(c, define.MANX_CODE_COUNT, fmt.Sprintf("今日获取短信数量上限，请%s后再来!", countDown))
			return
		}

		codeDown, err := db.Redis.TTL(c, phoneNumbers).Result()
		if err != nil {
			help.VczLog("get number code ttl from redis error", err)
			Res(c, -1, err.Error())
			return
		}
		if codeDown > 0 && count%2 == 0 {
			Res(c, define.REQUEST_OFTEN, "")
			return
		}
	}

	// 发送验证码
	code := help.GenerateCode()
	if define.SmsService == define.AlibabService {
		err = help.AlibabaSendSmsCode(phoneNumbers, code)
		if err != nil {
			help.VczLog("alibab send sms code failed", err)
			Res(c, define.CODE_SEND_FAILED, "")
			return
		}
	} else {
		err = help.TencentSendSmsCode(phoneNumbers, code)
		if err != nil {
			help.VczLog("tencent send sms code failed", err)
			Res(c, define.CODE_SEND_FAILED, "")
			return
		}
	}

	// 存储验证码
	err = db.Redis.Set(c, phoneNumbers, code, time.Second*time.Duration(define.CodeExpire)).Err()
	if err != nil {
		help.VczLog("set code to redis error", err)
		Res(c, -1, err.Error())
		return
	}

	// 存储验证码发送次数
	if count > 0 {
		err = db.Redis.Set(c, define.RedisCodeCountTitle+phoneNumbers, count+1, -1).Err()
		if err != nil {
			help.VczLog("set code count to redis error", err)
			Res(c, -1, err.Error())
			return
		}
	} else {
		err = db.Redis.Set(c, define.RedisCodeCountTitle+phoneNumbers, 1, time.Hour*24).Err()
		if err != nil {
			help.VczLog("set code count to redis error", err)
			Res(c, -1, err.Error())
			return
		}
	}

	Res(c, define.OK, "发送成功")
}
