package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	v := fn(3, 4)
	return v * v
}

func ajax(url string, method string, callback func(err string, data string)) {
	fmt.Println("AJAX start")
	msg := "Success"
	callback("", msg)
	fmt.Println("AJAX end")
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(3, 4))

	fmt.Println(compute(hypot))

	callback := func(err string, data string) {
		if err == "" {
			fmt.Println("xatolik")
		} else {
			fmt.Println(data)
		}
	}

	ajax("https://data.io", "POST", callback)
}
