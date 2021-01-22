package main

import "fmt"

type Hayvonlar interface {
	Goshxor()
}

type Tuzilishi struct {
	Ovozi string
	Oyogi int
	Rangi string
}

func (t *Tuzilishi) Goshxor() {
	fmt.Println(t.Ovozi, t.Oyogi, t.Rangi)
}

type Uchadimi bool

func (uchadimi Uchadimi) Goshxor() {
	fmt.Println(uchadimi)
}

func main() {
	var it Hayvonlar

	it = &Tuzilishi{
		"vov vov",
		4,
		"qora",
	}

	describe(it)
	it.Goshxor()

}

func describe(xayvonlar Hayvonlar) {
	fmt.Printf("(%v, %T)\n", xayvonlar, xayvonlar)
}
