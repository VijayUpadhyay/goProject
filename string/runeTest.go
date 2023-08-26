package main

import (
	"fmt"
	"unicode"
)

func main() {
	str := "Hello♧ hi♛♧!!"
	for _, c := range str {
		fmt.Printf("%c\n", unicode.ToUpper(c))
	}
}
