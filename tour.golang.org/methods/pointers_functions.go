package main

import "fmt"

type AddTwoSum struct {
	a, b int
}

func sum(n AddTwoSum) int {
	return n.a + n.b
}

func Scale(n *AddTwoSum, f int) {
	n.a += f
	n.b += f
}

func main() {

	addtwosum := AddTwoSum{1, 2}
	Scale(&addtwosum, 10)
	fmt.Println(sum(addtwosum))

}
