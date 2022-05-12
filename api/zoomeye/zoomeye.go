package zoomeye

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func Query(Username string, key string, query string) (string, error) {
	const api = "https://api.zoomeye.org/host/search?query=port:21%20city:beijing&page=1&facets=app,os"
	client := &http.Client{
		Timeout: 5 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	url := fmt.Sprintf(api, Username, key, base64.StdEncoding.EncodeToString([]byte(query)), time.Now().Year()-1, time.Now().Year())
	req, err := http.NewRequest("GET", url, strings.NewReader(""))
	if err != nil {
		//logsys.Error(err.Error())
		return "", err
	}
	resp, err := client.Do(req)
	if err != nil {
		//logsys.Error(err.Error())
		return "", err
	}
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		//logsys.Error(err.Error())
		return "", err
	}
	return string(result), err
}
