package main

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
	"unicode"
	"unicode/utf8"
)

// ============================================================================
// Example 1: charCount.go
// ============================================================================

var once sync.Once

func example1_charcount() {
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

// ============================================================================
// Example 2: iterateIn2Ways.go
// ============================================================================

func example2_iteratein2ways() {
	s := "vijay"
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i], " type: ", reflect.TypeOf(s[i]))
	}
	fmt.Println("Iterate using range")
	for _, v := range s {
		fmt.Println(v, " type: ", reflect.TypeOf(v))
	}
}

// ============================================================================
// Example 3: quotesBehaviour.go
// ============================================================================

func example3_quotesbehaviour() {
	s := `Go\tJava\nPython`
	fmt.Println(s)
	s = "Go\tJava\nPython"
	fmt.Println(s)
}

// ============================================================================
// Example 4: runeAndChar.go
// ============================================================================

func example4_runeandchar() {
	data := "We♥Go"
	fmt.Println(len(data)) // 7
	// If you count the number of characters is "We♥Go", it would be 5. So why 7?
	//In Go Strings are UTF-8 encoded, this means each character called rune can be of 1 to 4 bytes long.
	//Here,the character ♥ is taking 3 bytes, hence the total length of string is 7.
	data = "We♥Go"
	fmt.Println("rune count in string: ", utf8.RuneCountInString(data))
}

// ============================================================================
// Example 5: runeTest.go
// ============================================================================

func example5_runetest() {
	str := "Hello♧ hi♛♧!!"
	for _, c := range str {
		fmt.Printf("%c\n", unicode.ToUpper(c))
	}
}

// ============================================================================
// Example 6: trim.go
// ============================================================================

func example6_trim() {
	col := "F15_StopLight"
	v := strings.TrimSuffix(col, "_StopLight")
	fmt.Println(v)
}

// ============================================================================
// Main Function - Run All Examples
// ============================================================================

func main() {
	fmt.Println("Running 6 examples from string")
	fmt.Println(strings.Repeat("=", 60))

	fmt.Println("\n--- example1_charcount ---")
	example1_charcount()

	fmt.Println("\n--- example2_iteratein2ways ---")
	example2_iteratein2ways()

	fmt.Println("\n--- example3_quotesbehaviour ---")
	example3_quotesbehaviour()

	fmt.Println("\n--- example4_runeandchar ---")
	example4_runeandchar()

	fmt.Println("\n--- example5_runetest ---")
	example5_runetest()

	fmt.Println("\n--- example6_trim ---")
	example6_trim()

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("All examples completed!")
}
