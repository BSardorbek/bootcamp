package main

import "fmt"

func main() {
	var a interface{}
	fmt.Println(a) //nil

	a = "a"
	fmt.Println(a)          //a
	fmt.Println(a.(string)) //a

	do(a)
}

func do(i interface{}) {
	fmt.Println(i.(string))
}
