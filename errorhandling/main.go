package main

import (
	"fmt"
	"strings"
)

// ============================================================================
// Example 1: DeferTest.go
// ============================================================================

func example1_defertest() {
	defer fmt.Println("1st defer")
	defer fmt.Println("2nd defer")
	defer fmt.Println("3rd defer")
	fmt.Println("main 1st")
	fmt.Println("main 2nd")
	fmt.Println("Going to call method without defer")
	deferTest()
}

func deferTest() {
	i := 0
	defer fmt.Println(i)
	i++
	fmt.Println(i)
	return
}

// ============================================================================
// Example 2: errorChk.go
// ============================================================================

var err = "err msg"

func example2_errorchk() {
	if errr := returnError(); errr != nil {
		fmt.Printf("failed with error: %+v", errr)
		return
	}

}

func returnError() error {
	return fmt.Errorf("Error is: %s", err)
}

// ============================================================================
// Example 3: PanicAndDeferTest.go
// ============================================================================

func example3_panicanddefertest() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}

// ============================================================================
// Main Function - Run All Examples
// ============================================================================

func main() {
	fmt.Println("Running 3 examples from errorhandling")
	fmt.Println(strings.Repeat("=", 60))

	fmt.Println("\n--- example1_defertest ---")
	example1_defertest()

	fmt.Println("\n--- example2_errorchk ---")
	example2_errorchk()

	fmt.Println("\n--- example3_panicanddefertest ---")
	example3_panicanddefertest()

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("All examples completed!")
}
