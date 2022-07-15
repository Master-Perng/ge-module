package github

import (
	"fmt"
	logsys "github.com/Master-Perng/go-module/log"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const githubapi = "https://api.github.com/search/code?per_page=100&q=%s"

func SearchGithub(keyword string, token string) (string, error) {
	language := "  language:C# or language:java or language:php or language:go or language:jsp"
	//创建client结果，里面是http连接的参数 ， 比如超时 https策略 代理等等
	client := &http.Client{
		Timeout: 15 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	reqUrl := fmt.Sprintf(githubapi, url.PathEscape(keyword+language))
	req, err := http.NewRequest("GET", reqUrl, strings.NewReader(""))
	req.Header.Add("Authorization", "token "+token)
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	if err != nil {
		logsys.Error(err.Error())
		return "", err
	}
	resp, err := client.Do(req)
	if err != nil {
		if strings.Contains(err.Error(), "Timeout") {
			i := 0
			for {
				time.Sleep(2 * time.Second)
				resp, err = client.Do(req)
				i++
				if err == nil {
					break
				}
				if i > 5 {
					defer client.CloseIdleConnections()
					return "", err
				}
			}
		} else {
			defer client.CloseIdleConnections()
			logsys.Error(err.Error())
			return "", err
		}
	}
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		defer client.CloseIdleConnections()
		logsys.Error(err.Error())
		return "", err
	}
	defer client.CloseIdleConnections()
	return string(result), err

}
