package main

import (
	"fmt"
)

func main() {
	vals := make([]int, 2)
	fmt.Println("Capacity was:", cap(vals))
	fmt.Println("Length was:", len(vals))
	for i := 0; i < 5; i++ {
		vals = append(vals, i)
		fmt.Println("Capacity is now:", cap(vals))
		fmt.Println("Length is now:", len(vals))
	}
	fmt.Println(vals)
	vals = append(vals, 100)
	fmt.Println("1- Recent capacity: ", cap(vals))
	fmt.Println("1- Recent length: ", len(vals))
	vals = append(vals, 100)
	fmt.Println("2- Recent capacity: ", cap(vals))
	fmt.Println("2- Recent length: ", len(vals))

	fmt.Println("========================================")
	test()
}

func test() {
	text := "val: "
	//n := 10
	/* text += fmt.Sprintf("%d. %s\n", n, scanner.Text())
	text = append([]byte(text), fmt.Sprintf("%d. %s\n", n, scanner.Text())...) */
	text += "vijay"
	fmt.Println(text)
}
