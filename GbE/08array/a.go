package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a)

	a[4] = 100
	fmt.Println(a)
	fmt.Println("Length:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	var twoD [9][9]int
	for i := 1; i <= len(twoD); i++ {
		for j := 1; j <= len(twoD[i]); j++ {
			if i < j {
				continue
			}
			twoD[i-1][j-1] = i * j
		}
		fmt.Println(twoD[i-1])
	}
}
