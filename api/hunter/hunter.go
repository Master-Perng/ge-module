package hunter

import (
	"encoding/base64"
	"engine/logsys"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func Query(Username string, key string, query string) (string, error) {
	const api = "https://hunter.qianxin.com/openApi/search?username=%s&api-key=%s&search=%s&page=1&page_size=100&is_web=3&start_time=\"%d-01-01+00:00:00\"&end_time=\"%d-12-31+00:00:00\""
	client := &http.Client{
		Timeout: 5 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	url := fmt.Sprintf(api, Username, key, base64.StdEncoding.EncodeToString([]byte(query)), time.Now().Year()-1, time.Now().Year())
	req, err := http.NewRequest("GET", url, strings.NewReader(""))
	if err != nil {
		logsys.Error(err.Error())
		return "", err
	}
	resp, err := client.Do(req)
	if err != nil {
		logsys.Error(err.Error())
		return "", err
	}
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		logsys.Error(err.Error())
		return "", err
	}
	return string(result), err
}
