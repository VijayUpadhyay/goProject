package main

import "fmt"

func main() {
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
