package main

import (
	"testing"
)

func TestToString(t *testing.T) {
	contact := Contact{
		1,
		"Sartdor",
		"Buvashev",
		123,
	}
	fname := contact.getFirstName()
	if fname != "Sartdor" {
		t.Errorf("Error FName  ")
	}
}
