package js

import (
	v8 "rogchap.com/v8go"
	"testing"
)

func TestV8Go(t *testing.T) {
	iso := v8.NewIsolate()     // creates a new JavaScript VM
	ctx1 := v8.NewContext(iso) // new context within the VM
	ctx1.RunScript("const multiply = (a, b) => a * b", "math.js")
	ctx1.RunScript("multiply(1,2)", "math.js")
}
