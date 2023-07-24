package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	data := "We♥Go"
	fmt.Println(len(data)) // 7
	// If you count the number of characters is "We♥Go", it would be 5. So why 7?
	//In Go Strings are UTF-8 encoded, this means each character called rune can be of 1 to 4 bytes long.
	//Here,the character ♥ is taking 3 bytes, hence the total length of string is 7.
	data = "We♥Go"
	fmt.Println("rune count in string: ", utf8.RuneCountInString(data))
}
