package main

import "fmt"

var powNumber = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {

	for index, value := range powNumber {
		fmt.Printf("2**%d = %d\n", index, value)
	}

	println()
	numbers := make([]int, 63)
	for i := range numbers {
		numbers[i] = 1 << uint(i)
	}

	for _, value := range numbers {
		fmt.Printf("%d\n", value)
	}

	for index := range numbers {
		fmt.Printf("%d\n", index)
	}
}
