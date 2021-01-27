package main

import "fmt"

func main() {

	mySlice := make([]int, 0, 3)

	fmt.Println("-----------------")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println("-----------------")

	for i := 0; i < 5; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Len:", len(mySlice), "Capacity:", cap(mySlice), "Value: ", mySlice[i])
	}

	fmt.Println("Len:", len(mySlice), "Capacity:", cap(mySlice))

	l := mySlice[:6]

	fmt.Println(l)

	slcDel := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(slcDel)

	slcDel = append(slcDel, 8)

	fmt.Println(slcDel)

	//elementni o'chirish
	slcDel = append(slcDel[:3], slcDel[4:]...)
	fmt.Println(slcDel)

}
