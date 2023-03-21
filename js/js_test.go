package js

import (
	"fmt"
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
		ID:             0,
		Up:             0,
		Down:           0,
		Total:          0,
		Remark:         "",
		Enable:         false,
		ExpiryTime:     0,
		Autoreset:      false,
		Ipalert:        false,
		Iplimit:        0,
		Listen:         "",
		Port:           0,
		Protocol:       "",
		Settings:       "",
		StreamSettings: "",
		Tag:            "",
		Sniffing:       "",
	}
	link := node.GenLink()
	fmt.Println(link)
}
