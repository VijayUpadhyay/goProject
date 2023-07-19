package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, p := range a {
		fmt.Printf("2**%d = %d\n", i, p)
	}
	fmt.Println("_________")
	for i := range a {
		fmt.Printf("%d\n", i)
	}
}
