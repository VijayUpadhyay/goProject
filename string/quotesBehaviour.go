package main

import "fmt"

func main() {
	s := `Go\tJava\nPython`
	fmt.Println(s)
	s = "Go\tJava\nPython"
	fmt.Println(s)
}
