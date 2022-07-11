package amap

import (
	"fmt"
	logsys "github.com/Master-Perng/go-module/log"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"
)

const api = "https://restapi.amap.com/v3/config/district?keywords=%s&subdistrict=%d&key=%s"

func Districts(keyword string, subdistrict string, key string) ([]string, error) {
	client := &http.Client{
		Timeout: 15 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	AmapUrl := fmt.Sprintf(api, keyword, subdistrict, key)
	req, err := http.NewRequest("GET", AmapUrl, strings.NewReader(""))
	if err != nil {
		defer client.CloseIdleConnections()
		logsys.Error(err.Error())
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		defer client.CloseIdleConnections()
		logsys.Error(err.Error())
		return nil, err
	}
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		defer client.CloseIdleConnections()
		logsys.Error(err.Error())
		return nil, err
	}
	re := regexp.MustCompile(`"name":"(.+?)"`)
	result1 := re.FindAllStringSubmatch(string(result), -1)
	sub := make([]string, len(result1))
	for i := range result1 {
		sub[i] = result1[i][1]
	}
	defer client.CloseIdleConnections()
	return sub, err
}
