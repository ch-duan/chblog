package myStrings

import (
	"fmt"
	"log"
	"testing"
)

const (
	N = 2
)

func TestString(t *testing.T) {
	str := "Hello"
	log.Println(str[0])
	log.Println(str[0:1])
	str += "123"
	fmt.Println(str)
	// str[0] = 'm'
	var s [2 * N]struct{ x, y int32 }
	fmt.Println(len(s))
	var ss = [32]int{1, 2, 3}
	fmt.Println(ss)
}
