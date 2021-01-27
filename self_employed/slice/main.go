package main

import "fmt"

func main() {

	var arr [12]int
	mySlice := []int{1, 3, 5, 7, 9, 11}
	arr[1] = 12
	fmt.Printf("%T\n", mySlice)
	fmt.Printf("%T\n", arr)

	fmt.Println(mySlice)
	fmt.Println(arr)
}
