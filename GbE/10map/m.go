package main

import (
	"fmt"
)

func main() {
	// var n map[string]int
	// fmt.Println(n)
	m := make(map[string]int)
	fmt.Println(m)

	m["k1"] = 1
	m["k2"] = 2

	fmt.Println(m, len(m))

	v1 := m["k1"]
	fmt.Println(v1)

	delete(m, "k1")
	fmt.Println(m)

	// first is the value, second is presents
	_, prs := m["k2"]
	fmt.Println("prs", prs)

	n := map[string]int{"a": 1, "b": 2}
	fmt.Println(n)
}
