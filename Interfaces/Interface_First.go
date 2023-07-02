package main

import (
	"fmt"
	"math"
)

type TestInterface interface {
	abc() float64
}

type val float64

type vertexChk struct {
	x, y float64
}

func (f val) abc() float64 {
	return float64(f)
}

func (v *vertexChk) abc() float64 {
	return math.Sqrt(v.x*v.x + v.y + v.y)
}

func main() {
	var a TestInterface
	ab := val(-math.Sqrt2)
	m := vertexChk{3, 4}
	a = ab
	fmt.Println("After a = ab, value of a is:", a)
	a = &m
	fmt.Println("After a=&m, value of a is:", a)
	//a=m --> it will work if we get a method having non-pointer receiver or remove * from "*Vertex"
}
