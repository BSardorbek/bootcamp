package main

import "fmt"

func main() {

	x := 8
	fmt.Println(x)  //8
	fmt.Println(&x) //0xc00009e058

	b := &x //x ni xotiradagi joylashgan joy nomi

	fmt.Println(b)  //0xc00009e058
	fmt.Println(*b) //8

	d := true
	food := "1"
	if food += "Chocolate"; d {
		fmt.Println(food)
	}

}
