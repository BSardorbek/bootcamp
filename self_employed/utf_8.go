package main

import "fmt"

const SS := 123

func main() {

	for i := 60; i < 122; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}

}
