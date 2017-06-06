package main

import (
	"fmt"
	"math"
)

type rect struct {
	with, height float64
}

func (r rect) area() float64 {
	return r.with * r.height
}

func (r rect) perim() float64 {
	return 2*r.with + 2*r.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return (1.0 / 2) * math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

type geometry interface {
	area() float64
	perim() float64
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{10, 10}
	c := circle{5}

	measure(r)
	measure(c)
}
