package define

import "os"

var CodeLength int = 6   // 验证码长度
var CodeExpire int = 300 // 验证码过期时间

var SmsService = AlibabService // 使用哪个短信服务商    1:阿里云  2:腾讯云

var AlibabService = 1
var AlibabaAccessKeyId = os.Getenv("AccessKeyId")         // 阿里云权限Id
var AlibabaAccessKeySecret = os.Getenv("AccessKeySecret") // 阿里云权限密码
var AliyunSignName = "孙龙个人网站"                             // 阿里云短信签名
var AliyunTemplateCode = "SMS_270155458"                  // 阿里云短信模板

var TencentService = 2
var TencentSecretId = os.Getenv("TencentSecretId")   // 腾讯云权限id
var TencentSecretKey = os.Getenv("TencentSecretKey") // 腾讯云权限密码
var TencentSmsSdkAppId = "1400796655"                // 腾讯云应用程序ID
var TencentSignName = "vczs公众号"                      // 腾讯云短信签名
var TencentTemplateId = "1706699"                    // 腾讯云短信模板
