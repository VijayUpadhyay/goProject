package main

import "fmt"

func main() {
	// Create a channel with a capacity of 3
	ch := make(chan int, 3)
	// Send the value 2 to the channel
	ch <- 2
	// Call the process function with the channel as an argument
	process(ch)
	// Print a new line
	fmt.Println()
}

func process(ch <-chan int) {
	// Receive a value from the channel and assign it to the variable s
	s := <-ch
	// Print the value of s
	fmt.Println(s)
	// Comment out the line below to prevent sending it to a read-only channel
	//ch <- 2 // ===========================================================================================> NOT ALLOWED
}
