package mydefer

import (
	"fmt"
	"log"
	"testing"
)

type Slice []int

func NewSlice() *Slice {
	slice := make(Slice, 0)
	return &slice
}

func (s *Slice) AddSlice(val int) *Slice {
	*s = append(*s, val)
	log.Println(val)
	return s
}

func TestDeferCall(t *testing.T) {
	s := NewSlice()
	defer s.AddSlice(1).AddSlice(3)
	defer s.AddSlice(4).AddSlice(5)
	s.AddSlice(2)

}

func TestA(t *testing.T) {
	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
	fmt.Println("qqqq")
	for i := 0; i < 10; i++ {
		defer qqq(i)
	}
	fmt.Println("22:", qwe())
}

func qqq(value int) {
	fmt.Println(value)
}

func qwe() (res int) {
	defer func() {
		res++
		fmt.Println(res)
	}()
	defer func() {
		res++
		fmt.Println("1:", res)
	}()
	return 1
}
