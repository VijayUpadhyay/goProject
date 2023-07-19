package main

import (
	"bytes"
	"fmt"
)

func main() {
	slice1 := []byte{1, 2, 3, 4, 5}
	slice2 := []byte{1, 2, 3, 4, 5}
	slice3 := []byte{5, 4, 3, 2, 1}

	fmt.Println(bytes.Equal(slice1, slice2)) // true
	fmt.Println(bytes.Equal(slice1, slice3)) // false
	//
	fmt.Println(bytes.Compare(slice1, slice2)) // 0
	fmt.Println(bytes.Compare(slice1, slice3)) // -1
	//
	fmt.Println(bytes.EqualFold(slice1, slice2)) // true
	fmt.Println(bytes.EqualFold(slice1, slice3)) // false
}
