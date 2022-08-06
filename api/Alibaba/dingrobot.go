package Alibaba

import (
	"encoding/json"
	"fmt"
	log "github.com/Master-Perng/go-module/log"
	"io"
	"net/http"
	"strings"
)

type Markdown struct {
	Msgtype  string `json:"msgtype"`
	Markdown struct {
		Title string `json:"title"`
		Text  string `json:"text"`
	} `json:"markdown"`
	isAtAll bool `json:"isAtAll"`
}

func DingBotMarkDown(title string, text string, api string) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	GoData := Markdown{
		Msgtype: "markdown",
		isAtAll: true,
	}
	GoData.Markdown.Title = title
	GoData.Markdown.Text = text
	Postbody, err := json.Marshal(GoData)
	req, err := http.NewRequest("POST", api, strings.NewReader(string(Postbody)))
	if err != nil {
		log.Error("Error :", err.Error())
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Error("Error :", err.Error())
	}
	context, _ := io.ReadAll(resp.Body)
	fmt.Println(string(context))
}
