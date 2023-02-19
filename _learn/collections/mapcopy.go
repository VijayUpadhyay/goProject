package main

import "fmt"

func main() {
	map1 := make(map[string]string)
	map1["a"] = "b"
	map1["b"] = "c"
	map2 := make(map[string]string)
	map2["a1"] = "b1"
	map2["b1"] = "c1"
	map2 = map1
	fmt.Print(map2)

}
