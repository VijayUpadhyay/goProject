package main

import (
	"fmt"
	"time"
)

func main() {
	v := make(chan int)
	select {
	case <-v:
		fmt.Println("inside v")
	case <-time.After(1 * time.Minute):
		fmt.Println("Inside time limit check")
	}
}
