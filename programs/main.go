package main

import (
	"fmt"
	"strings"
)

// ============================================================================
// Example 1: factorial.go
// ============================================================================

func example1_factorial() {
	num := 7
	fmt.Println(factorial(num))
}

func factorial(num int) int {
	if num == 0 || num == 1 {
		return 1
	}
	return num * factorial(num-1)
}

// ============================================================================
// Example 2: fibonaci.go
// ============================================================================

func example2_fibonaci() {
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

// ============================================================================
// Example 3: swapList.go
// ============================================================================

func example3_swaplist() {
	listObj := [3]int{1, 2, 3}
	listObj1 := []int{1, 2, 66, 3}
	swapContents(listObj)
	swapContents1(listObj1)
	fmt.Println(listObj)
	fmt.Println(listObj1)
}

func swapContents(listObj [3]int) {
	if len(listObj) == 0 {
		return
	}
	for i, j := 0, len(listObj)-1; i <= j; i, j = i+1, j-1 {
		listObj[i], listObj[j] = listObj[j], listObj[i]
	}
}

func swapContents1(listObj []int) {
	if len(listObj) == 0 {
		return
	}
	for i, j := 0, len(listObj)-1; i < j; i, j = i+1, j-1 {
		listObj[i], listObj[j] = listObj[j], listObj[i]
	}
}

// ============================================================================
// Main Function - Run All Examples
// ============================================================================

func main() {
	fmt.Println("Running 3 examples from programs")
	fmt.Println(strings.Repeat("=", 60))

	fmt.Println("\n--- example1_factorial ---")
	example1_factorial()

	fmt.Println("\n--- example2_fibonaci ---")
	example2_fibonaci()

	fmt.Println("\n--- example3_swaplist ---")
	example3_swaplist()

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("All examples completed!")
}
