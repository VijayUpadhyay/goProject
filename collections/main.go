package main

import (
	"fmt"
	"strings"
	"unsafe"
)

// ============================================================================
// Example 1: mapcopy.go
// ============================================================================

func example1_mapcopy() {
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

// ============================================================================
// Example 2: mapSize.go
// ============================================================================

func example2_mapsize() {
	map1 := make(map[string]string)
	fmt.Println("Initial map size: ", unsafe.Sizeof(map1))
	for i := 0; i < 10000; i++ {
		map1[fmt.Sprintf("a%d", i)] = fmt.Sprintf("b%d", i)
		fmt.Println("After 1st entry, map size: ", unsafe.Sizeof(map1))
	}
}

// ============================================================================
// Main Function - Run All Examples
// ============================================================================

func main() {
	fmt.Println("Running 2 examples from collections")
	fmt.Println(strings.Repeat("=", 60))

	fmt.Println("\n--- example1_mapcopy ---")
	example1_mapcopy()

	fmt.Println("\n--- example2_mapsize ---")
	example2_mapsize()

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("All examples completed!")
}
