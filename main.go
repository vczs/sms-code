package main

import (
	"fmt"
	"sms-code/config"
	"sms-code/router"
)

func main() {
	config.InitConfig()
	engine := router.Router()
	engine.Run(fmt.Sprintf(":%d", config.GetConfig().Port))
}
