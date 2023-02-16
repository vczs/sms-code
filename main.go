package main

import "sms-code/router"

func main() {
	engine := router.Router()
	engine.Run(":8080")
}
