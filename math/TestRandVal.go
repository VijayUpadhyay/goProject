package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
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
