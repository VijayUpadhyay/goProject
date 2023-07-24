package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(strings.NewReader(`one
two
three
four
`))
	var (
		text strings.Builder
		n    int
	)
	for scanner.Scan() {
		n++
		text.WriteString(fmt.Sprintf("%d. %s\n", n, scanner.Text()))
	}
	fmt.Print(text.String())
	//os.Stdout.Write(text.String())
	// Output:
	// 1. One
	// 2. Two
	// 3. Three
	// 4. Four
}
