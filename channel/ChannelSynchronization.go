package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting sync")
	done := make(chan bool)
	go worker(done)
	<-done

}

func worker(c chan bool) {
	fmt.Println("Inside method....")
	time.Sleep(time.Millisecond)
	fmt.Println("done")

	c <- true
}
