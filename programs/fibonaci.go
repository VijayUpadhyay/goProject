package main

import "fmt"

func main() {
	num := 7
	fmt.Println(fibonacci(num))
}

func fibonacci(num int) int {
	if num == 0 || num == 1 {
		return num
	}
	fmt.Println(num)
	return fibonacci(num-1) + fibonacci(num-2)
}

func test() {
	a := new([]int)
	//b:=make([]int)
	fmt.Print(a)
}
