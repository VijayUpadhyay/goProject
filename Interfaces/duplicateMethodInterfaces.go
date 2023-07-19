package main

import (
	"fmt"
)

type StringCheck interface {
	Println()
}

type Reader interface {
	Read(p []byte) (n int, err error)
	Close()
}

type Writer interface {
	Write(p []byte) (n int, err error)
	Close() bool
}

type ReadWriter interface {
	Reader
	// Writer --> duplicate methods are not allowed
}

func main() {
	fmt.Println("Hello, playground")
}
