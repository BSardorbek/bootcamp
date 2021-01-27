package main

import "fmt"

func main() {
	myStudent := map[int]string{
		1: "sardor",
		2: "navzur",
		3: "doston",
	}

	fmt.Println(myStudent)
	fmt.Println(len(myStudent))

	myStudent[2] = "xxxxx"

	fmt.Println(myStudent)
	delete(myStudent, 8)
	delete(myStudent, 3)
	fmt.Println(myStudent)

	if val, exists := myStudent[3]; exists {
		delete(myStudent, 3)
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	} else {
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	}

	txt := "sartdor"
	fmt.Println(txt)
	fmt.Printf("%T\n", txt[0])

	l := rune("As"[0])
	fmt.Printf("%T\n", l)
	fmt.Println(l)

	word := "Hello"
	letter := rune(word[0])
	fmt.Println(letter)
	fmt.Println(string(72))

}
