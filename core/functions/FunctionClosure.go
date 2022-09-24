package main

import "fmt"

func closureEx() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	p, q, r := closureEx(), closureEx(), closureEx()
	for i := 0; i < 10; i++ {
		fmt.Println("p is:", p(i), "q is:", q(i*2), "r is:", r(i*3))
	}
}
