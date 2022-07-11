package Alibaba

import (
	log "engine/logsys"
	dysmsapi "github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

//Message函数使用阿里云sdk
func Message(phone string, message string) string {
	client, err := dysmsapi.NewClientWithAccessKey("cn-shanghai", "LTAI5tEaMFHzJ7ZbVHJHdaCG", "k12myDRsOqM6AJBrmHu9a9IQ1zrVmS")
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = phone
	request.SignName = "白菜帽子的心路历程"
	request.TemplateCode = "SMS_225394548"
	request.TemplateParam = "{\"code\":\"" + message + "\"}"
	request.SmsUpExtendCode = "123321"
	request.OutId = "abcdefg"
	response, err := client.SendSms(request)
	if err != nil {
		defer client.CloseIdleConnections()
		log.Error(err.Error())
	}
	return response.Message

}
