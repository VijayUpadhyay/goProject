package main

import (
	"fmt"
)

func main() {
	var a, b int = 100, 200
	fmt.Println("Value of a: ", a)
	fmt.Printf("%d", &a)
	address := &a
	fmt.Println()
	fmt.Println("Value from address: ", *(&a))
	fmt.Printf("Addess is: %d\n", address)
	fmt.Println("Value of a through address", *address)

	fmt.Println("-------------------------")
	*address = 45
	fmt.Println("Changed value of a:", a)
	address = &b
	*address = 500 // dereferencing
	fmt.Println("Changed value of b:", b)
}
