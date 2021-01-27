package main

import "fmt"

func main() {
	student := []int{}
	students := [][]int{{1, 2, 3}, {1, 2, 3}}

	fmt.Println(student)
	fmt.Println(students)

	var s []int

	fmt.Println(student == nil)
	fmt.Println(s == nil)

}
