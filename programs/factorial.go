package main

import "fmt"

func main() {
	num := 7
	fmt.Println(factorial(num))
}

func factorial(num int) int {
	if num == 0 || num == 1 {
		return 1
	}
	return num * factorial(num-1)
}
