package help

import (
	"encoding/json"
	"errors"
	"sms-code/define"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

func AlibabaSendSmsCode(phoneNumbers string, code string) error {
	config := &openapi.Config{
		AccessKeyId:     tea.String(define.AlibabaAccessKeyId),
		AccessKeySecret: tea.String(define.AlibabaAccessKeySecret),
		Endpoint:        tea.String("dysmsapi.aliyuncs.com"),
	}
	client, err := dysmsapi20170525.NewClient(config)
	if err != nil {
		return err
	}

	param, err := json.Marshal(map[string]any{"code": code})
	if err != nil {
		return err
	}
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  tea.String(phoneNumbers),
		SignName:      tea.String(define.AliyunSignName),
		TemplateCode:  tea.String(define.AliyunTemplateCode),
		TemplateParam: tea.String(string(param)),
	}

	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		result, _err := client.SendSmsWithOptions(sendSmsRequest, runtime)
		if _err != nil {
			return _err
		}
		if *result.Body.Code != "OK" {
			return errors.New(*result.Body.Message)
		}
		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		result, _err := util.AssertAsString(error.Message)
		if _err != nil {
			return _err
		}
		return errors.New(*result)
	}

	return nil
}
