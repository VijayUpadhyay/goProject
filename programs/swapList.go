package main

import "fmt"

func main() {
	listObj := [3]int{1, 2, 3}
	listObj1 := []int{1, 2, 66, 3}
	swapContents(listObj)
	swapContents1(listObj1)
	fmt.Println(listObj)
	fmt.Println(listObj1)
}

func swapContents(listObj [3]int) {
	if len(listObj) == 0 {
		return
	}
	for i, j := 0, len(listObj)-1; i <= j; i, j = i+1, j-1 {
		listObj[i], listObj[j] = listObj[j], listObj[i]
	}
}

func swapContents1(listObj []int) {
	if len(listObj) == 0 {
		return
	}
	for i, j := 0, len(listObj)-1; i < j; i, j = i+1, j-1 {
		listObj[i], listObj[j] = listObj[j], listObj[i]
	}
}
