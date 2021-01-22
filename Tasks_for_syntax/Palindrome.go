package main

import (
	"fmt"
)

func reverse(msg string) string {

	var tmp string

	for _, v := range msg {
		fmt.Println(string(v), tmp)
		tmp = string(v) + tmp
	}

	return tmp

}

func isPalindrom(msg string) bool {
	if reverse(msg) == msg {
		return true
	} else {
		return false
	}

}

func main() {
	fmt.Print("So'z yoki son kiriting ")

	var msg string

	fmt.Scanf("%s", &msg)

	if isPalindrom(msg) {
		fmt.Println("Palindrom.")
	} else {
		fmt.Println("Palindrom emas.")
	}
}

/*
	golang da char turi mavjud emas shuning uchun
	uning o'rniga uint8 va int32 tularini nomlangani mavjud
	ular:
		byte va rune
		byte ASCII ni kengaytiradi
		rune UTF-8 ni kengaytiradi
*/
