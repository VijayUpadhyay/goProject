package main

import (
	"fmt"
	"reflect"
)

type mobile struct {
	price float64
	color string
}

func main() {
	// Slice DeepEqual - String
	s1 := []string{"A", "B", "C", "D", "E"}
	s2 := []string{"D", "E", "F"}
	s3 := []string{"A", "B", "C", "D", "E"}
	result := reflect.DeepEqual(s1, s2)
	fmt.Println("S1 and S2: ", result)
	result = reflect.DeepEqual(s1, s3)
	fmt.Println("S1 and S3: ", result)
	// Slice DeepEqual - int
	s4 := [3]int{1, 2, 3}
	s5 := [3]int{4, 2, 3}
	s6 := [3]int{1, 2, 3}
	result = reflect.DeepEqual(s4, s5)
	fmt.Println("S4 and S5: ", result)
	result = reflect.DeepEqual(s4, s6)
	fmt.Println("S4 and S6: ", result)

	// DeepEqual is used to check two arrays are equal or not
	n1 := [5]int{1, 2, 3, 4, 5}
	n2 := [5]int{1, 2, 3, 4, 5}
	result = reflect.DeepEqual(n1, n2)
	fmt.Println(result)

	// DeepEqual is used to check two structures are equal or not
	m1 := mobile{500.50, "red"}
	m2 := mobile{400.50, "black"}
	m3 := mobile{500.50, "red"}
	result = reflect.DeepEqual(m1, m2)
	fmt.Println(result)
	result = reflect.DeepEqual(m1, m3)
	fmt.Println(result)
}
