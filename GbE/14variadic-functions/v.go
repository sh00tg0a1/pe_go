package main

import "fmt"

func sum(nums ...int) int {
	fmt.Print(nums, " ")
	total := 0
	for _, v := range nums {
		total += v
	}
	fmt.Println(total)

	return total
}

func main() {
	sum(1)
	sum(1, 2, 3)
	nums := []int{4, 5, 6}
	sum(nums...)
}
