package main

import (
	"fmt"
	"sms-code/config"
	"sms-code/db"
	"sms-code/router"
)

func main() {
	config.InitConfig()
	db.Init()
	engine := router.Router()
	engine.Run(fmt.Sprintf(":%d", config.GetConfig().Port))
}
