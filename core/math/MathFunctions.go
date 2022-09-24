package main

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
)

func foo() {
	fmt.Println("Square root is: ", math.Sqrt(18))
	fmt.Println("Celeing value is:  ", math.Ceil(4.12))
	fmt.Println("Floor value is: ", math.Floor(4.12))
	fmt.Println("Random value is: ", rand.Intn(100))
	fmt.Println("Random value is: ", rand.Intn(100))
}

func add1(a, b float32) float32 {
	return a + b
}

func add2(a, b float64) float64 {
	return a + b
}

func single(a, b string) string {
	return a + b
}

func multiple1(a, b, c string) (string, string, string) {
	return a, b, c
}

func multiple(a, b string) (string, string) {
	return a, b
}

func typecastIt() {
	var a int = 64
	var b float64 = float64(a)
	fmt.Println(b)
}
func test3(a float32) {
	fmt.Printf("a is of type: %T\n", a)
}

func main() {
	foo()
	var v1, v2 float32 = 5.6, 5.6
	v3, v4 := 5.6, 6.7
	fmt.Println("Added value of v1 and v2 is: ", add1(v1, v2))
	//fmt.Println(fmt.Sprintf("%T"),v1)
	fmt.Println(reflect.TypeOf(v1).String())
	fmt.Println("Added value is: ", add2(v3, v4))

	fmt.Println("------------------------------")
	str1, str2 := "vijay", "upadhyay"
	str3, str4 := "test1", "test4"
	fmt.Println("Single return: ", single(str1, str2))
	fmt.Println("Multiple return is:")
	fmt.Println(multiple(str4, str3))
	fmt.Println("Calling typecast")
	typecastIt()
	test3(v2)
	rand.Seed(100)
	fmt.Printf("my random number is: %v \n", rand.Seed)
	fmt.Println(math.Pi)
}
