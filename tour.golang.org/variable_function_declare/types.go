// asosiy turlar (basic types)

package main

import "fmt"

func main() {

	// integer
	var X uint8 = 255 //1<<8-1 dan oshib ketsa 0 qaytaradi
	var x int8 = 127  //-128 <-> 127
	fmt.Println(X+1, X, x)
	var Y int16 = 32767 //-32768<->32767
	fmt.Println(Y, Y+1)

	// float
	a := 20.45
	b := 34.89
	var c float32 = float32(b - a)
	fmt.Printf("Natija: %f", c)
	fmt.Printf("\n c ning turi : %T", c)

	println()

	// complex
	var d complex128 = complex(6, 2)
	var e complex64 = complex(9, 2)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Printf("d ning turi %T va e ning turi %T", d, e)

	println()
	// boolean
	str1 := "BootCamp"
	str2 := "bootCamp"
	str3 := "bootcamp"
	result1 := str1 == str2
	result2 := str1 == str3
	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Printf("result1 ning turi %T va result2 ning turi %T", result1, result2)

	// string
	str := "Qanaqadur matnlar"
	fmt.Printf("\nstr matnning uzunligi %d", len(str))
	fmt.Printf("\nMatn: %s", str)
	fmt.Printf("\nstr ning turi: %T", str)
	//constanta
	const PI = 3.14
	fmt.Printf("\nconst ning turi: %T", PI)

}
