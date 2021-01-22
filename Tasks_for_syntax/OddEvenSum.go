package main

import "fmt"

func OddEvenSum(n int) (oddSum, evenSum int) {

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			evenSum = evenSum + i
		} else {
			oddSum = oddSum + i
		}
	}

	return

}

func main() {
	fmt.Print("Son kiriting ")

	var n int

	fmt.Scanf("%d", &n)
	var a, b = OddEvenSum(n)
	fmt.Println("Odd sum: ", a)
	fmt.Println("Even sum: ", b)

}

// 1 3 5 7 9 = 25
// 2 4 6 8 = 20
