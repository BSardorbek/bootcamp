package main

import (
	"fmt"
	"strconv"
)

type Contact struct {
	id          int
	firstName   string
	lastName    string
	phoneNumber int16
}

func (c Contact) getFirstName() string {
	return c.firstName
}

func (c Contact) getLastName() string {
	return c.lastName
}

func (c Contact) getID() int {
	return c.id
}

func (c *Contact) setFirstName(firstname string) {
	c.firstName = firstname
}

func (c *Contact) setLastName(lastname string) {
	c.lastName = lastname
}

func (c *Contact) toString() string {
	return c.getFirstName() + " " + c.getLastName() + " " + strconv.Itoa(c.getID())
}

func main() {

	var newContact Contact

	fmt.Println(newContact)

	newContact.setLastName("sardor")
	newContact.setFirstName("buvashev")
	newContact.id = 1
	fmt.Println(newContact)

	fmt.Println(newContact.toString())
	newContact.setFirstName("testtt")
	fmt.Println(newContact.toString())

}
