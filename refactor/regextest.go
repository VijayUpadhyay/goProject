package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	//str1 := "Hello X42 I'm a Y-32.35 string Z30"
	str1 := "GrossMargin"
	fmt.Println("getSectionName(str1): ", getSectionName(str1))
	fmt.Println("name: ", getName("model.BalanceSheet"))
	getNameTrimR()
}

func getSectionName(str1 string) string {
	re := regexp.MustCompile(`[A-Z][^A-Z]*`)

	fmt.Printf("Pattern: %v\n", re.String()) // Print Pattern

	submatchall := re.FindAllString(str1, -1)
	/*	for _, element := range submatchall {
		fmt.Println(element)
	}*/
	return strings.Join(submatchall, " ")
}

func getName(s string) string {
	return strings.TrimPrefix(s, "model.")
}

func getNameTrimR() {
	fmt.Println(strings.TrimSuffix("invalid section: model.Test", "model."))
}
