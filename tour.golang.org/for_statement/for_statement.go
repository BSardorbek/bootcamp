package main

import (
	"fmt"
	"time"
)

func main() {

	//1-usuli
	sum := 10
	for i := 0; i <= sum; i++ {
		fmt.Println(i)
	}

	//2-usuli
	for {
		fmt.Println("ddos")
		time.Sleep(1 * time.Second) //1sekund kutadi

		if time.Now().Second()%2 == 0 {
			break
		}
	}

	//3-usuli
	i := 5
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	//4-usul arraylarda

}
