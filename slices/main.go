package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

// ============================================================================
// Example 1: append.go
// ============================================================================

func example1_append() {
	var a [6]int = [6]int{0, 1, 2} // 0,1,2,0,0
	fmt.Println(len(a), cap(a), a)
	var t []int = []int{0, 1, 2} // 0,1,2,0,0
	fmt.Println(len(t), cap(t), t)
	fmt.Println("-------------------------------------")
	b := a[1:4]       //1,2,0
	b = append(b, 10) //b: 1,2,0,10, a: 0,1,2,0,0,10
	b = append(b, 10)
	fmt.Println("-------------------------------------")
	fmt.Println("a: ", len(a), cap(a), a)
	fmt.Println("b: ", len(b), cap(b), b)
	fmt.Println("-------------------------------------")
	b = append(b, 10)
	b = append(b, 10)
	fmt.Println("a: ", len(a), cap(a), a)
	fmt.Println("b: ", len(b), cap(b), b)
}

// ============================================================================
// Example 2: capCalc.go
// ============================================================================

func example2_capcalc() {
	sliceTest()
}

func sliceTest() {
	var a [6]int = [6]int{0, 1, 2} // 0,1,2,0,0
	b := a[1:3]                    // 1,2,0
	fmt.Println("Before add")
	fmt.Println(len(b), cap(b), b) // cap = len(a)-1 ==> 3 5 [1,2,0]
	b = append(b, 2000)
	add(b)
	fmt.Println("After add")
	fmt.Println(len(b), cap(b), b) // 8,8, //{9,10,100,20,30,0,0,0}
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	fmt.Printf("Slice Header(main): %+v\n", sh)
}

func add(s []int) {
	fmt.Println("Received add")
	fmt.Println(len(s), cap(s), s)
	s[1] = 10         // {1,10,0}
	s = append(s, 20) // {1,10,0,20}
	//s = append(s, 30) //{1,10,0,20,30}
	s[0] = 9   //{9,10,0,20,30}
	s[2] = 100 //{9,10,100,20,30}
	fmt.Println("Inside add -- end")
	fmt.Println(len(s), cap(s), s)
	s[3] = 5000
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	fmt.Printf("Slice Header: %+v\n", sh)
}

// ============================================================================
// Example 3: copyChk.go
// ============================================================================

func example3_copychk() {
	fmt.Println("test1")
	test1()
	fmt.Println("test2")
	test2()
}

func test1() {
	a := []int{1, 2}
	b := []int{3, 4}
	check := a
	copy(a, b) // the check variable is used to hold a reference to the original slice description to make sure it is really copied. So copied to their reference variables also
	fmt.Println(a, b, check)
	// Output: [3 4] [3 4] [3 4]
}

func test2() {
	a := []int{1, 2}
	b := []int{3, 4}
	check := a
	a = b // operation does not copy the slice contents, only the slice description
	fmt.Println(a, b, check)
	// Output: [3 4] [3 4] [1 2]
}

// ============================================================================
// Example 4: DeleteChk.go
// ============================================================================

func example4_deletechk() {
	var NumOfObjectToAddOrDelete = 3
	var intArray = []int{31, 3, 4, 66, 77}
	size := len(intArray) - NumOfObjectToAddOrDelete
	intArray = intArray[:size]
	fmt.Println(intArray)
}

// ============================================================================
// Example 5: sizeChk.go
// ============================================================================

func example5_sizechk() {
	sliceSizeCheck()
}

func sliceSizeCheck() {
	var a [5]int = [5]int{0, 1, 2} // 0,1,2,0,0
	fmt.Println(len(a), cap(a), a)
	b := a[1:4]       //1,2,0
	b = append(b, 10) //b: 1,2,0,10, a: 0,1,2,0,0,10
	fmt.Println("a: ", len(a), cap(a), a)
	fmt.Println("b: ", len(b), cap(b), b)
	sliceSizeCheckViaMethod(b)
	fmt.Println("After function call")
	fmt.Println(len(a), cap(a), a)
	fmt.Println(len(b), cap(b), b)
}

func sliceSizeCheckViaMethod(b []int) {
	b = append(b, 10)
	b = append(b, 10)
	b = append(b, 10)
}

// ============================================================================
// Example 6: sliceCheck.go
// ============================================================================

func example6_slicecheck() {
	//s1 := new([]int)
	// fmt.Printf("Slice len: %d and Cap: %d", len(*s1), cap(*s1))
	arr := [6]int{1, 2, 3, 4, 5, 6}
	//s1 := make([]int, 6)
	s1 := arr[2:5]
	fmt.Printf("Slice: %+v len: %d and Cap: %d, Address: %p \n", s1, len(s1), cap(s1), &s1)
	s1 = append(s1, 10)
	fmt.Println(s1, arr)
	s1[3] = 1000
	fmt.Println(s1, arr)
	s1 = append(s1, 100) // once capacity is full, new slice will be created and all elements will be copied into it. From here, no link between slice and array
	fmt.Println(s1, arr)
	fmt.Printf("Now, Slice: %+v len: %d and Cap: %d, Address: %p \n", s1, len(s1), cap(s1), &s1)
	arr[3] = 200
	s1[1] = 20000
	fmt.Println(s1, arr)
}

// ============================================================================
// Example 7: sliceChk.go
// ============================================================================

func example7_slicechk() {
	vals := make([]int, 2)
	fmt.Println("Capacity was:", cap(vals))
	fmt.Println("Length was:", len(vals))
	for i := 0; i < 5; i++ {
		vals = append(vals, i)
		fmt.Println("Capacity is now:", cap(vals))
		fmt.Println("Length is now:", len(vals))
	}
	fmt.Println(vals)
	vals = append(vals, 100)
	fmt.Println("1- Recent capacity: ", cap(vals))
	fmt.Println("1- Recent length: ", len(vals))
	vals = append(vals, 100)
	fmt.Println("2- Recent capacity: ", cap(vals))
	fmt.Println("2- Recent length: ", len(vals))

	fmt.Println("========================================")
	test()
}

func test() {
	text := "val: "
	//n := 10
	/* text += fmt.Sprintf("%d. %s\n", n, scanner.Text())
	text = append([]byte(text), fmt.Sprintf("%d. %s\n", n, scanner.Text())...) */
	text += "vijay"
	fmt.Println(text)
}

// ============================================================================
// Example 8: sliceLenCap.go
// ============================================================================

func example8_slicelencap() {
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
	fmt.Println(strings.Join(str, "::"))
}

// ============================================================================
// Example 9: slicePtr.go
// ============================================================================

func example9_sliceptr() {
	sliceTest2()
}

func sliceTest2() {
	var a [6]int = [6]int{0, 1, 2} // 0,1,2,0,0
	b := a[1:4]                    // 1,2,0
	fmt.Println("Before add")
	fmt.Println(len(b), cap(b), b) // cap = len(a)-1 ==> 3 5 [1,2,0]
	b = append(b, 2000)
	addPtr(&b)
	fmt.Println("After add")
	fmt.Println(len(b), cap(b), b) // 8,8, //{9,10,100,20,30,0,0,0}
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	fmt.Printf("Slice Header(main): %+v\n", sh)

}

func addPtr(s *[]int) {
	fmt.Println("Received addPtr--------------------")
	fmt.Println(len(*s), cap(*s), s)
	(*s)[1] = 10        // {1,10,0}
	*s = append(*s, 20) // {1,10,0,20}
	//s = append(s, 30) //{1,10,0,20,30}
	(*s)[0] = 9   //{9,10,0,20,30}
	(*s)[2] = 100 //{9,10,100,20,30}
	fmt.Println(len(*s), cap(*s), s)
	(*s)[3] = 5000
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	fmt.Printf("Slice Header: %+v\n", sh)
	fmt.Println("Inside addPtr ------------------- end")
}

// ============================================================================
// Example 10: sliceRef.go
// ============================================================================

func example10_sliceref() {
	originalArr := [9]int{1, 2, 3, 22, 33, 11, 9, 909, 6576} // mentioning/ not mentioning size makes it array else it is slice
	fmt.Printf("Type is %T \n", originalArr)
	var sliceArr []int
	sliceArr = originalArr[:4]
	fmt.Println("originalArr1: ", originalArr)
	fmt.Println("sliceArr1: ", sliceArr)
	sliceArr[2] = 100
	fmt.Println("originalArr2: ", originalArr)
	fmt.Println("sliceArr2: ", sliceArr)
	for p := 0; p < 3; p++ {
		sliceArr = append(sliceArr, p)
	}
	fmt.Println("originalArr3: ", originalArr)
	fmt.Println("sliceArr3: ", sliceArr)
	for p := 0; p < 5; p++ {
		sliceArr = append(sliceArr, p)
	}
	fmt.Println("originalArr4: ", originalArr)
	fmt.Println("sliceArr4: ", sliceArr)
	sliceArr[1] = 500
	fmt.Println("originalArr5: ", originalArr)
	fmt.Println("sliceArr5: ", sliceArr)
}

// ============================================================================
// Example 11: slickWorking.go
// ============================================================================

func example11_slickworking() {
	originalArr := []int{999, 2, 3, 22, 33, 11, 9, 909, 6576}
	var sliceArr []int
	sliceArr = originalArr[:4]
	fmt.Println("originalArr1: ", originalArr)
	fmt.Println("sliceArr1: ", sliceArr)
	sliceArr[2] = 100
	fmt.Println("originalArr2: ", originalArr)
	fmt.Println("sliceArr2: ", sliceArr)
	for p := 0; p < 3; p++ {
		sliceArr = append(sliceArr, p)
	}
	fmt.Println("originalArr3: ", originalArr)
	fmt.Println("sliceArr3: ", sliceArr)
	for p := 0; p < 5; p++ {
		sliceArr = append(sliceArr, p)
	}
	fmt.Println("originalArr4: ", originalArr)
	fmt.Println("sliceArr4: ", sliceArr)
	sliceArr[1] = 500
	fmt.Println("originalArr5: ", originalArr)
	fmt.Println("sliceArr5: ", sliceArr)
}

// ============================================================================
// Example 12: StructToJSON_1.go
// ============================================================================

type UCSUserAccess struct {
	Groups                 map[string]string  `json:"groups"`
	Access                 map[string]string  `json:"access"`
	IsExpert               string             `json:"isExpert"`
	SubscriptionGroupDates map[string]GroupID `json:"subscription_group_dates"`
}

type GroupID struct {
	StartDate int `json:"START_DATE"`
	EndDate   int `json:"END_DATE"`
}

func example12_structtojson_1() {
	var test = UCSUserAccess{
		Access:                 map[string]string{"STOCK": "1", "FOREX": "1", "WEBFOREX": "1", "WEBSTOCK": "1"},
		SubscriptionGroupDates: map[string]GroupID{"32": {1464753600, 1472616000}, "42": {1470024000, 1472616000}},
	}
	body, err := json.Marshal(test)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

// ============================================================================
// Example 13: StructToJSON_2.go
// ============================================================================

type StructData struct {
	name       string
	rollNo     int
	address    string
	fatherName string
	motherName string
}

func example13_structtojson_2() {
	var test = StructData{"Vijay", 4, "address unknown", "umesh", "indu"}
	body, err := json.Marshal(test)
	if err != nil {
		fmt.Println("Unable to marshall to JSON")
		return
	} else {
		fmt.Println(string(body))
	}

	//Unmarshall response
	res := StructData{}
	json.Unmarshal([]byte(string(body)), &res)
	fmt.Println(res)
}

// ============================================================================
// Example 14: TestRandVal.go
// ============================================================================

func example14_testrandval() {
	var NumOfObjectToAddOrDelete = 6
	var existingNum = []string{"31", "1", "47", "18", "9", "37"}
	var intArray []int
	for _, i := range existingNum {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		intArray = append(intArray, j)
	}
	test := make([]int, NumOfObjectToAddOrDelete)
	for i := 0; i < NumOfObjectToAddOrDelete; i++ {
		//rand.Seed(time.Now().Unix())
		randVal := rand.Intn(50-0) + 0
		if isNumbExists(randVal, intArray) {
			i--
			continue
		} else {
			test[i] = randVal
		}

	}
	fmt.Println("test array is: ", test)
}

func isNumbExists(val int, intArr []int) bool {
	var flag = false
	for i := 0; i < len(intArr); i++ {
		if val == intArr[i] {
			flag = true
			break
		}
	}
	return flag
}

// ============================================================================
// Example 15: UnaryOperators.go
// ============================================================================

func example15_unaryoperators() {
	fmt.Print(1 << 3)
}

// ============================================================================
// Example 16: UnmarshallTimestamp.go
// ============================================================================

//json unmarshal time that isn't in RFC 3339 format

type CustomTime struct {
	time.Time
}

const ctLayout = "2006/01/02|15:04:05"

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	ct.Time, err = time.Parse(ctLayout, string(b))
	return
}

func (ct *CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(ct.Time.Format(ctLayout)), nil
}

type Args struct {
	Time CustomTime
}

var data = `
	{"Time": "2014/08/01|11:27:18"}
`

func example16_unmarshalltimestamp() {
	a := Args{}
	fmt.Println(json.Unmarshal([]byte(data), &a))
	fmt.Println(a.Time.String())
}

// ============================================================================
// Main Function - Run All Examples
// ============================================================================

func main() {
	fmt.Println("Running 16 examples from slices")
	fmt.Println(strings.Repeat("=", 60))

	fmt.Println("\n--- example1_append ---")
	example1_append()

	fmt.Println("\n--- example2_capcalc ---")
	example2_capcalc()

	fmt.Println("\n--- example3_copychk ---")
	example3_copychk()

	fmt.Println("\n--- example4_deletechk ---")
	example4_deletechk()

	fmt.Println("\n--- example5_sizechk ---")
	example5_sizechk()

	fmt.Println("\n--- example6_slicecheck ---")
	example6_slicecheck()

	fmt.Println("\n--- example7_slicechk ---")
	example7_slicechk()

	fmt.Println("\n--- example8_slicelencap ---")
	example8_slicelencap()

	fmt.Println("\n--- example9_sliceptr ---")
	example9_sliceptr()

	fmt.Println("\n--- example10_sliceref ---")
	example10_sliceref()

	fmt.Println("\n--- example11_slickworking ---")
	example11_slickworking()

	fmt.Println("\n--- example12_structtojson_1 ---")
	example12_structtojson_1()

	fmt.Println("\n--- example13_structtojson_2 ---")
	example13_structtojson_2()

	fmt.Println("\n--- example14_testrandval ---")
	example14_testrandval()

	fmt.Println("\n--- example15_unaryoperators ---")
	example15_unaryoperators()

	fmt.Println("\n--- example16_unmarshalltimestamp ---")
	example16_unmarshalltimestamp()

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("All examples completed!")
}
