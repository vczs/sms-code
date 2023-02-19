package define

import "os"

var CodeLength int = 6   // 验证码长度
var CodeExpire int = 300 // 验证码过期时间

var AlibabaAccessKeyId = os.Getenv("AccessKeyId")         // 阿里云keyId
var AlibabaAccessKeySecret = os.Getenv("AccessKeySecret") // 阿里云keySecret

var AliyunSignName = "孙龙个人网站"            // 阿里云短信签名
var AliyunTemplateCode = "SMS_270155458" // 阿里云模板
