package help

import (
	"math/rand"
	"sms-code/define"
	"time"
)

func GenerateCode() string {
	s := "1234567890"
	code := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < define.CodeLength; i++ {
		code += string(s[rand.Intn(10)])
	}
	return code
}
