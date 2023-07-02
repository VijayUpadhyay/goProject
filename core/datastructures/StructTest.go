package main

import (
	"fmt"
)

func main() {
	type VariableName struct {
		a   int
		str string
	}
	fmt.Println(VariableName{5, "madhu"})

	//z:= make([]struct{},5)
	v := VariableName{10, "Rakesh"}
	var d = VariableName{}
	fmt.Println(d)
	fmt.Println("Roll Number:", v.a, "Name:", v.str)

	fmt.Println("Changing data in sruct using pointer")
	p := &v
	fmt.Println("Value in p:", p)
	p.a = 900
	fmt.Println("Changed struct:", v)

	//Array of structures/union
	/*for i:=0;i<5 ;i++  {
		//z[i] :=
		fmt.Println(z[i])
	}*/
}
