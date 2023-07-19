package main

import (
	"fmt"
	"unsafe"
)

func main() {
	map1 := make(map[string]string)
	fmt.Println("Initial map size: ", unsafe.Sizeof(map1))
	map1["a"] = "b"
	fmt.Println("After 1st entry, map size: ", unsafe.Sizeof(map1))
	map1["b"] = "c"
	fmt.Println("Initial 2nd entry, map size: ", unsafe.Sizeof(map1))
	map2 := make(map[string]string)
	map2["a1"] = "b1"
	map2["b1"] = "c1"
	map2 = map1
	fmt.Print(map2)
}
