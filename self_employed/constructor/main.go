package main

import "fmt"

type Person struct {
	Name string
	age  int
}

func New(name string, age int) *Person {
	return &Person{
		Name: name,
		age:  age,
	}
}

func Default() *Person {
	return &Person{
		Name: "default",
		age:  18,
	}
}

func testt(i interface{}) {
	fmt.Println(i)
}

func main() {

	sardor := New("sardor", 26)

	navruz := Default()

	testt(navruz)
	testt(sardor)
	arr := []int{1, 2}
	testt(arr)

	fmt.Println(sardor)
	fmt.Println(navruz)

}
