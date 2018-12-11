package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("Starting go func")

	c1:= make(chan string)
	c2:= make(chan string)
	go func() {
		time.Sleep(time.Millisecond*3)
		c1<-"Hi"
	}()

	go func() {
		time.Sleep(time.Millisecond*2)
		c2<-"Vijay"
	}()

	for i:=0;i<2 ;i++  {
		select {
		case msg1:= <-c1 :
			fmt.Println(msg1)
		case msg2:= <-c2:
			fmt.Println(msg2)

			
		}
	}

}
