package main

import "fmt"

func main() {
	var a [6]int = [6]int{0, 1, 2} // 0,1,2,0,0
	fmt.Println(len(a), cap(a), a)
	var t []int = []int{0, 1, 2} // 0,1,2,0,0
	fmt.Println(len(t), cap(t), t)
	fmt.Println("-------------------------------------")
	b := a[1:4]       //1,2,0
	b = append(b, 10) //b: 1,2,0,10, a: 0,1,2,0,0,10
	b = append(b, 10)
	fmt.Println("-------------------------------------")
	fmt.Println("a: ", len(a), cap(a), a)
	fmt.Println("b: ", len(b), cap(b), b)
	fmt.Println("-------------------------------------")
	b = append(b, 10)
	b = append(b, 10)
	fmt.Println("a: ", len(a), cap(a), a)
	fmt.Println("b: ", len(b), cap(b), b)
}
