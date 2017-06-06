package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func fab() func() int {
	a, b, s := 0, 1, 0
	return func() int {
		s, a, b = a, b, a+b
		return s
	}

}

func fabN(n int) int {
	res := 0
	num := fab()
	for i := 0; i < n; i++ {
		res = num()
	}
	return res
}

func main() {
	// nextInt := intSeq()
	// fmt.Println(nextInt())
	// fmt.Println(nextInt())
	// fmt.Println(nextInt())

	// newInts := intSeq()
	// fmt.Println(newInts())

	// fmt.Println(intSeq()())
	// fmt.Println(intSeq()())
	// fmt.Println(intSeq()())

	seq := make([]int, 0)
	fabNum := fab()

	for i := 0; i < 20; i++ {
		seq = append(seq, fabNum())
	}

	fmt.Println(seq)

	fmt.Println(fabN(20))
}
