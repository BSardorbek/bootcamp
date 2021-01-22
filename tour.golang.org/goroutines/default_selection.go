package main

import (
	"fmt"
	"time"
)

func main() {
	quyon1 := time.Tick(1000 * time.Millisecond)
	quyon2 := time.After(5000 * time.Millisecond)

	x := 123

	// fmt.Println(<-quyon1) //1sekund turadi
	// fmt.Println(<-quyon2) //5 sekund turadi

	fmt.Println(x)

	for {
		select {
		case <-quyon1:
			fmt.Println("quton1") //1sekund turadi
		case <-quyon2:
			fmt.Println("quyon2") //5sekund turadi
			return
		default:
			fmt.Println("kutish vaqtida ishlovchi jarayon") //agar yuqorida kutish vaqti cho'ilib ketsa default holat ishlab turadi
			time.Sleep(6000 * time.Millisecond)
		}
	}

}
