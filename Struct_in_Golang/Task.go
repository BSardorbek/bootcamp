package main

import (
	"fmt"
	"strconv"
)

//task
type Task struct {
	id    int
	title string
	body  string
}

func (t *Task) toString() string {
	return t.title + " " + t.body + " " + strconv.Itoa(t.id)
}

func main() {

	var newTask Task

	fmt.Println(newTask)

	newTask.title = "adjlksdj"
	newTask.body = "asdkldkl sdhkj"
	newTask.id = 1
	fmt.Println(newTask)

	fmt.Println(newTask.toString())

}
