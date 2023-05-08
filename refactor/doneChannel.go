package main

import (
	"fmt"
)

func main() {
	done := make(chan bool)
	//done <- false
	go isDone(done)
	<-done
	fmt.Println("completed")
}

func isDone(done chan bool) {
	//time.Sleep(3 * time.Second)
	done <- true
}
