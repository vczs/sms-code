package router

import (
	"sms-code/controller"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.POST("/send_code", controller.SendCode)
	r.POST("/check_code", controller.CheckCode)

	return r
}
