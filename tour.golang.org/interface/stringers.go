package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Printer() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{
		"Sardor",
		25,
	}

	b := Person{
		"navruz",
		23,
	}

	fmt.Println(a, b)
}
