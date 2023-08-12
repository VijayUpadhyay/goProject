package main

import (
	"fmt"
	"time"
)

func sayHello(str string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("str at %d is: %s \n", i+1, str)
	}
}
func main() {
	fmt.Println()
	go sayHello("hello")
	sayHello("hi")
}
