package main

import "fmt"

func main() {

	day, year := 26, 1996

	p := &day
	fmt.Println(*p) //26
	*p = 22
	fmt.Println(day)  //22
	fmt.Println(year) //1996

}
