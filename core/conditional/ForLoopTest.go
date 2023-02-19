package main

import (
	"fmt"
)

func main() {
	var sum int = 0
	fmt.Println()
	for sum < 1000 {
		sum += 100
		fmt.Println(sum)
	}
}
