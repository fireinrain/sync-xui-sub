package js

import (
	"fmt"
	"os"
	"reflect"
	v8 "rogchap.com/v8go"
	"testing"
)

func TestV8Go(t *testing.T) {
	ctx := v8.NewContext()                           // new context with a default VM
	obj := ctx.Global()                              // get the global object from the context
	obj.Set("version", "v1.0.0")                     // set the property "version" on the object
	val, _ := ctx.RunScript("version", "version.js") // global object will have the property set within the JS VM
	fmt.Printf("version: %s", val)

	if obj.Has("version") { // check if a property exists on the object
		obj.Delete("version") // remove the property from the object
	}
}

func TestGenLink(t *testing.T) {
	node := ServerNode{
		ID:             5,
		Up:             619517876,
		Down:           619517876,
		Total:          53687091200,
		Remark:         "speedtest",
		Enable:         false,
		ExpiryTime:     0,
		Autoreset:      false,
		Ipalert:        false,
		Iplimit:        0,
		Listen:         "",
		Port:           59876,
		Protocol:       "vmess",
		Settings:       "{\n  \"clients\": [\n    {\n      \"id\": \"14d6853c-813c-49ff-8d09-7edbe832af44\",\n      \"alterId\": 0\n    }\n  ],\n  \"disableInsecureEncryption\": false\n}",
		StreamSettings: "{\n  \"network\": \"ws\",\n  \"security\": \"tls\",\n  \"tlsSettings\": {\n    \"serverName\": \"cloud2.131433.xyz\",\n    \"certificates\": [\n      {\n        \"certificateFile\": \"/nginxweb/cert/fullchain.cer\",\n        \"keyFile\": \"/nginxweb/cert/private.key\"\n      }\n    ]\n  },\n  \"wsSettings\": {\n    \"path\": \"/fire\",\n    \"headers\": {}\n  }\n}",
		Tag:            "inbound-59876",
		Sniffing:       "{\n  \"enabled\": true,\n  \"destOverride\": [\n    \"http\",\n    \"tls\"\n  ]\n}",
	}
	link, _ := node.GenLink()
	fmt.Println(link)
}

func TestV8Go2(t *testing.T) {
	ctx := v8.NewContext()
	file, err := os.ReadFile("core.js")
	if err != nil {
		panic(err)
	}
	coreJs := string(file)
	ctx.RunScript(coreJs, "corelink.js")
	// new context with a default VM
	ctx.RunScript("const result = selfTestAdd(1,2)", "sample.js") // global object will have the property set within the JS VM

	val, _ := ctx.RunScript("result", "result.js") // global object will have the property set within the JS VM
	fmt.Printf("version: %s", val)

	reflect.DeepEqual(val, "3")

}
