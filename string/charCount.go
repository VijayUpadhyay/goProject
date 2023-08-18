package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello♧ hi♛"
	//m := make(map[string]int)

	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
		fmt.Printf("%c \n", str[i])
	}

	// str = strings.ToUpper(str)
	fmt.Println("*******************RUNE*************************")
	for _, c := range str {
		fmt.Println(c)
		fmt.Printf("%c \n", c)
	}
	fmt.Println("*******************Char Count*************************")
	/* charToSearch := "h"
	fmt.Printf("Count of h is: %d", m[strings.ToUpper(charToSearch)]) */
	charCount(str, 'h')
}

func charCount(str string, char int32) {
	str = strings.ToUpper(str)
	count := 0
	for _, r := range str {
		if runeVal == r {
			count++
		}
	}
	fmt.Printf("Size of character:%c is: %d", runeVal, count)
}

func reverse(str string) string {
	return ""
}
