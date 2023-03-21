package js

import (
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

var V8COntext *v8.Context

func init() {
	V8COntext = v8.NewContext() // new context with a default VM
}

func (receiver ServerNode) GenLink() string {

	return ""
}
