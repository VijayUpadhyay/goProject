package main

import "fmt"

func main() {
	//s1 := new([]int)
	// fmt.Printf("Slice len: %d and Cap: %d", len(*s1), cap(*s1))
	arr := [6]int{1, 2, 3, 4, 5, 6}
	//s1 := make([]int, 6)
	s1 := arr[2:5]
	fmt.Printf("Slice: %+v len: %d and Cap: %d, Address: %p \n", s1, len(s1), cap(s1), &s1)
	s1 = append(s1, 10)
	fmt.Println(s1, arr)
	s1[3] = 1000
	fmt.Println(s1, arr)
	s1 = append(s1, 100) // once capacity is full, new slice will be created and all elements will be copied into it. From here, no link between slice and array
	fmt.Println(s1, arr)
	fmt.Printf("Now, Slice: %+v len: %d and Cap: %d, Address: %p \n", s1, len(s1), cap(s1), &s1)
	arr[3] = 200
	s1[1] = 20000
	fmt.Println(s1, arr)
}
