package main

import "fmt"

func main() {
	fmt.Println("Creating channels")
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pongs, pings)

	fmt.Println(<-pongs)

	////
	/* t:=make(chan string)
	pong() */
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}
func pong(pongs chan<- string, pings <-chan string) {
	msg := <-pings
	pongs <- msg
}
