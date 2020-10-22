package mymap

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	var q map[string]int
	w := make(map[string]int)
	fmt.Println(q, w, q == nil, w == nil)

	e := make(map[string]int, 10)
	e["qqq"] = 1
	value, ok := e["qqq"]
	value2, ok2 := e["www"]
	fmt.Println(value, ok, value2, ok2)
	fmt.Println(len(e))
	delete(e, "qqq")
	value, ok = e["qqq"]
	fmt.Println(value, ok)
}
