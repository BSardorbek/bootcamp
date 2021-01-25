package main

import (
	"fmt"
	"time"
)

func visit(numbers int, cb func(int)) {
	fmt.Println(numbers)
	time.Sleep(3 * time.Second)

	cb(1)
}

func main() {

	cb := func(n int) {
		time.Sleep(3 * time.Second)
		fmt.Println(n)
	}

	visit(0, cb)
}
