package main

import "fmt"

func FibonacciClosure() func() int {

	f1, f2 := 0, 1
	//fn = fn-1 + fn-2
	return func() int {
		f1, f2 = f2, f1+f2
		return f1
	}

}

func FibonacciRecursion(i int) int {

	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}

	return FibonacciRecursion(i-1) + FibonacciRecursion(i-2)
}

func main() {
	fmt.Print("Son kiriting ")
	var n int

	fmt.Scanf("%d", &n)

	fc := FibonacciClosure()

	for i := 0; i < n; i++ {
		fmt.Println(fc())
	}

	println("=========")

	for i := 0; i <= n; i++ {
		fmt.Println(FibonacciRecursion(i))
	}

}
