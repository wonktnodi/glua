/*
	对glau功能的测试
*/
package glua

import (
	"fmt"
	"runtime"
	"testing"
)

func TestDofile(t *testing.T) {
	L := NewState()
	L.Openlibs()
	if !L.Dofile("testdofile.lua") {
		t.Error("Dofile error.")
	}
}

func TestDofilegc(t *testing.T) {
	runtime.GC()
}

func TestDostring(t *testing.T) {
	L := NewState()
	L.Openlibs()
	if !L.Dostring("print(2); print(math.pi)") {
		t.Error("Dofile error.")
	}
}

func TestDostringgc(t *testing.T) {
	runtime.GC()
}

func printf(b bool, i int, s string) (bool, int, string) {
	fmt.Println("this is a func in go test.")
	fmt.Println(b, i, s)
	return b, i, s
}

func getSlice() []int {
	return []int{1, 2, 3}
}

func printSlice(si []int) {
	fmt.Println(si)
}

var testfunc int

var tlib = Libfuncs{
	"gotest", // lib name
	map[string]interface{}{
		"printf":     printf,     // lua function name, go function
		"getSlice":   getSlice,   //
		"printSlice": printSlice, //
		"goprintln":  fmt.Println,
		"testfunc":   testfunc,  // error
	},
}

func TestRegLib(t *testing.T) {
	L := NewState()
	L.Openlibs()
	if ok, err := L.Register(&tlib); !ok {
		t.Error(err.Error())
	}
	if !L.Dofile("testregister.lua") {
		t.Error("Dofile error.")
	}

}
