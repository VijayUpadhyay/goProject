package main

import "fmt"

type ByteCounter int

func (b *ByteCounter) Write(p []byte) (n int, err error) {
	//*b = ByteCounter(len(p))
	*b = ByteCounter(2000)
	return 2000, nil
}

func main() {
	var b ByteCounter
	fmt.Fprintf(&b, "Vijay Upadhyay")
	fmt.Println(b)
}
