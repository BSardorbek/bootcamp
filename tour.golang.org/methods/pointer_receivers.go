package main

import (
	"fmt"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) power() float64 {
	return v.X * v.Y
}

func (v *Vertex) add(f float64) {
	v.X = v.X + f
	v.Y = v.Y + f
}

func main() {

	multi := Vertex{
		3., 2.,
	}

	multi.add(2.)
	fmt.Println(multi.power())

}
