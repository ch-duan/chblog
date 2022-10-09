package function

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	var a N = 25
	var p *N = &a
	fmt.Printf("%p\r\n", &a)
	a.copy()
	N.copy(a)
	a.ref()
	p.ref()
	p.copy()
	m := [...]int{
		1: 1, 4: 2, 3,
	}
	fmt.Println(m)
	data := make(chan int)
	quit := make(chan struct{})
	go func() {
		data <- 11
	}()
	go func() {
		defer close(quit)
		fmt.Println(<-data)
		fmt.Println(<-data)
	}()
	data <- 22
	<-quit
}
