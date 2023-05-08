package main

import (
	"fmt"
	"math"
	"time"
)

const quarter int = 3

func main() {

	/*fmt.Println(int(t.Month()))
	fmt.Println(int(time.Now().UTC().Month()))*/
	fmt.Println(GetCurrentQuarterAndYear())

}

func GetCurrentQuarterAndYear() (int, int) {
	t := time.Date(2022, time.Month(10), 11, 12, 13, 44, 213, time.Local)
	return int(math.Ceil(float64((int(t.UTC().Month())-1)/quarter))) + 1, t.Year()
}
