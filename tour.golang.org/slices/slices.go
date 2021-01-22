package main

import (
	"fmt"
	"strings"
)

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)

	names := [...]string{
		"sardor",
		"navruz",
		"doston",
		"gishmat",
		"toshmat",
	}

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "xxx"
	fmt.Println(names)

	r := []bool{true, true, false, true}
	fmt.Println(r)

	ss := []struct {
		id     int
		isShow bool
	}{
		{1, true},
		{2, false},
		{3, true},
	}

	fmt.Println(ss)

	sonlar := []int{2, 3, 5, 7, 11, 13}
	slicess := sonlar[1:4]
	fmt.Println(slicess)
	fmt.Println(sonlar[1:])
	fmt.Println(sonlar[:2])

	printSlices(sonlar)
	printSlices(sonlar[:0])
	printSlices(sonlar[:4])
	printSlices(sonlar[2:])
	printSlices(sonlar[2:3])

	var isNiL []int
	var isOpen [5]bool
	println()
	isOpen[0] = true
	fmt.Println(" massiv ", isOpen)

	fmt.Println(isNiL, len(isNiL), cap(isNiL))
	if isNiL == nil {
		fmt.Println("bo'sh massiv nil ga teng")
	}

	aa := make([]int, 5)
	fmt.Println(aa, len(aa), cap(aa))

	bb := make([]int, 0, 5)
	fmt.Println(bb, len(bb), cap(bb))

	cc := bb[:2]
	fmt.Println(cc, len(cc), cap(cc))

	dd := aa[3:3]
	fmt.Println(dd, len(dd), cap(dd))

	println()
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	fmt.Println(board, len(board), cap(board))

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	board[0][1] = "O"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	var test []int
	printSlices(test)

	test = append(test, 0)
	printSlices(test)

	test = append(test, 1)
	printSlices(test)

	test = append(test, 11, 12, 23)
	printSlices(test)
}

func printSlices(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

/*
	slices bu array ga ko'rsatkich desa ham bo'ladi
*/
