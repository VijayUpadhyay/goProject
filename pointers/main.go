package main

import (
	"fmt"
	"strings"
)

// ============================================================================
// Example 1: PointersTest.go
// ============================================================================

func example1_pointerstest() {
	var a, b int = 100, 200
	fmt.Println("Value of a: ", a)
	fmt.Printf("%d", &a)
	address := &a
	fmt.Println()
	fmt.Println("Value from address: ", *(&a))
	fmt.Printf("Addess is: %d\n", address)
	fmt.Println("Value of a through address", *address)

	fmt.Println("-------------------------")
	*address = 45
	fmt.Println("Changed value of a:", a)
	address = &b
	*address = 500 // dereferencing
	fmt.Println("Changed value of b:", b)
}

// ============================================================================
// Example 2: ptrValueAssignment.go
// ============================================================================

func example2_ptrvalueassignment() {
	// var v string=nil ==> not allowed, Nil is not a type but a reserved word, you cannot use it in assignment.
	var data *string = nil
	if data == nil {
		fmt.Println(data)
	}
}

// ============================================================================
// Main Function - Run All Examples
// ============================================================================

func main() {
	fmt.Println("Running 2 examples from pointers")
	fmt.Println(strings.Repeat("=", 60))

	fmt.Println("\n--- example1_pointerstest ---")
	example1_pointerstest()

	fmt.Println("\n--- example2_ptrvalueassignment ---")
	example2_ptrvalueassignment()

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("All examples completed!")
}
