package main

import "fmt"

func main() {
	ch := make(chan int)
	go func(chan int) {
		ch <- 10
	}(ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch) // passed once, read once
	/*fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)*/
}
