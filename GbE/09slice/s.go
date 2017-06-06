package main

import "fmt"

func main() {
	// var s [10]string
	s := make([]string, 3)
	fmt.Println(s)

	s[0] = "aaa"
	s[1] = "bbb"
	s[2] = "ccc"
	fmt.Println(s, len(s))

	s = append(s, "ddd")
	s = append(s, "eee", "fff")
	fmt.Println(s, len(s))

	// := is copy, not ref
	c := s
	c = append(c, "11")
	fmt.Println("c:", c)
	fmt.Println("s:", s)

	c = make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	l := s[2:5]
	fmt.Println("2:5", l)

	l = s[:5]
	fmt.Println(":5", l)

	t := []string{"a", "b", "c"}
	fmt.Println("new:", t)

	twoD := make([][]int, 3)
	fmt.Println(twoD)

	for i := 0; i < len(twoD); i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
