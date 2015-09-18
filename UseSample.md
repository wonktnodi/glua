这是一个简单的Go Lua binding，希望对你有用。

简单举例如下：

```
package main

import (
	"fmt"
	"glua"
)

type Int struct {
	I int
}

func NewInt() *Int {
	return &Int{10}
}

func (i Int) PrintInt(str string) {
	fmt.Println(str, i.I)
}

func main() {
	L := glua.NewState()
	
	L.Openlibs()

	var tlib = glua.Libfuncs{
		"gotest", // lib name
		map[string]interface{}{
			"NewInt":    NewInt,          // lua function name, go function
			"PrintInt":  (*Int).PrintInt, // lua function name, go function
			"goprintln": fmt.Println,
		},
	}
	if ok, err := L.Register(&tlib); !ok {
		fmt.Println(err.Error())
		return
	}
	
	L.Dostring(`gotest.PrintInt(gotest.NewInt(), "Int is")`) 

	L.Dofile("test.lua")
	
	L.Call("gotest.goprintln", "Call lua function.", 123456)
}
```

```
--test.lua
gotest.goprintln(true, 123, "lua", gotest.NewInt())
```