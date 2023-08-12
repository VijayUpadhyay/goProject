package main

import "fmt"

func main() {
	a := make([]int, 6)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Println(a)
	a[0] = 100
	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Println(a)
	/**/
	fmt.Println("Going to add value at indexes")
	for i := 0; i < 6; i++ {
		a[i] = (i + 1) * 100
		fmt.Println(len(a))
		fmt.Println(cap(a))
		fmt.Println(a)
	}
	a = append(a, 90000)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Println(a)
}
