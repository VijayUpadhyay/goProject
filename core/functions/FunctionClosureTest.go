package main

import "fmt"

func main() {
	fmt.Println()
	test()
}

// in closure --> 2 return for each function closure
func test() func(int) int {

	return func(i int) int {
		sum := 0
		sum += sum + i
		return sum
	}
}
