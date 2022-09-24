package main

import (
	"fmt"
	"strings"
	"io"
)

func main()  {
	r:= strings.NewReader("Vijay Upadhyay")
	b:= make([]byte,3)

	fmt.Println()

	for  {
		n,err:= r.Read(b)
		fmt.Printf("n: %v, err: %v, b: %v, ",n,err,b)
		fmt.Printf(" b[:n] = %q\n",b[:n])
		if err==io.EOF {
			break
		}

	}
}
