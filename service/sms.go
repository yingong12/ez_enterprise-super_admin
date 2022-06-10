// This file is auto-generated, don't edit it. Thanks.
package service

import (
	"encoding/json"
	"super_admin/library/env"
	"super_admin/repository"
	"super_admin/utils"

	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 域名有问题需要认真
	config.Endpoint = tea.String(env.GetStringVal("SERVICE_DOMAIN"))
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

type TemplateParam struct {
	Code string `json:"code"`
}

func send(phone, code string) (_err error) {
	client, _err := CreateClient(tea.String("SMS_API_KEY"), tea.String("SMS_API_KEY_SECRET"))
	if _err != nil {
		return _err
	}
	tp := TemplateParam{
		Code: code,
	}
	j, _err := json.Marshal(tp)
	if _err != nil {
		return
	}
	tpStr := (string)(j)
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		SignName:      tea.String(env.GetStringVal("COMPANY_NAME")),
		TemplateCode:  tea.String(env.GetStringVal("SMS_TEMPLATE")),
		PhoneNumbers:  tea.String(phone),
		TemplateParam: tea.String(tpStr),
	}
	// 复制代码运行请自行打印 API 的返回值
	_, _err = client.SendSms(sendSmsRequest)
	if _err != nil {
		return _err
	}
	return _err
}

func SendVerifyCode(phone string) (ok bool, err error) {
	code := utils.GenerateVerifyCode()
	//这里先写redis，因为如果失败了就不发短信。不会造成用户收到短信结果没法使用的bug
	ok, err = repository.SetSMSEntry(env.GetStringVal("KEY_PREFIX_SMS")+phone, code)
	if err != nil {
		return
	}
	// err = send(phone, code)
	return
}
