package js

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	v8 "rogchap.com/v8go"
)

type ServerNode struct {
	ID             int    `json:"id"`
	Up             int64  `json:"up"`
	Down           int64  `json:"down"`
	Total          int    `json:"total"`
	Remark         string `json:"remark"`
	Enable         bool   `json:"enable"`
	ExpiryTime     int    `json:"expiryTime"`
	Autoreset      bool   `json:"autoreset"`
	Ipalert        bool   `json:"ipalert"`
	Iplimit        int    `json:"iplimit"`
	Listen         string `json:"listen"`
	Port           int    `json:"port"`
	Protocol       string `json:"protocol"`
	Settings       string `json:"settings"`
	StreamSettings string `json:"streamSettings"`
	Tag            string `json:"tag"`
	Sniffing       string `json:"sniffing"`
}

//在golang中执行js

var V8Context *v8.Context
var CoreJS string

func init() {
	V8Context = v8.NewContext() // new context with a default VM
	file, err := os.ReadFile("js/core.js")
	if err != nil {
		fmt.Println("read core.js error:", err)
	} else {
		CoreJS = string(file)
	}
}

func (receiver ServerNode) GenLink() (string, error) {
	standard := FormatJsonStrForStandard(receiver)

	ctx := v8.NewContext()
	coreJs := CoreJS
	ctx.RunScript(coreJs, "corelink.js")
	sprintf := fmt.Sprintf("const jsonStr = %s", standard)
	ctx.RunScript(sprintf, "variable.js")
	_, err2 := ctx.RunScript("const temp = JSON.stringify(jsonStr); const result = getLinkFromJsonStr(temp);", "run.js")
	if err2 != nil {
		fmt.Println("v8 run js error: ", err2.Error())
		return "", err2
	}
	val, _ := ctx.RunScript("result", "result.js") // global object will have the property set within the JS VM
	//fmt.Printf("result: %s", val)
	return val.String(), nil
}

func FormatJsonStrForStandard(node ServerNode) string {
	nodeData, err := json.Marshal(node)
	if err != nil {
		log.Println("json.Marshal failed: ", err)
		return ""
	}
	//nodeDataStr := strings.ReplaceAll(string(nodeData), "\\n", "")

	//var data map[string]interface{}
	//err = json.Unmarshal([]byte(nodeDataStr), &data)
	//if err != nil {
	//	log.Println("json.UnMarshal failed: ", err)
	//	return ""
	//}
	//
	//// convert settings and sniffing fields to JSON objects
	//settingsStr := data["settings"].(string)
	//var settingsObj map[string]interface{}
	//err = json.Unmarshal([]byte(settingsStr), &settingsObj)
	//if err != nil {
	//	panic(err)
	//}
	//data["settings"] = settingsObj
	//
	//streamSettingsStr := data["streamSettings"].(string)
	//var streamSettingsObj map[string]interface{}
	//err = json.Unmarshal([]byte(streamSettingsStr), &settingsObj)
	//if err != nil {
	//	panic(err)
	//}
	//data["streamSettings"] = streamSettingsObj
	//
	//sniffingStr := data["sniffing"].(string)
	//var sniffingObj map[string]interface{}
	//err = json.Unmarshal([]byte(sniffingStr), &sniffingObj)
	//if err != nil {
	//	panic(err)
	//}
	//data["sniffing"] = sniffingObj
	//
	//// marshal data to standard JSON format
	//jsonData, err := json.MarshalIndent(data, "", "  ")
	//if err != nil {
	//	panic(err)
	//}
	// print standard JSON string
	//return string(jsonData)
	//return nodeDataStr
	return string(nodeData)
}
