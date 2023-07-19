package main

import "fmt"

func main() {
	originalArr := [9]int{1, 2, 3, 22, 33, 11, 9, 909, 6576} // mentioning/ not mentioning size makes it array else it is slice
	fmt.Printf("Type is %T \n", originalArr)
	var sliceArr []int
	sliceArr = originalArr[:4]
	fmt.Println("originalArr1: ", originalArr)
	fmt.Println("sliceArr1: ", sliceArr)
	sliceArr[2] = 100
	fmt.Println("originalArr2: ", originalArr)
	fmt.Println("sliceArr2: ", sliceArr)
	for p := 0; p < 3; p++ {
		sliceArr = append(sliceArr, p)
	}
	fmt.Println("originalArr3: ", originalArr)
	fmt.Println("sliceArr3: ", sliceArr)
	for p := 0; p < 5; p++ {
		sliceArr = append(sliceArr, p)
	}
	fmt.Println("originalArr4: ", originalArr)
	fmt.Println("sliceArr4: ", sliceArr)
	sliceArr[1] = 500
	fmt.Println("originalArr5: ", originalArr)
	fmt.Println("sliceArr5: ", sliceArr)
}
