package myStrings

import (
	"log"
	"testing"
)

func TestString(t *testing.T) {
	str := "Hello"
	log.Println(str[0])
	log.Println(str[0:1])
	// str[0] = 'm'

}
