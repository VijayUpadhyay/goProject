package main

import (
	"fmt"
	"math"
)

var a, b, c bool
var d int = 10

func main() {
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
