package main

import "fmt"

func makeGreeter(args string) func() string {
	return func() string {
		return "Hello world!"
	}
}

func main() {
	greet := makeGreeter("123")()
	fmt.Println(greet)
	fmt.Printf("%T\n", makeGreeter)
}
