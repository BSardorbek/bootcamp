package main

import (
	"fmt"
	"strconv"
)

func FizzBuzz(n int) string {

	if n%3 == 0 && n%5 == 0 {
		return "FizzBuzz"
	} else if n%3 == 0 {
		return "fizz"
	}
	if n%5 == 0 {
		return "buzz"
	} else {
		return strconv.Itoa(n)
	}

}

func main() {
	fmt.Print("Son kiriting ")
	var n int

	fmt.Scanf("%d", &n)

	for i := 0; i <= n; i++ {
		fmt.Println(FizzBuzz(i))
	}
}
