package main

import (
	"fmt"
	"math"
	"runtime"
)

func main() {
	var sum int = 100

	if p := math.Pow10(sum); p < 1000 {
		fmt.Println("p is less than 1000")
	} else {
		fmt.Println("p is equalTo or greater than 1000")
	}

	fmt.Println()

	os := runtime.GOOS
	fmt.Println("os is: ", os)
}
