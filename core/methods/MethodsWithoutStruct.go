package main

import (
	"fmt"
)

type floatValue float64
func main() {
	value:=floatValue(10.909789999999999999999999)
	fmt.Println(value.withoutStruct())
}

func (f floatValue) withoutStruct() float64 {
	return float64(f)
}
