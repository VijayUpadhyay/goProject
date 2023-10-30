package main

import "fmt"

type Staff struct {
	Name string
	Id   int32
}

type Student struct {
	Staff
	Subjects []string
}

type Faculty struct {
	Staff
	Subjects []string
}

func main() {

	s := Student{Staff{Name: "Ram", Id: 100}, []string{"Maths", "Science"}}
	fmt.Println(s)

}
