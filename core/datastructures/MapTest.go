package main

import "fmt"

type VarName struct {
	a int
	b int
}

type VarName2 struct {
	a1 string
	b1 string
}

var m map[string]VarName
var p2 map[int64]VarName2

func main() {
m = make(map[string]VarName)
m["key"]=VarName{10,20}

p2 = make(map[int64]VarName2)
p2[10]=VarName2{"p","q"}
p2[20]=VarName2{"dfvdf","dvd"}
	fmt.Println(m["key"])
	fmt.Println(m)

	fmt.Println("=======================")
	fmt.Println(p2)
	fmt.Println(p2[10])
	fmt.Println(p2[30])
}
