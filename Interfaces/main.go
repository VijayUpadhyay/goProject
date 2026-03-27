package main

import (
	"fmt"
	"io"
	"math"
)

// ============================================================================
// Example 1: Basic Interface Implementation
// ============================================================================

type TestInterface interface {
	abc() float64
}

type val float64

type vertexChk struct {
	x, y float64
}

func (f val) abc() float64 {
	return float64(f)
}

func (v *vertexChk) abc() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func example1_BasicInterface() {
	fmt.Println("\n=== Example 1: Basic Interface ===")
	var a TestInterface
	ab := val(-math.Sqrt2)
	m := vertexChk{3, 4}
	a = ab
	fmt.Println("After a = ab, value of a is:", a)
	a = &m
	fmt.Println("After a=&m, value of a is:", a)
	// a=m --> it will work if we get a method having non-pointer receiver or remove * from "*Vertex"
}

// ============================================================================
// Example 2: Type Assertion
// ============================================================================

func example2_TypeAssertion() {
	fmt.Println("\n=== Example 2: Type Assertion ===")
	var i interface{} = "hello"

	p, ok2 := i.(int)
	fmt.Println("int assertion:", p, ok2)

	s := i.(string)
	fmt.Println("string assertion:", s)

	s, ok := i.(string)
	fmt.Println("string assertion with ok:", s, ok)

	f, ok := i.(float64)
	fmt.Println("float64 assertion with ok:", f, ok)

	// f = i.(float64) // This would panic - commented out for safety
	// fmt.Println(f)
}

// ============================================================================
// Example 3: Empty Interface
// ============================================================================

func example3_EmptyInterface() {
	fmt.Println("\n=== Example 3: Empty Interface ===")
	var i interface{}
	i = 50
	printDetails(i)
	i = "vijay"
	printDetails(i)
}

func printDetails(i interface{}) {
	fmt.Printf("value is: %v and type is: %T\n", i, i)
}

// ============================================================================
// Example 4: Type Switches
// ============================================================================

func example4_TypeSwitches() {
	fmt.Println("\n=== Example 4: Type Switches ===")
	var i interface{}
	typeCheck(i)
	i = 10
	typeCheck(i)
	i = "vijay"
	typeCheck(i)
	i = 79.00
	typeCheck(i)
}

func typeCheck(i interface{}) {
	switch val := i.(type) {
	case int:
		fmt.Printf("Type is int for: %v\n", val)
	case string:
		fmt.Println("Type is string for", val)
	case float64:
		fmt.Printf("Type is %T for: %v\n", val, val)
	default:
		fmt.Printf("Type for value %v is : %T\n", val, val)
	}
}

// ============================================================================
// Example 5: Stringer Interface
// ============================================================================

type Person struct {
	firstName string
	age       int
}

func (p Person) String() string {
	return fmt.Sprintf("%v is %d years", p.firstName, p.age)
}

func example5_StringerInterface() {
	fmt.Println("\n=== Example 5: Stringer Interface ===")
	v := Person{"vijay", 10}
	fmt.Println("Person:", v)
}

// ============================================================================
// Example 6: Custom String Method
// ============================================================================

type myInterface interface{}

type myStruct struct{}

func (ms myStruct) String() string {
	return "myStruct"
}

func example6_CustomString() {
	fmt.Println("\n=== Example 6: Custom String Method ===")
	var i myInterface = myStruct{}
	fmt.Println(i)
}

// ============================================================================
// Example 7: Polymorphism
// ============================================================================

type Triangle struct {
	base, height float32
}

type Square struct {
	length float32
}

type Rectangle struct {
	length, breadth float32
}

func (triangle Triangle) Area() float32 {
	return 0.5 * triangle.base * triangle.height
}

func (square Square) Area() float32 {
	return square.length * square.length
}

func (rect Rectangle) Area() float32 {
	return rect.length * rect.breadth
}

type Area interface {
	Area() float32
}

func example7_Polymorphism() {
	fmt.Println("\n=== Example 7: Polymorphism ===")
	triangleObject := Triangle{base: 20, height: 10}
	squareobject := Square{length: 25}
	rectObject := Rectangle{length: 15, breadth: 20}

	var shapeObject Area

	shapeObject = triangleObject
	fmt.Println("Triangle Area =", shapeObject.Area())

	shapeObject = squareobject
	fmt.Println("Square Area =", shapeObject.Area())

	shapeObject = rectObject
	fmt.Println("Rectangle Area =", shapeObject.Area())
}

// ============================================================================
// Example 8: Nil Interface
// ============================================================================

type TestNilInterface interface {
	checkNil()
}

type Vertex_2 struct {
	name string
}

func (v Vertex_2) checkNil() {
	fmt.Println("Inside check nil method")
	fmt.Println(v.name)
}

func example8_NilInterface() {
	fmt.Println("\n=== Example 8: Nil Interface ===")
	var TestInt_2 TestNilInterface = Vertex_2{"ram"}
	fmt.Println("1st call - method reference:")
	fmt.Printf("%v\n", TestInt_2.checkNil)
	fmt.Println("2nd call - method execution:")
	TestInt_2.checkNil()
}

// ============================================================================
// Example 9: Nil Interface Reference Call (Commented - would panic)
// ============================================================================

type Ij interface {
	M()
}

func example9_NilRefCall() {
	fmt.Println("\n=== Example 9: Nil Interface Reference ===")
	var i Ij
	describe2(i)
	// i.M() // This would panic - commented out for safety
	fmt.Println("Note: Calling M() on nil interface would panic")
}

func describe2(i Ij) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// ============================================================================
// Example 10: Implicit Implementation
// ============================================================================

type InterfaceTest interface {
	Test() string
	M()
}

type StructTest struct {
	firstName, lastName string
}

func (v StructTest) Test() string {
	return v.firstName + " " + v.lastName
}

func (t StructTest) M() {
	fmt.Println("Inside method M")
	fmt.Println(t.firstName)
}

func example10_ImplicitImplementation() {
	fmt.Println("\n=== Example 10: Implicit Implementation ===")
	var i InterfaceTest
	l := StructTest{"vijay", "upadhyay"}
	i = l
	fmt.Printf("Method reference: %v\n", i.M)
	fmt.Println("Full name:", i.Test())
}

// ============================================================================
// Example 11: Duplicate Method Interfaces
// ============================================================================

type StringCheck interface {
	Println()
}

type Reader interface {
	Read(p []byte) (n int, err error)
	Close()
}

type Writer interface {
	Write(p []byte) (n int, err error)
	Close() bool
}

type ReadWriter interface {
	Reader
	// Writer --> duplicate methods are not allowed (Close has different signatures)
}

func example11_DuplicateMethods() {
	fmt.Println("\n=== Example 11: Duplicate Method Interfaces ===")
	fmt.Println("Note: Interfaces cannot embed other interfaces with conflicting method signatures")
	fmt.Println("Reader.Close() and Writer.Close() bool would conflict")
}

// ============================================================================
// Example 12: Generic Interface Check
// ============================================================================

func example12_GenericCheck() {
	fmt.Println("\n=== Example 12: Generic Interface Check ===")
	printAny((*int)(nil))
	printAny[*int](nil)
	printInterface((*int)(nil))

	var r io.Reader
	printAny(r)
	printInterface(r)
}

func printInterface[T any](foo T) {
	fmt.Printf("%v\n", foo)
}

func printAny[T interface{}](foo T) {
	fmt.Printf("%v\n", foo)
}

// ============================================================================
// Example 13: Interface Without Method Declaration
// ============================================================================

type I interface {
	M()
}

type tVertex struct {
	str string
}

func (v tVertex) M() {
	fmt.Println("inside M()", v.str)
}

func example13_NoMethodDeclaration() {
	fmt.Println("\n=== Example 13: Interface Without Explicit Declaration ===")
	var i I
	l := tVertex{"hello"}
	i = l
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// ============================================================================
// Example 14: Stringer Implementation Over IPAddr
// ============================================================================

type IPAddr [4]byte

func (m IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", m[0], m[1], m[2], m[3])
}

func example14_StringerIPAddr() {
	fmt.Println("\n=== Example 14: Stringer Implementation Over IPAddr ===")
	var hosts = map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	fmt.Println("Loopback:", hosts["loopback"])
	fmt.Println("Google DNS:", hosts["googleDNS"])
}

// ============================================================================
// Example 15: Abser Interface Test
// ============================================================================

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func example15_AbserInterface() {
	fmt.Println("\n=== Example 15: Abser Interface ===")
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f // MyFloat implements Abser
	fmt.Println("MyFloat Abs:", a.Abs())

	a = &v // *Vertex implements Abser
	fmt.Println("Vertex Abs:", a.Abs())

	// a = v // This would fail - Vertex (not *Vertex) doesn't implement Abser
}

// ============================================================================
// Main Function - Run All Examples
// ============================================================================

func main() {
	fmt.Println("=================================================")
	fmt.Println("Go Interfaces - Comprehensive Examples")
	fmt.Println("=================================================")

	example1_BasicInterface()
	example2_TypeAssertion()
	example3_EmptyInterface()
	example4_TypeSwitches()
	example5_StringerInterface()
	example6_CustomString()
	example7_Polymorphism()
	example8_NilInterface()
	example9_NilRefCall()
	example10_ImplicitImplementation()
	example11_DuplicateMethods()
	example12_GenericCheck()
	example13_NoMethodDeclaration()
	example14_StringerIPAddr()
	example15_AbserInterface()

	fmt.Println("\n=================================================")
	fmt.Println("All interface examples completed!")
	fmt.Println("=================================================")
}

// Made with Bob
