package main

import "fmt"

type Animal struct {
	sound string
}

type Dog struct {
	Animal
	Friendly bool
}

type Cat struct {
	Animal
	Annoying bool
}

func specs(a interface{}) {
	fmt.Println(a)
}
func main() {
	fido := Dog{
		Animal{"wov wov"},
		true,
	}
	fifi := Cat{Animal{"meow"}, true}

	specs(fido)
	specs(fifi)

	critters := []interface{}{fido, fifi}
	fmt.Println(critters)

}
