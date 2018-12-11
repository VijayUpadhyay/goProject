package Basics

import (
	"fmt"
)
func main() {
	var  x interface{}

	switch i:= x.(type) {
	case nil:
		fmt.Println("type of x is: %T",i)
	case int:
		fmt.Println("type of x is: int ")
	case float32:
		fmt.Println("type of x is float32")
	case float64:
		fmt.Println("type of x is float64")
	case string:
		fmt.Println("type of x is string")
	default:
		fmt.Println("Unknown type")
	}
	fmt.Println()
	}
