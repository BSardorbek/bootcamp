package main

import "fmt"

func fibonacci() func() int {
	f1, f2 := 1, 1
	var fn int
	//fn = fn-1 + fn-2

	return func() int {
		if f1 == 1 && f2 == 1 {
			fn = f1 + f2
			f1 = f2
			f2 = fn
		} else {
			fn = f1 + f2
			f1, f2, = f2, fn
		}

		return fn
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
