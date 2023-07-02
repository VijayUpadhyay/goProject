package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "$VA_REVENUE_TOTAL_5YR_CAGR$"
	fmt.Print(strings.Trim(str, "$"))
}
