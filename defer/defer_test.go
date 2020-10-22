package mydefer

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	for i := 0; i < 10; i++ {
		defer func() {
			// fmt.Println(i)
		}()
	}
	fmt.Println("qqqq")
	for i := 0; i < 10; i++ {
		defer qqq(i)
	}
}

func qqq(value int) {
	fmt.Println(value)
}
