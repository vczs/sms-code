package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

var c Config

type Config struct {
	Port  int
	Redis struct {
		Host string
		Port int
	}
}

func InitConfig() {
	c = Config{}
	data, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		log.Fatalln("read ymal file failed", err)
	}
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		log.Fatalln("yaml unmarshal failed", err)
	}
}

func GetConfig() *Config {
	return &c
}
