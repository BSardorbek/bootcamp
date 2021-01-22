package main

import "fmt"

type Phone interface {
	SimCard()
}

type Mobile struct {
	number string
}

func (mobile Mobile) SimCard() {
	fmt.Println(mobile.number)
}

func main() {

	var ucell Phone = Mobile{
		"+998936439414",
	}

	ucell.SimCard()

	var beeline Phone = Mobile{
		"+998907705701",
	}
	beeline.SimCard()

}
