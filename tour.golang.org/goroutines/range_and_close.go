package main

import "fmt"

func fibonacci(n int, ch chan int) {
	f1, f2 := 0, 1

	for i := 0; i < n; i++ {
		ch <- f1
		f1, f2 = f2, f1+f2
	}

	close(ch)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)

	for i := range c {
		fmt.Println(i)
	}
}
