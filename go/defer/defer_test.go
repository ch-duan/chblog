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

//defer的顺序是FILO先入后出，defer后面是表达式时会先计算表达式，直到最后一个方法
func TestDeferCall(t *testing.T) {
	s := NewSlice()
	defer s.AddSlice(1).AddSlice(3)
	defer s.AddSlice(4).AddSlice(5)
	s.AddSlice(2)

}

//测试传入变量或指针
func TestB(t *testing.T) {
	x := 100
	p := &x
	defer fmt.Println("defer:", x)
	defer fmt.Println("defer:", *p)
	x++
	defer fmt.Println("defer:", *p)
	fmt.Println("main:", x, *p)
}

//闭包需要传递一个i的参数，不然会使用同一个i的内存
func TestA(t *testing.T) {
	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Printf("%p,%d\r\n", &i, i)
		}()
	}
	fmt.Println("qqqq")
	for i := 0; i < 10; i++ {
		defer qqq(i)
	}
	fmt.Println("22:", qwe())
}

func qqq(value int) {
	fmt.Printf("%p,%v\r\n", &value, value)
}

func qwe() (res int) {
	fmt.Println("qwe:", res)
	defer func() {
		res++
		fmt.Println("qwe1:", res)
	}()
	defer func() {
		res++
		fmt.Println("qwe2:", res)
	}()
	return 1
}
