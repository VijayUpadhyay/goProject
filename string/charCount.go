package main

import (
	"fmt"
	"strings"
	"sync"
	"unicode"
)

var once sync.Once

func main() {
	str := "Hello♧ hi♛♧!!"
	charCount(str)
	str2 := reverse(str)
	fmt.Printf("String after reversal: %s \n", str2)
	fmt.Printf("Count of char ♧ is: %d \n", charCountByChar(str, '♧'))
	fmt.Printf("Count of char ♛ is: %d \n", charCountByChar(str, '♛'))
	fmt.Printf("Count of char h is: %d \n", charCountByChar(str, 'h'))
	fmt.Printf("Count of char(ignoreCase) h is: %d \n", charCountIgnoreCaseByChar(str, 'h'))
}

func charCount(str string) map[int32]int64 {
	m := make(map[int32]int64)
	for _, c := range str {
		if val, ok := m[c]; ok {
			m[c] = val + 1
		} else {
			m[c] = 1
		}
	}
	once.Do(func() {
		fmt.Printf("Map with char count: %+v \n", m)
		fmt.Println("---------- Char count by chars---------")
		for k, v := range m {
			fmt.Printf("%c : %d\n", k, v)
		}
	})
	return m
}

func charCountIgnoreCase(str string) map[int32]int64 {
	strUpper := strings.ToUpper(str)
	m2 := make(map[int32]int64)
	for _, c := range strUpper {
		if val, ok := m2[c]; ok {
			m2[c] = val + 1
		} else {
			m2[c] = 1
		}
	}
	return m2
}

func charCountIgnoreCaseByChar(str string, charVal rune) int64 {
	upperCharVal := unicode.ToUpper(charVal)
	mapData := charCountIgnoreCase(str)
	count, _ := mapData[upperCharVal]
	return count

}

func charCountByChar(str string, charVal rune) int64 {
	m := charCount(str)
	count, _ := m[charVal]
	return count
}

func reverse(str string) string {
	runeSlice := []rune(str)
	for i, j := 0, len(runeSlice)-1; i < j; i, j = i+1, j-1 {
		runeSlice[i], runeSlice[j] = runeSlice[j], runeSlice[i]
	}
	return string(runeSlice)
}
