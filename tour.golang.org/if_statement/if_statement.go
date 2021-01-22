package main

import (
	"fmt"
	"math"
)

func toqYokiJuft(number int64) string {
	if number%2 == 0 {
		return "Juft"
	}

	return "toq"
}

func daraja(n, m, chegara float64) (float64, string) {
	if d := math.Pow(n, m); d > chegara {
		return d, "chegaradan kotta"
	}

	return 0, "chegaradan kichik"
}

func main() {
	a := int64(55)
	var b int64 = 32
	var c int64
	c = 54

	fmt.Println(toqYokiJuft(c))
	fmt.Println(toqYokiJuft(b))
	fmt.Println(toqYokiJuft(a))

	n, m := 2., 2.
	var chegara float64
	chegara = 121

	fmt.Println(daraja(n, m, chegara))
}
