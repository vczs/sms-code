package router

import (
	"sms-code/controller"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	engine := gin.Default()

	engine.POST("/send_code", controller.SendCode)

	return engine
}
