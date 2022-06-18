package tyc

/*
type Tyc struct {
	Name    string `json:"name"`
	Percent string `json:"percent"`
}
type Tycsub struct {
	Name string `json:"name"`
}
*/
/*func GetSub(name string, token string) []Tyc {
	//返回一个数组，每组都对应着一家
	/*result1, err := TycSub(1, name, token)
	if err != nil {
		logsys.Error(err.Error())
	}
	totil1 := jsoniter.Get([]byte(result1), "result").Get("total").ToInt()
	result, err := TycInvestment(1, name, token)
	if err != nil {
		logsys.Error(err.Error())
	}
	totil := jsoniter.Get([]byte(result), "result").Get("total").ToInt()
	AllMap := make([]Tyc, totil+totil1)
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
	//天眼查下属单位
	SubMap1 := make(map[string]map[string][]Tycsub)
	json.Unmarshal([]byte(result1), &SubMap1)
	//资产入数组
	for s := range SubMap1["result"]["items"] {
		AllMap[i].Name = SubMap1["result"]["items"][s].Name
		AllMap[i].Percent = "分支"
		i++
	}
	if totil1 > 20 {
		for k := 2; k <= totil1/20+1; k++ {
			result, err := TycInvestment(k, name, token)
			if err != nil {
				logsys.Error(err.Error())
			}
			SubMap1 := make(map[string]map[string][]Tycsub)
			json.Unmarshal([]byte(result), &SubMap1)
			for s := range SubMap1["result"]["items"] {
				AllMap[i].Name = SubMap1["result"]["items"][s].Name
				AllMap[i].Percent = "分支"
				i++
			}
		}
	}
	return AllMap
}*/
