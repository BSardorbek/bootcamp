package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func addTwoSum(x int, y int) int {
	return x + y
}

func subTwoSum(a, b int) int {
	return a - b
}

func swap(txt1, txt2 string) (string, string) {
	return txt2, txt1
}

func nakedReturn(sum int) (a, b int) {
	a = sum / 10
	b = sum % 10

	return a, b
}

/* global elon qilish var kalit so'zi ishlatishi
kerak aks xolda xatolik beradi
a := 1 //error syntax
*/
var n, m bool
var a, b int = 1, 2
var k, l, isShow = 12, "yes", true

func main() {

	fmt.Println("Salom bootcamp")
	fmt.Println("Hozirgi vaqt ", time.Now())
	fmt.Println("Random ", rand.Intn(15))
	fmt.Println("Math kutubxonasi", math.Sqrt(81))
	fmt.Println("math kutubxonasidan PI ", math.Pi)

	fmt.Println(addTwoSum(100, 200))
	fmt.Println(subTwoSum(50, 23))

	txt1, txt2 := swap("dunyo", "salom")
	fmt.Println(txt1 + " " + txt2)

	fmt.Println(swap("world", "hello"))

	//ikki xonali son uchun
	birlar, onlar := nakedReturn(21)
	fmt.Println(birlar)
	fmt.Println(onlar)
	fmt.Println(123)

	fmt.Println(n, m) //bool default holatda fasle ni oladi
	var n = true      //new var override global declare p
	fmt.Println(n)
	/*
		n := false //bu qisaqa elon qilish
		var n = true // bu toliq elon qilish

		xatolik beradi chunki oldin elon qilingan localda
	*/
	fmt.Println(n)

}

/*

uint8	unsigned 8-bit integer	(0 to 255)
uint16	unsigned 16-bit integer	(0 to 65535)
uint32	unsigned 32-bit integer	(0 to 4294967295)
uint64	unsigned 64-bit integer	(0 to 18446744073709551615)
int8	signed 8-bit integer	(-128 to 127)
int16	signed 16-bit integer	(-32768 to 32767)
int32	signed 32-bit integer	(-2147483648 to 2147483647)
int64	signed 64-bit integer	(-9223372036854775808 to 9223372036854775807)


*/
