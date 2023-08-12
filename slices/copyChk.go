package main

import "fmt"

func main() {
	fmt.Println("test1")
	test1()
	fmt.Println("test2")
	test2()
}

func test1() {
	a := []int{1, 2}
	b := []int{3, 4}
	check := a
	copy(a, b) // the check variable is used to hold a reference to the original slice description to make sure it is really copied. So copied to their reference variables also
	fmt.Println(a, b, check)
	// Output: [3 4] [3 4] [3 4]
}

func test2() {
	a := []int{1, 2}
	b := []int{3, 4}
	check := a
	a = b // operation does not copy the slice contents, only the slice description
	fmt.Println(a, b, check)
	// Output: [3 4] [3 4] [1 2]
}
