package main

import (
<<<<<<<< HEAD:core/slices/TestRandVal.go
	"strconv"
	"math/rand"
	"fmt"
========
	"fmt"
	"math/rand"
	"strconv"
>>>>>>>> merge_conflicts:math/TestRandVal.go
)

func main() {
	var NumOfObjectToAddOrDelete = 6
	var existingNum = []string{"31", "1", "47", "18", "9", "37"}
<<<<<<<< HEAD:core/slices/TestRandVal.go
	var intArray = []int{}
========
	var intArray []int
>>>>>>>> merge_conflicts:math/TestRandVal.go
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
<<<<<<<< HEAD:core/slices/TestRandVal.go
		randVal := rand.Intn(50 - 0) + 0
		if isNumbExists(randVal, intArray) {
			i--;
========
		randVal := rand.Intn(50-0) + 0
		if isNumbExists(randVal, intArray) {
			i--
>>>>>>>> merge_conflicts:math/TestRandVal.go
			continue
		} else {
			test[i] = randVal
		}

	}
	fmt.Println("test array is: ", test)
}

<<<<<<<< HEAD:core/slices/TestRandVal.go
func isNumbExists(val int, intArr []int) (bool) {
========
func isNumbExists(val int, intArr []int) bool {
>>>>>>>> merge_conflicts:math/TestRandVal.go
	var flag = false
	for i := 0; i < len(intArr); i++ {
		if val == intArr[i] {
			flag = true
			break
		}
	}
	return flag
<<<<<<<< HEAD:core/slices/TestRandVal.go
}
========
}
>>>>>>>> merge_conflicts:math/TestRandVal.go
