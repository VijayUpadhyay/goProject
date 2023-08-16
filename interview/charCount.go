package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello hi"
	m := make(map[string]int)
	str = strings.ToUpper(str)
	for _, c := range str {
		s := fmt.Sprintf("%s", c)
		count, ok := m[s]
		if ok {
			m[s] = count + 1
		} else {
			m[s] = 0
		}
	}

	charToSearch := "h"
	fmt.Printf("Count of h is: %d", m[strings.ToUpper(charToSearch)])

}

func reverse(str string) string {
if str =="" {
	return
}
var revStr string
//
revStr=  str[l]+reverse(str[:len(str)-1])
return revStr
}
