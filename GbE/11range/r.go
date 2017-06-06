package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	for i, num := range nums {
		fmt.Println("i:", i, "v:", num)
	}

	kvs := map[string]int{"a": 1, "b": 2}
	fmt.Println(kvs)

	for k, v := range kvs {
		fmt.Println("k:", k, "v:", v)
	}

	for k := range kvs {
		fmt.Println("k:", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
