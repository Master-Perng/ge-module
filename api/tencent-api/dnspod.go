package tencent_api

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
)

func GetAllDns(SecretId string, SecretKey string) (string, error) {

	credential := common.NewCredential(
		SecretId,
		SecretKey,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "dnspod.tencentcloudapi.com"
	client, _ := dnspod.NewClient(credential, "", cpf)

	request := dnspod.NewDescribeDomainListRequest()

	request.Type = common.StringPtr("ALL")

	response, err := client.DescribeDomainList(request)
	return response.ToJsonString(), err
}
