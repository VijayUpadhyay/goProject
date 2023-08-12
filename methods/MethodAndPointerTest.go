package main

import "fmt"

type Vertx struct {
	X, Y float64
}

func (v *Vertx) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertx, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertx{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertx{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
