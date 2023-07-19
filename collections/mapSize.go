package main

import (
	"fmt"
	"unsafe"
)

func main() {
	map1 := make(map[string]string)
	fmt.Println("Initial map size: ", unsafe.Sizeof(map1))
	for i := 0; i < 10000; i++ {
		map1[fmt.Sprintf("a%d", i)] = fmt.Sprintf("b%d", i)
		fmt.Println("After 1st entry, map size: ", unsafe.Sizeof(map1))
	}
}
