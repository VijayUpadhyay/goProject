package main

import (
	"fmt"
	"math"
	"strings"
)

// ============================================================================
// Example 1: MethodAndPointerTest.go
// ============================================================================

type Vertx struct {
	X, Y float64
}

func (v *Vertx) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertx, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func example1_methodandpointertest() {
	v := Vertx{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertx{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}

// ============================================================================
// Example 2: MethodsTest.go
// ============================================================================

type Vertex struct {
	a float64
	b float64
}

func (v Vertex) test() float64 {
	return math.Sqrt(v.a*v.a + v.b*v.b)
}
func example2_methodstest() {
	v := Vertex{3.0, 4.0}
	fmt.Println(v.test())
}

// ============================================================================
// Example 3: MethodsWithoutStruct.go
// ============================================================================

type floatValue float64

func example3_methodswithoutstruct() {
	value := floatValue(10.909789999999999999999999)
	fmt.Println(value.withoutStruct())
}

func (f floatValue) withoutStruct() float64 {
	return float64(f)
}

// ============================================================================
// Example 4: PointerReceivers.go
// ============================================================================

type Ver struct {
	a float64
	b float64
}

func (v *Ver) pointerReceiver() float64 {
	return math.Sqrt(v.a*v.a + v.b*v.b)
}

func (v *Ver) changeValueUsingPointer() {
	v.a = 3.0
	v.b = 4.0
}

func funcWithoutPointer(v Ver) {

}

func example4_pointerreceivers() {
	v := Ver{10.23, 10.234}
	v.changeValueUsingPointer()
	funcWithoutPointer(v)
	//funcWithoutPointer(*v) -> won't work
	fmt.Println(v.pointerReceiver())
}

// ============================================================================
// Main Function - Run All Examples
// ============================================================================

func main() {
	fmt.Println("Running 4 examples from methods")
	fmt.Println(strings.Repeat("=", 60))

	fmt.Println("\n--- example1_methodandpointertest ---")
	example1_methodandpointertest()

	fmt.Println("\n--- example2_methodstest ---")
	example2_methodstest()

	fmt.Println("\n--- example3_methodswithoutstruct ---")
	example3_methodswithoutstruct()

	fmt.Println("\n--- example4_pointerreceivers ---")
	example4_pointerreceivers()

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("All examples completed!")
}
