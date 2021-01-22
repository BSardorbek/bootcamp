package main

import "fmt"

func fibonacci(c, quit chan int) {
	f1, f2 := 0, 1
	for {
		select {
		case c <- f1:
			f1, f2 = f2, f1+f2
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)

	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}

		quit <- 0
	}()

	fibonacci(c, quit)

}
