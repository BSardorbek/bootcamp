package main

import (
	"fmt"
	"time"
)

func main() {

	//oddiy usul

	// a := "BAHOR"
	// a := "YOZ"
	a := "X"

	switch a {
	case "QISH":
		fmt.Println("qor yogadi")
	case "BAHOR":
		fmt.Println("sumalak pishiriladi")
	case "YOZ":
		fmt.Println("mazza qilib daryoda suzamiz")
	case "KUZ":
		fmt.Println("daraxt barklari to'kiladi")
	default:
		fmt.Println("Bunday fasl yo'q")

	}

	//2-usul
	switch time.Now().Weekday() {
	case time.Sunday, time.Saturday, time.Thursday:
		fmt.Println("dam olish kuni")
	default:
		fmt.Println("o'qish  kuni")

	}

	//3-usul
	t := time.Now().Hour()
	switch {
	case t > 16:
		fmt.Println("kech qolding")
	default:
		fmt.Println("ayni vaqti")

	}

	//4-usul

	baxo := func(a int) {
		switch a {
		case 2:
			fmt.Println("Yomon")
		case 3:
			fmt.Println("Qoniqarli")
		case 4:
			fmt.Println("O'rtacha")
		case 5:
			fmt.Println("A'lo")
		default:
			fmt.Println("Bebaxo)")
		}
	}

	baxo(1)
	baxo(4)
	baxo(3)
	baxo(5)

	defer fmt.Println("defer  --> Bu barcha jarayon tugagaidan so'ng ishlaydi")
}
