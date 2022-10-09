package mygo

import (
	"fmt"
	"testing"
)

func TestType(t *testing.T) {
	var a, b interface {
		test()
	}
	fmt.Println(a == b)
}

func TestArray(t *testing.T) {
	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := a[2:6:7]
	fmt.Println(s)
	fmt.Println(len(s), cap(s))

	fmt.Printf("a:%p ~ %p\t%d ~ %d\n", &a[0], &a[len(a)-1], a[0], a[len(a)-1])
	fmt.Printf("s:%p ~ %p\t%d ~ %d\n", &s[0], &s[len(s)-1], s[0], s[len(s)-1])

}
