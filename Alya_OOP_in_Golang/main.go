package main

import (
	"fmt"
)

const USD1UZS = 10500

type Bank struct {
	money float64
}

func New(money float64) *Bank {
	return &Bank{
		money: money,
	}
}

func Default() *Bank {
	return &Bank{
		money: 1,
	}
}

func (bank Bank) uzs() float64 {
	return USD1UZS * bank.money
}

func (bank Bank) usd() float64 {
	return bank.money / float64(USD1UZS)
}

func (bank Bank) tranzaction(cash float64, callback func(err string, data float64)) {
	fmt.Println("tranzaction start")
	if bank.money > cash {
		msg := bank.money - cash
		callback("", msg)
	} else {
		callback("xatolik", -1)
	}
	fmt.Println("tranzaction end")
}

func (bank Bank) info() interface{} {
	if bank.money > 15000 {
		return bank.money
	}

	return "less money"
}

func main() {

	hamkorbank := New(660) //usd da

	fmt.Println(hamkorbank.uzs())

	infinbank1 := Default() //sum

	fmt.Println(infinbank1.info())

	infinbank2 := New(1050000.0) //sum

	fmt.Println(infinbank1.usd())
	fmt.Println(infinbank2.usd())

	callback := func(err string, data float64) {
		if err == "" {
			fmt.Println(data)
		} else {
			fmt.Println("error not enough money")
		}
	}

	infinbank1.tranzaction(1, callback)
	infinbank1.tranzaction(1200, callback)
	infinbank2.tranzaction(50, callback)

}
