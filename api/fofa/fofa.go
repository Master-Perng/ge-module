package fofa

import (
	"encoding/base64"
	"fmt"
	logsys "github.com/Master-Perng/go-module/log"
	"io"
	"net/http"
	"strings"
	"time"
)

func Query(Email string, key string, query string, fields string) (string, error) {
	const api = "https://fofa.info/api/v1/search/all?email=%s&key=%s&qbase64=%s&size=2000&fields=%s"
	client := &http.Client{
		Timeout: 15 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	url := fmt.Sprintf(api, Email, key, base64.StdEncoding.EncodeToString([]byte(query)), fields)
	req, err := http.NewRequest("GET", url, strings.NewReader(""))
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
		logsys.Error(err.Error())
		return "", err
	}
	return string(result), err
}
