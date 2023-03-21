package js

import (
	"encoding/json"
	"fmt"
	"log"
	v8 "rogchap.com/v8go"
	"strings"
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

func init() {
	V8Context = v8.NewContext() // new context with a default VM
}

func (receiver ServerNode) GenLink() (string, error) {
	nodeData, err := json.Marshal(receiver)
	if err != nil {
		log.Println("json.Marshal failed: ", err)
		return "", err
	}
	nodeDataStr := strings.ReplaceAll(string(nodeData), "\\n", "")

	//parse, err := v8.JSONParse(V8Context, nodeDataStr)
	//fmt.Println(parse)
	//将nodeData转换为js object
	jsonToObj := fmt.Sprintf("const nodeData = JSON.parse('%s')", nodeDataStr)
	script, err := V8Context.RunScript(jsonToObj, "jsonToObj.js")
	fmt.Println(script)
	//V8Context.RunScript("const result = add(3, 4)", "main.js")
	val, err := V8Context.RunScript("nodeData", "jsonToObj.js")
	fmt.Printf("addition result: %s", val)
	fmt.Printf(err.Error())
	return "", nil
}
