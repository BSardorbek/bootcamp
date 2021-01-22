package main

import (
	"fmt"
	"time"
)

type Person struct {
	name string
	year int
}

func (p Person) age() int {
	return time.Now().Year() - p.year
}

func age(p Person) int {
	return time.Now().Year() - p.year
}

func main() {
	sardor := Person{
		"sardorbek",
		1996,
	}

	//bu metod
	fmt.Println(sardor.age())

	//bu funksiya
	fmt.Println(age(Person{
		"navruz",
		1998,
	}))
}

/*
	class bo'lmagani uchun
	metod sifatida funksiya ishlatiladi
*/
