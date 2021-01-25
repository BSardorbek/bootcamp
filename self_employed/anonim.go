package main

import "fmt"

func main() {
	func(args int) {
		fmt.Println("I'm driving!" + string(args))
	}(1)
}
