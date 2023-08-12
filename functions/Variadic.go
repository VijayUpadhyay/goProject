package main

import "fmt"

func main() {
	variadicExample("red", "blue", "green", "yellow")
}

func variadicExample(s string, t ...string) {
	fmt.Println(s[0])
	fmt.Println(t[2])
}
