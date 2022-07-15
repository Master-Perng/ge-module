package hunter

import (
	"encoding/base64"
	"fmt"
	logsys "github.com/Master-Perng/go-module/log"
	"io"
	"net/http"
	"strings"
	"time"
)

func Query(Username string, page int, key string, query string) (string, error) {
	const api = "https://hunter.qianxin.com/openApi/search?username=%s&api-key=%s&search=%s&page=%d&page_size=100&is_web=3&start_time=\"%d-01-01+00:00:00\"&end_time=\"%d-12-31+00:00:00\""
	client := &http.Client{
		Timeout: 5 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	url := fmt.Sprintf(api, Username, key, base64.StdEncoding.EncodeToString([]byte(query)), page, time.Now().Year()-1, time.Now().Year())
	req, err := http.NewRequest("GET", url, strings.NewReader(""))
	if err != nil {
		defer client.CloseIdleConnections()
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
