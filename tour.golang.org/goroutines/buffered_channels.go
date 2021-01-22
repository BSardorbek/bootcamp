package main

import "fmt"

func concat(a, b string, ch chan string, isValid bool) {
	if isValid {
		ch <- "msg: " + a + b
	} else {
		ch <- "phone: " + a + b
	}
}

func main() {
	ch := make(chan string, 3)

	concat("salom ", "dunyo", ch, true)
	concat("+998", "936439414", ch, false)

	value := <-ch
	phone := <-ch

	fmt.Println(value, len(value))
	fmt.Println(phone, len(phone))

}
