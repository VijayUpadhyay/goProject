package main

import (
	"fmt"
	"strings"
)

func main() {
	col := "F15_StopLight"
	v := strings.TrimSuffix(col, "_StopLight")
	fmt.Println(v)
}
