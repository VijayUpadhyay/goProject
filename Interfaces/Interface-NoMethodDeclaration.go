package main

import "fmt"

type I interface {
	M()
}

type tVertex struct {
	str string
}

func (v tVertex) M() {
	fmt.Println("inside M()", v.str)
}

func main() {
	var i I
	l := tVertex{"hello"}
	i = l
	describe(i)
	i.M()
}




func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
