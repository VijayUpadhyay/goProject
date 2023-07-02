package main

import (
	"fmt"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)
	//c3 <- "test"
	fmt.Println(c3)
	go func() {
		c1 <- "inside channel 1"
	}()

	go func() {
		c2 <- "inside channel 2"
		c3 <- "test"
	}()
	fmt.Println(c3)
	for i:=0;i<2 ;i++  {
		select {
		case msg1:= <- c1:
			fmt.Println(msg1)
		case msg2:= <- c2:
			fmt.Println(msg2)
		case msg3 := <-c3:
			fmt.Println(msg3)

		/*default:
			fmt.Println("No data in channels")*/
		}
	}
	fmt.Println("End")
}
