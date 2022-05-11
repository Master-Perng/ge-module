package fofa

import (
	"encoding/base64"
	"engine/logsys"
	"fmt"
	"net/http"
	"strings"
)

func Query(Email string, key string, query string) (string, error) {
	const api = "https://fofa.info/api/v1/search/all?email=%s&key=%s&qbase64=%s&size=10000"
	client := &http.Client{
		Timeout: 5 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	url := fmt.Sprintf(api, Email, key, base64.StdEncoding.EncodeToString([]byte(query)))
	req, err := http.NewRequest("GET", url, strings.NewReader(""))
	if err != nil {
		logsys.Error(err.Error())
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		logsys.Error(err.Error())
		return nil, err
	}
	result, err := io.ReadAll(resp)
	if err != nil {
		logsys.Error(err.Error())
		return nil, err
	}
	return result, err
}
