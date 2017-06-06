package main

import "fmt"

func zeroVal(val int) {
	val = 0
}

func zeroPtr(ptr *int) {
	*ptr = 0
}

func main() {
	a, b := 100, 200
	fmt.Println("before", a, b)
	zeroVal(a)
	zeroPtr(&b)

	fmt.Println("after", a, b)

	aptr := &a
	fmt.Println("pointer", &a, &aptr)
}
