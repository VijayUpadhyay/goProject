package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	a float64
	b float64
}

func (v Vertex) test() float64  {
	return math.Sqrt(v.a*v.a + v.b*v.b)
}
func main() {
	v:= Vertex{3.0,4.0}
	fmt.Println(v.test())
}
