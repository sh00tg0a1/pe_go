package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for i := 1; i <= 3; {
		fmt.Println(i)
		i++
	}

	for {
		fmt.Println("Loop")
		break
	}

	for n := 1; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
