package main

import (
	"fmt"
	"math"
)

//
type ShapeI interface {
	area() float64
}

//
type Circle struct {
	radius float64
}

//
type Square struct {
	side float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s Square) area() float64 {
	return s.side * s.side
}

func info(i ShapeI) {
	fmt.Println(i)
	fmt.Println(i.area())
}

func main() {
	s := Square{10}
	c := Circle{5}
	info(s)
	info(c)
}
