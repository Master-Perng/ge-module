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
		Timeout: 30 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	AmapUrl := fmt.Sprintf(api, keyword, subdistrict, key)
	req, err := http.NewRequest("GET", AmapUrl, strings.NewReader(""))
	if err != nil {
		logsys.Error(err.Error())
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		logsys.Error(err.Error())
		return nil, err
	}
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		logsys.Error(err.Error())
		return nil, err
	}
	re := regexp.MustCompile(`"name":"(.+?)"`)
	result1 := re.FindAllStringSubmatch(string(result), -1)
	fmt.Println(result1)
	keylist := make([]string, len(result1))
	return keylist, err
}
