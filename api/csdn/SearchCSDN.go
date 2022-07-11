package csdn

import (
	"fmt"
	logsys "github.com/Master-Perng/go-module/log"
	jsoniter "github.com/json-iterator/go"
	"io"
	"net/http"
	"strings"
	"time"
)

const csdn = "https://so.csdn.net/api/v3/search?q=%s&t=code&p=%d"

func SearchCSDN(keyword string) (string, error) {
	//创建client结果，里面是http连接的参数 ， 比如超时 https策略 代理等等
	client := &http.Client{
		Timeout: 15 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	reqUrl := fmt.Sprintf(csdn, keyword, 1)
	req, err := http.NewRequest("GET", reqUrl, strings.NewReader(""))
	if err != nil {
		defer client.CloseIdleConnections()
		return "", err
	}
	resp, err := client.Do(req)
	if err != nil {
		defer client.CloseIdleConnections()
		return "", err
	}
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		defer client.CloseIdleConnections()
		return "", err
	}
	defer client.CloseIdleConnections()
	page := jsoniter.Get(result, "total_page").ToInt()
	if page > 1 {
		reqUrl := fmt.Sprintf(csdn, keyword, page)
		req, err = http.NewRequest("GET", reqUrl, strings.NewReader(""))
		resp, err = client.Do(req)
		resp, err = client.Do(req)
		result, err = io.ReadAll(resp.Body)
	}
	if err != nil {
		logsys.Error(err.Error())
		return "", err
	}
	return string(result), nil

}
