package function

import "fmt"

type N int

func (n N) copy() {
	fmt.Printf("%p,%v\n", &n, n)
}

func (n *N) ref() {
	fmt.Printf("%p,%v\n", n, *n)
}
