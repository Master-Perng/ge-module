package tyc

import (
	"encoding/json"
	logsys "github.com/Master-Perng/go-module/log"
	jsoniter "github.com/json-iterator/go"
)

type Tycsub struct {
	Name string `json:"name"`
}

func GetSub(name string, token string) {
	//返回一个数组，每组都对应着一家
	result, err := TycInvestment(1, name, token)
	if err != nil {
		logsys.Error(err.Error())
	}
	totil := jsoniter.Get([]byte(result), "result").Get("total").ToInt()
	AllMap := make([]Tyc, totil)
	SubMap := make(map[string]map[string][]Tyc)
	json.Unmarshal([]byte(result), &SubMap)
	i := 0
	//资产入数组
	for s := range SubMap["result"]["items"] {
		AllMap[i].Name = SubMap["result"]["items"][s].Name
		AllMap[i].Percent = SubMap["result"]["items"][s].Percent
		i++
	}

	if totil > 20 {
		for k := 2; k <= totil/20+1; k++ {
			result, err := TycInvestment(k, name, token)
			if err != nil {
				logsys.Error(err.Error())
			}
			SubMap := make(map[string]map[string][]Tyc)
			json.Unmarshal([]byte(result), &SubMap)
			for s := range SubMap["result"]["items"] {
				AllMap[i].Name = SubMap["result"]["items"][s].Name
				AllMap[i].Percent = SubMap["result"]["items"][s].Percent
				i++
			}
		}
	}
}
