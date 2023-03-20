package js

import v8 "rogchap.com/v8go"

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

func init() {
	iso := v8.NewIsolate()     // creates a new JavaScript VM
	ctx1 := v8.NewContext(iso) // new context within the VM
	ctx1.RunScript("const multiply = (a, b) => a * b", "math.js")

	ctx2 := v8.NewContext(iso) // another context on the same VM
	if _, err := ctx2.RunScript("multiply(3, 4)", "main.js"); err != nil {
		// this will error as multiply is not defined in this context
	}
}
