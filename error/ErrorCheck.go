package main

import (
	"errors"
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, error) {

	if x < 0 {
		return -1, errors.New("Given number is negative")
	} else {
		return math.Sqrt(x), nil
	}

}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
