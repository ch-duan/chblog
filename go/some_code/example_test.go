package mExample

import (
	"fmt"
	"testing"
)

type aa struct {
	name string
}

func TestTest(t *testing.T) {
	a := &aa{name: "111"}
	m := map[string]*aa{"uno": a}
	p := m["uno"]
	fmt.Println(m["uno"])
}
