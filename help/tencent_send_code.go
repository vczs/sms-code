package help

import (
	"errors"
	"fmt"
	"sms-code/define"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

func TencentSendSmsCode(phoneNumbers string, code string) error {
	credential := common.NewCredential(define.TencentSecretId, define.TencentSecretKey)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
	client, _ := sms.NewClient(credential, "ap-guangzhou", cpf)

	request := sms.NewSendSmsRequest()
	request.SmsSdkAppId = common.StringPtr(define.TencentSmsSdkAppId)
	request.SignName = common.StringPtr(define.TencentSignName)
	request.TemplateId = common.StringPtr(define.TencentTemplateId)
	request.PhoneNumberSet = common.StringPtrs([]string{phoneNumbers})
	request.TemplateParamSet = common.StringPtrs([]string{code})

	response, err := client.SendSms(request)
	if err != nil {
		return err
	}
	if *response.Response.SendStatusSet[0].Code != "Ok" {
		fmt.Println(*response.Response.SendStatusSet[0].Code)
		return errors.New(*response.Response.SendStatusSet[0].Message)
	}

	return nil
}
