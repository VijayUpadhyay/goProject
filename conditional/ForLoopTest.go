package main

import (
	"fmt"
)

func main() {
	var sum = 0
	fmt.Println(sum)
	for sum < 1000 {
		sum += 100
		fmt.Println(sum)
	}
}
