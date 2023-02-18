package db

import (
	"fmt"
	"sms-code/config"

	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client

func Init() {
	InitRedis(config.GetConfig().Redis.Host, config.GetConfig().Redis.Port)
}

func InitRedis(host string, port int) {
	Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
