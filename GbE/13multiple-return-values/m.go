package main

import (
	"fmt"
)

func vals() (int, int) {
	return 1, 2
}

func main() {
	a, b := vals()
	fmt.Printf("a, b = %d, %d\n", a, b)

	_, c := vals()
	fmt.Println(c)
}
