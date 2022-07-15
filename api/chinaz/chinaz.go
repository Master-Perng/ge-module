package chinaz

import (
	"fmt"
	logsys "github.com/Master-Perng/go-module/log"
	"io"
	"net/http"
	"strings"
	"time"
)

const api = "https://apidatav2.chinaz.com/single"

func Re_whois(key string, queryData string, queryType string) (string, error) {
	client := &http.Client{
		Timeout: 15 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	re_icp_api := api + "/whoisreverse?key=%s&queryData=%s&queryType=%s"
	url := fmt.Sprintf(re_icp_api, key, queryData, queryType)
	req, err := http.NewRequest("GET", url, strings.NewReader(""))
	if err != nil {
		defer client.CloseIdleConnections()
		logsys.Error(err.Error())
		return "", err
	}
	resp, err := client.Do(req)
	if strings.Contains(err.Error(), "Timeout") {
		for {
			time.Sleep(2 * time.Second)
			resp, err = client.Do(req)
			if !strings.Contains(err.Error(), "Timeout") {
				break
			}
		}
	}

	if err != nil {
		defer client.CloseIdleConnections()
		logsys.Error(err.Error())
		return "", err
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
func Whois(key string, domain string) (string, error) {
	client := &http.Client{
		Timeout: 15 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	re_icp_api := api + "/whois?key=%s&domain=%s"
	url := fmt.Sprintf(re_icp_api, key, domain)
	req, err := http.NewRequest("GET", url, strings.NewReader(""))
	if err != nil {
		defer client.CloseIdleConnections()
		logsys.Error(err.Error())
		return "", err
	}
	resp, err := client.Do(req)
	if strings.Contains(err.Error(), "Timeout") {
		for {
			time.Sleep(2 * time.Second)
			resp, err = client.Do(req)
			if !strings.Contains(err.Error(), "Timeout") {
				break
			}
		}
	}
	if err != nil {
		defer client.CloseIdleConnections()
		logsys.Error(err.Error())
		return "", err
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
func Icp(key string, domain string) (string, error) {
	client := &http.Client{
		Timeout: 15 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	re_icp_api := api + "/newicp?key=%s&domain=%s"
	url := fmt.Sprintf(re_icp_api, key, domain)
	req, err := http.NewRequest("GET", url, strings.NewReader(""))
	if err != nil {
		defer client.CloseIdleConnections()
		logsys.Error(err.Error())
		return "", err
	}
	resp, err := client.Do(req)
	if strings.Contains(err.Error(), "Timeout") {
		for {
			time.Sleep(2 * time.Second)
			resp, err = client.Do(req)
			if !strings.Contains(err.Error(), "Timeout") {
				break
			}
		}
	}

	if err != nil {
		defer client.CloseIdleConnections()
		logsys.Error(err.Error())
		return "", err
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
func Re_Icp(key string, companyname string) (string, error) {
	client := &http.Client{
		Timeout: 15 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	re_icp_api := api + "/newsponsorunit?key=%s&companyname=%s"
	url := fmt.Sprintf(re_icp_api, key, companyname)
	req, err := http.NewRequest("GET", url, strings.NewReader(""))
	if err != nil {
		defer client.CloseIdleConnections()
		logsys.Error(err.Error())
		return "", err
	}
	resp, err := client.Do(req)
	if strings.Contains(err.Error(), "Timeout") {
		for {
			time.Sleep(2 * time.Second)
			resp, err = client.Do(req)
			if !strings.Contains(err.Error(), "Timeout") {
				break
			}
		}
	}

	if err != nil {
		defer client.CloseIdleConnections()
		logsys.Error(err.Error())
		return "", err
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
