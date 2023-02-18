package controller

import (
	"net/http"
	"sms-code/define"

	"github.com/gin-gonic/gin"
)

type Body struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Res(c *gin.Context, code int, msg string) {
	body := Body{Code: code, Message: msg}
	if msg == "" {
		body.Message = define.Message(code)
	}
	c.JSON(http.StatusOK, body)
}

type BodyData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func ResData(c *gin.Context, code int, msg string, data any) {
	body := BodyData{Code: code, Message: msg, Data: data}
	if msg == "" {
		body.Message = define.Message(code)
	}
	c.JSON(http.StatusOK, body)
}
