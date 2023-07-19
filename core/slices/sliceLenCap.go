<<<<<<<< HEAD:core/slices/sliceLenCap.go
package main

import (
	"fmt"
	"sort"
)

func main() {
	var s []int
	s = []int{2, 3, 5, 7, 8, 4, 5}
	fmt.Printf("Array is: %v, len is: %d and cap is: %d\n", s, len(s), cap(s))
	s1 := s[2:5]
	fmt.Printf("Array is: %v, len is: %d and cap is: %d\n", s1, len(s1), cap(s1))
	s1[2] = 100
	fmt.Println("After update................")
	fmt.Printf("Array is: %v, len is: %d and cap is: %d\n", s, len(s), cap(s))
	s1 = append(s1, 200)
	fmt.Printf("Array is: %v, len is: %d and cap is: %d\n", s1, len(s1), cap(s1))

	// string sorting using ASCII value
	str := []string{"ram", "al", "xx", "zz", "AA", "PP", "pp"}
	sort.Strings(str)
	fmt.Println(str)
}
========
package main

import (
	"fmt"
	"sort"
)

func main() {
	var s []int
	s = []int{2, 3, 5, 7, 8, 4, 5}
	fmt.Printf("Array is: %v, len is: %d and cap is: %d\n", s, len(s), cap(s))
	s1 := s[2:5]
	fmt.Printf("Array is: %v, len is: %d and cap is: %d\n", s1, len(s1), cap(s1))
	s1[2] = 100
	fmt.Println("After update................")
	fmt.Printf("Array is: %v, len is: %d and cap is: %d\n", s, len(s), cap(s))
	s1 = append(s1, 200)
	fmt.Printf("Array is: %v, len is: %d and cap is: %d\n", s1, len(s1), cap(s1))
	str := []string{"ram", "al", "xx", "zz", "AA", "PP", "pp"}
	sort.Strings(str)
	fmt.Println(str)
}
>>>>>>>> merge_conflicts:main/sliceofslice.go
