package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)
	go func(chan int) {
		for i := 0; i < 2; i++ {
			ch <- 10
			time.Sleep(3 * time.Second)
			ch <- 20
		}
	}(ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	//TODO: check why below 2 lines are working fine?
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	//fmt.Println(<-ch) // this will cause deadlock
}
