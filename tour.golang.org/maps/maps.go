package main

import (
	"fmt"
	"strings"
)

func logmap(m map[string]int) {
	for key, value := range m {
		fmt.Println(key, "->", value)
	}

	for _, value := range m {
		fmt.Println("Value: ", value)
	}

	for key := range m {
		fmt.Println(key)
	}
}

type Vertex struct {
	Lat, Long float64
}

type Adge struct {
	id int64
}

func main() {

	// Create new map string -> int

	/*
		key -> value
		string -> int
		map[key type]value type
	*/
	m := make(map[string]int)

	m["a"] = 10
	m["b"] = 20
	m["c"] = 30

	fmt.Println(m)
	logmap(m)

	var ver = map[string]Vertex{
		"Uzbekistan": {
			42.5, 41.23,
		},
		"Boston": {
			72.5, 33.23,
		},
	}

	fmt.Println(ver["Uzbekistan"])

	mm := make(map[string]int)

	mm["Javob"] = 1996

	fmt.Println(mm)
	delete(mm, "Javob")
	fmt.Println(mm)

	value, isKey := m["Javob"]
	fmt.Println(value, isKey)

	fmt.Println(test("I am learning Go!"))

	var tt string = "I am learning Go!"
	fmt.Println(test(tt))

}

func test(txt string) map[string]int {

	m := make(map[string]int)
	arr := strings.Fields(txt)

	for _, value := range arr {
		m[value]++
	}

	return m
}
