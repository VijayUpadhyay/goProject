package main

import (
	"fmt"
	"math"
	"strings"
)

// ============================================================================
// Example 1: FunctionClosure.go
// ============================================================================

func closureEx() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func example1_functionclosure() {
	p, q, r := closureEx(), closureEx(), closureEx()
	for i := 0; i < 10; i++ {
		fmt.Println("p is:", p(i), "q is:", q(i*2), "r is:", r(i*3))
	}
}

// ============================================================================
// Example 2: FunctionClosureTest.go
// ============================================================================

func example2_functionclosuretest() {
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

// ============================================================================
// Example 3: FunctionValues.go
// ============================================================================

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func example3_functionvalues() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

// ============================================================================
// Example 4: NakedReturn.go
// ============================================================================

func example4_nakedreturn() {
	fmt.Println(nakedReturn(100))
}

func nakedReturn(sum int) (x, y int) {
	x = sum / 10
	y = sum - 10
	return
}

// ============================================================================
// Example 5: Variadic.go
// ============================================================================

func example5_variadic() {
	variadicExample("red", "blue", "green", "yellow")
}

func variadicExample(s string, t ...string) {
	fmt.Println(s[0])
	fmt.Println(t[2])
}

// ============================================================================
// Main Function - Run All Examples
// ============================================================================

func main() {
	fmt.Println("Running 5 examples from functions")
	fmt.Println(strings.Repeat("=", 60))

	fmt.Println("\n--- example1_functionclosure ---")
	example1_functionclosure()

	fmt.Println("\n--- example2_functionclosuretest ---")
	example2_functionclosuretest()

	fmt.Println("\n--- example3_functionvalues ---")
	example3_functionvalues()

	fmt.Println("\n--- example4_nakedreturn ---")
	example4_nakedreturn()

	fmt.Println("\n--- example5_variadic ---")
	example5_variadic()

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("All examples completed!")
}
