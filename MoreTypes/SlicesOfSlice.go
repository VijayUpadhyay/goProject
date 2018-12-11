package main

import (
	"fmt"
	"strings"
)

func main() {
	arrOfArrays:= [][]string{[]string{"_","_","_"},
							  []string{"_","_","_"},
	}
	arrOfArrays[0][0]="vijay"
	arrOfArrays[0][1]="vijay"
	arrOfArrays[0][2]="vijay"
	arrOfArrays[1][0]="vijay"

	fmt.Println(arrOfArrays)
	fmt.Println("_____________________")
	for i:=0;i<len(arrOfArrays);i++ {
		fmt.Printf("%s\n", strings.Join(arrOfArrays[i], "add"))
	}
}
