package main

import (
	"fmt"
	"io"
)

func main() {
	printAny((*int)(nil))
	printAny[*int](nil)
	printInterface((*int)(nil))

	var r io.Reader
	printAny(r)
	printInterface(r)
}

func printInterface[T any](foo T) {
	fmt.Printf("%v\n", foo)
}

func printAny[T interface{}](foo T) {
	fmt.Printf("%v\n", foo)
}
