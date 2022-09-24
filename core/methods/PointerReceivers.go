package main

import (
	"fmt"
	"math"
)

type Ver struct {
	a float64
	b float64
}
func (v *Ver) pointerReceiver() float64  {
	return math.Sqrt(v.a*v.a + v.b*v.b)
}

func (v *Ver) changeValueUsingPointer()  {
	v.a=3.0
	v.b=4.0
}

func funcWithoutPointer(v Ver)  {

}

func main() {
	v:= Ver{10.23,10.234}
	v.changeValueUsingPointer()
	funcWithoutPointer(v)
	//funcWithoutPointer(*v) -> won't work
	fmt.Println(v.pointerReceiver())
}


