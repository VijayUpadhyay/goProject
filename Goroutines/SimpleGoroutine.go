package main

import (
	"fmt"
	"time"
)

func sayHello(str string)  {
	for i:=0;i<5 ;i++  {
		time.Sleep(100*time.Millisecond)
		fmt.Println("str is: ",str)
	}
}
func main()  {
	fmt.Println()
	go sayHello("hello")
	sayHello("hi")
}
