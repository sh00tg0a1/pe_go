package main

import (
	"fmt"
)

type rect struct {
	with, height int
}

func (r *rect) area() int {
	return r.with * r.height
}

func (r *rect) perim() int {
	return 2*r.with + 2*r.height
}

func main() {
	r := rect{2, 3}
	fmt.Println(r.area(), r.perim())

	rp := &r
	rp.with = 10
	r.height = 10

	fmt.Println(rp.area(), rp.perim())
}
