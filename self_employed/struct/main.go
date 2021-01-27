package main

import "fmt"

type foo int

type Person struct {
	fname  string
	lname  string
	age    int
	salary float64
	company
}

type company struct {
	comp_id int
	cname   string
}

func (p Person) fullname() string {
	return p.fname + "|" + p.lname
}

func (p Person) info() string {
	return p.lname + "|" + p.fname + "|"
}

func (c company) info() string {
	return c.cname
}

func main() {
	p1 := &Person{
		"sarrsdf",
		"adds",
		21,
		1231321321.1,
		company{
			2,
			"adjdkj",
		},
	}

	fmt.Println(p1)
	fmt.Printf("%T\n", p1)
	fmt.Println(p1.age)
	fmt.Println(p1.company.comp_id)
)

	a := Person{
		"Sardor",
		"Buvashve",
		25,
		1260000.0,
		company{
			1,
			"UzXIA",
		},
	}

	fmt.Println(a)
	fmt.Println(a.fullname())
	fmt.Printf("%T %v \n", a, a)

	fmt.Println(a.info())
	fmt.Println(a.company.info())

}
