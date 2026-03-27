package main

import (
	"fmt"
	"strings"
)

// ============================================================================
// Example 1: AppendInSlice.go
// ============================================================================

func example1_appendinslice() {
	var arr []int
	s := arr[0:]
	r := []int{7, 8, 9}
	s = append(s, 0, 1, 2, 3, 4)
	s = append(s, r...)
	fmt.Println("s is:", s)
	fmt.Println(arr)
}

// ============================================================================
// Example 2: ArrayTest.go
// ============================================================================

func example2_arraytest() {
	var str [2]string
	str[0] = "vijay"
	str[1] = "Upadhyay"
	fmt.Println(str)

	fmt.Println("---------------------------")
	var arr = []int{2, 3}
	fmt.Println(arr)

	arr1 := []int{2, 3}
	fmt.Println(arr1)
}

// ============================================================================
// Example 3: MakeForSlice.go
// ============================================================================

func example3_makeforslice() {
	a := make([]int, 5)
	a[0] = 5
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println("------------------------")
	b := make([]int, 2, 5) // make(type,len,capability)
	b[0] = 10
	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println("------------------------")
	c := a[1:3]
	fmt.Println(c)
	fmt.Println(len(c))
}

// ============================================================================
// Example 4: mapEleSize.go
// ============================================================================

func example4_mapelesize() {
	arr := []int{10, 20, 20, 10, 10, 20, 5, 20}
	m := make(map[int]int)
	for _, val := range arr {
		if count, ok := m[val]; ok {
			m[val] = count + 1
		} else {
			m[val] = 1
		}
	}

	for k, v := range m {
		fmt.Printf("Elements: %d, Counts: %d \n", k, v)
	}
}

// ============================================================================
// Example 5: MapLiterals.go
// ============================================================================

type s struct {
	a int
	b float64
}

var ma = map[string]s{
	"key1": {10, 10.89},
	"key2": {20, 30.90},
}

func example5_mapliterals() {
	fmt.Println(ma)

	ele, ok := ma["key1"]
	fmt.Println("ele:", ele, "ok:", ok)
	ele1, ok1 := ma["key3"]
	fmt.Println("ele1:", ele1, "ok1:", ok1)
}

// ============================================================================
// Example 6: MapTest.go
// ============================================================================

type VarName struct {
	a int
	b int
}

type VarName2 struct {
	a1 string
	b1 string
}

var m map[string]VarName
var p2 map[int64]VarName2

func example6_maptest() {
	m = make(map[string]VarName)
	m["key"] = VarName{10, 20}

	p2 = make(map[int64]VarName2)
	p2[10] = VarName2{"p", "q"}
	p2[20] = VarName2{"dfvdf", "dvd"}
	fmt.Println(m["key"])
	fmt.Println(m)

	fmt.Println("=======================")
	fmt.Println(p2)
	fmt.Println(p2[10])
	fmt.Println(p2[30])
}

// ============================================================================
// Example 7: RangeOnSlice.go
// ============================================================================

func example7_rangeonslice() {
	a := []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, p := range a {
		fmt.Printf("2**%d = %d\n", i, p)
	}
	fmt.Println("_________")
	for i := range a {
		fmt.Printf("%d\n", i)
	}
}

// ============================================================================
// Example 8: SliceLengthAndCapacity.go
// ============================================================================

func example8_slicelengthandcapacity() {
	var s = []int{2, 3, 4, 56, 6, 7, 89, 9}
	fmt.Println(s)
	s = s[:0]
	fmt.Println(s)
	s = s[2:4]
	fmt.Println(s)
	s = s[0:]
	fmt.Println("without ", s)
	s = s[0:3]
	fmt.Println(s)

	s = s[0:5]
	fmt.Println(s)

	var arr []int
	fmt.Println(arr)
	if arr == nil {
		fmt.Println("Inside if for nil")
	}
}

// ============================================================================
// Example 9: SlicesOfSlice.go
// ============================================================================

func example9_slicesofslice() {
	arrOfArrays := [][]string{{"_", "_", "_"},
		{"_", "_", "_"},
	}

	arrOfArrays[0][0] = "vijay"
	arrOfArrays[0][1] = "vijay"
	arrOfArrays[0][2] = "vijay"
	arrOfArrays[1][0] = "vijay"
	fmt.Println(arrOfArrays)
	fmt.Println("_____________________")
	for i := 0; i < len(arrOfArrays); i++ {
		fmt.Printf("%s\n", strings.Join(arrOfArrays[i], "add"))
	}
}

// ============================================================================
// Example 10: SlicesTest.go
// ============================================================================

func example10_slicestest() {
	var arr = [6]int{5, 6, 7, 8, 9}
	var sliceEx = arr[1:4] // length is (4-1)
	fmt.Println("arr:", arr)
	fmt.Println("Slice:", sliceEx)
	fmt.Println("Slice length", len(sliceEx))
	p := arr[:2] // length is (2-0)
	fmt.Println("p:", p)
	q := arr[3:]
	fmt.Println("q:", q)

	q[0] = 100
	fmt.Println("--------------------------")
	fmt.Println("Updated arr:", arr)
	fmt.Println("Upadated sliceEx:", sliceEx)
	fmt.Println("Updated p:", p)
	fmt.Println("-------------Array of struct and implementing slice over that----------------")
	var r = []struct {
		a int
		b string
	}{
		{10, "ram"},
		{20, "sita"},
	}
	fmt.Println("r is:", r)
	fmt.Println("r is:", r[0])
}

// ============================================================================
// Example 11: StructTest.go
// ============================================================================

func example11_structtest() {
	type VariableName struct {
		a   int
		str string
	}
	fmt.Println(VariableName{5, "madhu"})

	//z:= make([]struct{},5)
	v := VariableName{10, "Rakesh"}
	var d = VariableName{}
	fmt.Println(d)
	fmt.Println("Roll Number:", v.a, "Name:", v.str)

	fmt.Println("Changing data in sruct using pointer")
	p := &v
	fmt.Println("Value in p:", p)
	p.a = 900
	fmt.Println("Changed struct:", v)

	//Array of structures/union
	/*for i:=0;i<5 ;i++  {
		//z[i] :=
		fmt.Println(z[i])
	}*/
}

// ============================================================================
// Example 12: Test2.go
// ============================================================================

func example12_test2() {
	fmt.Println("vijay upadhyay")

	a := make([]int, 1)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Println("slice")

	s, t := slice(5, 9)
	fmt.Println(s)
	fmt.Println(t)
	//fmt.Println(slice(5,8))
}

func slice(x, y int) (p, q int) {
	s := x - y
	t := x + y
	return s, t
}

// ============================================================================
// Main Function - Run All Examples
// ============================================================================

func main() {
	fmt.Println("Running 12 examples from datastructures")
	fmt.Println(strings.Repeat("=", 60))

	fmt.Println("\n--- example1_appendinslice ---")
	example1_appendinslice()

	fmt.Println("\n--- example2_arraytest ---")
	example2_arraytest()

	fmt.Println("\n--- example3_makeforslice ---")
	example3_makeforslice()

	fmt.Println("\n--- example4_mapelesize ---")
	example4_mapelesize()

	fmt.Println("\n--- example5_mapliterals ---")
	example5_mapliterals()

	fmt.Println("\n--- example6_maptest ---")
	example6_maptest()

	fmt.Println("\n--- example7_rangeonslice ---")
	example7_rangeonslice()

	fmt.Println("\n--- example8_slicelengthandcapacity ---")
	example8_slicelengthandcapacity()

	fmt.Println("\n--- example9_slicesofslice ---")
	example9_slicesofslice()

	fmt.Println("\n--- example10_slicestest ---")
	example10_slicestest()

	fmt.Println("\n--- example11_structtest ---")
	example11_structtest()

	fmt.Println("\n--- example12_test2 ---")
	example12_test2()

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("All examples completed!")
}
