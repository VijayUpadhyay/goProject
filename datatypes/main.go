package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

// ============================================================================
// Example 1: autoTypeCast.go
// ============================================================================

func example1_autotypecast() {
	var x = 13 / 3.5
	fmt.Println(x)
	fmt.Println(reflect.TypeOf(x))
}

// ============================================================================
// Example 2: nilChecks.go
// ============================================================================

func example2_nilchecks() {
	var a interface{}
	v, ok := a.(int)
	fmt.Println(v, ok)
	//fmt.Println(nil.(int))
	fmt.Println(a.(int))
}

// ============================================================================
// Example 3: VarDeclaration.go
// ============================================================================

func example3_vardeclaration() {
	var a, b, c = 3, 4, "vj"
	const LENGTH int = 10
	const WIDTH int = 20
	var area int
	p := 10
	q := "ram"
	var r float32 = 3.23333399999999999999 // round off at 6th pos
	fmt.Println(r)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(p)
	fmt.Println(q)
	fmt.Println()
	area = findArea(LENGTH, WIDTH)
	fmt.Println("Area is: ", area)
}

func findArea(length int, width int) int {
	return length * width
}

// ============================================================================
// Example 4: VariablesTest.go
// ============================================================================

var a, b, c bool
var d int = 10

func example4_variablestest() {
	var e float64
	var d int = 5
	fmt.Println(a, b, c, e)
	fmt.Println(d)

	fmt.Println("-----------------------------------")
	var m, n, o = 5, 5.5, "string data"
	fmt.Println(m, n, o)
	fmt.Println("Calling testMath")
	testMath()
}

func testMath() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	p := x * x
	fmt.Println("P is: ", p)
	var z uint = uint(f)
	fmt.Println(x, y, z)
	const (
		BIG   = 1 << 10
		SMALL = BIG >> 9
	)
	fmt.Println("Big value: ", testBig(BIG))
	fmt.Println("Small value: ", SMALL)

}

func testBig(a int64) int64 {
	return a*100 + 1
}

// ============================================================================
// Main Function - Run All Examples
// ============================================================================

func main() {
	fmt.Println("Running 4 examples from datatypes")
	fmt.Println(strings.Repeat("=", 60))

	fmt.Println("\n--- example1_autotypecast ---")
	example1_autotypecast()

	fmt.Println("\n--- example2_nilchecks ---")
	example2_nilchecks()

	fmt.Println("\n--- example3_vardeclaration ---")
	example3_vardeclaration()

	fmt.Println("\n--- example4_variablestest ---")
	example4_variablestest()

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("All examples completed!")
}
