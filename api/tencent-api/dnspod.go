package tencent_api

import "fmt"

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
)

func GetAllRootDomain(SecretId string, SecretKey string) (string, error) {
	credential := common.NewCredential(
		SecretId,
		SecretKey,
	)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "dnspod.tencentcloudapi.com"
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := dnspod.NewClient(credential, "", cpf)

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := dnspod.NewDescribeDomainListRequest()

	// 返回的resp是一个DescribeDomainListResponse的实例，与请求对象对应
	response, err := client.DescribeDomainList(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return "", err
	}
	if err != nil {
		panic(err)
	}
	// 输出json格式的字符串回包
	return response.ToJsonString(), nil
}

func GetFQDN(SecretId string, SecretKey string, Domain string) (string, error) {
	credential := common.NewCredential(
		SecretId,
		SecretKey,
	)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "dnspod.tencentcloudapi.com"
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := dnspod.NewClient(credential, "", cpf)

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := dnspod.NewDescribeRecordListRequest()

	request.Domain = common.StringPtr(Domain)

	// 返回的resp是一个DescribeRecordListResponse的实例，与请求对象对应
	response, err := client.DescribeRecordList(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return "", err
	}
	if err != nil {
		return "", err
	}
	// 输出json格式的字符串回包
	return response.ToJsonString(), nil
}

func ChangeFQDNStatus(SecretId string, SecretKey string, Domain string, RecordId uint64, StatusBool bool) (string, error) {
	credential := common.NewCredential(
		SecretId,
		SecretKey,
	)
	Status := ""
	switch StatusBool {
	case true:
		Status = "ENABLE"
	case false:
		Status = "DISABLE"

	}
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "dnspod.tencentcloudapi.com"
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := dnspod.NewClient(credential, "", cpf)

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := dnspod.NewModifyRecordStatusRequest()
	request.Domain = common.StringPtr(Domain)
	request.RecordId = common.Uint64Ptr(RecordId)
	request.Status = common.StringPtr(Status)

	// 返回的resp是一个ModifyRecordStatusResponse的实例，与请求对象对应
	response, err := client.ModifyRecordStatus(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return "", err
	}
	if err != nil {
		return "", err
	}
	// 输出json格式的字符串回包
	return response.ToJsonString(), nil

}

func AddFQDN(SecretId string, SecretKey string, Domain string, RecordType string, RecordLine string, Value string) error {
	credential := common.NewCredential(
		SecretId,
		SecretKey,
	)
	if RecordLine == "" {
		RecordLine = "默认"
	}
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "dnspod.tencentcloudapi.com"
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := dnspod.NewClient(credential, "", cpf)

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := dnspod.NewCreateRecordRequest()

	request.Domain = common.StringPtr(Domain)
	request.RecordType = common.StringPtr(RecordType)
	request.RecordLine = common.StringPtr(RecordLine)
	request.Value = common.StringPtr(Value)

	// 返回的resp是一个CreateRecordResponse的实例，与请求对象对应
	response, err := client.CreateRecord(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return err
	}
	if err != nil {
		return err
	}
	// 输出json格式的字符串回包
	return err

}

func ChangeDomainStatus(SecretId string, SecretKey string, Domain string, StatusBool bool) error {
	credential := common.NewCredential(
		SecretId,
		SecretKey,
	)
	Status := ""
	switch StatusBool {
	case true:
		Status = "ENABLE"
	case false:
		Status = "DISABLE"

	}
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "dnspod.tencentcloudapi.com"
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := dnspod.NewClient(credential, "", cpf)

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := dnspod.NewModifyDomainStatusRequest()

	request.Domain = common.StringPtr(Domain)
	request.Status = common.StringPtr(Status)

	// 返回的resp是一个ModifyDomainStatusResponse的实例，与请求对象对应
	_, err := client.ModifyDomainStatus(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return err
	}
	if err != nil {
		return err
	}
	// 输出json格式的字符串回包
	return err
}
func ChangeFQDN(SecretId string, SecretKey string, Domain string, RecordType string, RecordLine string, Value string, RecordId uint64) error {
	credential := common.NewCredential(
		SecretId,
		SecretKey,
	)
	if RecordLine == "" {
		RecordLine = "默认"
	}
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "dnspod.tencentcloudapi.com"
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := dnspod.NewClient(credential, "", cpf)

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := dnspod.NewModifyRecordRequest()

	request.Domain = common.StringPtr(Domain)
	request.RecordType = common.StringPtr(RecordType)
	request.RecordLine = common.StringPtr(RecordLine)
	request.Value = common.StringPtr(Value)
	request.RecordId = common.Uint64Ptr(RecordId)

	// 返回的resp是一个ModifyRecordResponse的实例，与请求对象对应
	response, err := client.ModifyRecord(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return err
	}
	if err != nil {
		return err
	}
	// 输出json格式的字符串回包
	return err
}

func DeleteFQDN(SecretId string, SecretKey string, Domain string, RecordId uint64) error {
	credential := common.NewCredential(
		SecretId,
		SecretKey,
	)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "dnspod.tencentcloudapi.com"
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := dnspod.NewClient(credential, "", cpf)

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := dnspod.NewDeleteRecordRequest()

	request.Domain = common.StringPtr(Domain)
	request.RecordId = common.Uint64Ptr(RecordId)

	// 返回的resp是一个DeleteRecordResponse的实例，与请求对象对应
	response, err := client.DeleteRecord(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return err
	}
	if err != nil {
		return err
	}
	// 输出json格式的字符串回包
	return err
}
