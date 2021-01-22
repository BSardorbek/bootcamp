package main

import (
	"fmt"
	"strings"
)

func main() {

	var arr [2]string
	arr[0] = "Hello"
	arr[1] = "World"

	fmt.Println(arr)

	primes := [6]int{2, 3, 5, 7, 11, 13}

	fmt.Println(primes)

	number := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(number)

	//array uzunligini olish uchun len() funksiyasini ishlatish kerak
	fmt.Println(len(arr))
	fmt.Println(cap(arr))

	//yangi massiv yaratish
	b := make([]string, 8)
	fmt.Println(b)
	fmt.Println(len(b), cap(b))

	fmt.Println(primes[:5])  //boshidan 5 tani ajratib oladi
	fmt.Println(primes[2:])  //boshidan 2 tani tashlab ketadi
	fmt.Println(primes[2:5]) //birinchi beshtani oladi keyin ikkitani tashlab yuboradi

	a := []string{"s", "a", "r", "d", "o", "r"}
	fmt.Println(strings.Join(a, ""))
}
