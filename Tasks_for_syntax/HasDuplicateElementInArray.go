package main

import "fmt"

func hasDuplicateBad(arr []int) {
	for i := 0; i < len(arr); i++ {
		tmp := arr[i]
		for j, v := range arr {
			if tmp == v && i != j {
				fmt.Printf("array duplicate number %d \n ", arr[j])
				arr[i], arr = arr[len(arr)-1], arr[:len(arr)-1]
			}
		}
	}
}

func hasDuplicateGood(arr []int) {

	m := make(map[int]int)

	for _, v := range arr {
		m[v]++
	}

	fmt.Printf("array duplicate number %d \n ", m)

}

func main() {

	arr := []int{3, 4, 8, 6, 11, 4, 11, 4, 6}

	hasDuplicateBad(arr)

	hasDuplicateGood(arr)

}
