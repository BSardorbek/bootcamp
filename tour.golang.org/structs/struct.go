package main

import "fmt"

type Person struct {
	name string
	age  int
	id   int8
}

type Student struct {
	group string
	gpa   int
	p_id  int
}

func main() {
	fmt.Println(Person{
		"sardorbek",
		24,
		1,
	})

	fmt.Println(Student{
		"M414",
		5,
		1,
	})

	p := Person{
		"navruz",
		22,
		2,
	}

	p.id = 3

	fmt.Println(p)

	m := &p
	m.age = 25
	m.name = "Doston"
	m.id = 4
	fmt.Println(m)
}
