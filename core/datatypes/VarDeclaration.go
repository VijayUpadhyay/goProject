package main

import "fmt"

func main() {
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
