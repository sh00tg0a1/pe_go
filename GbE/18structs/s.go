package main

import "fmt"

type person struct {
	name   string
	age    int
	gender bool
}

func main() {
	fmt.Println(person{"Bob", 20, true})
	fmt.Println(person{name: "Alice", age: 20, gender: false})

	fmt.Println(&person{"Ann", 20, false})

	s := person{"Sean", 45, true}
	fmt.Println(s.name, s.age)

	// 都是深度拷贝
	sp := &s
	sp.age = 51

	fmt.Println(s.age)
}
