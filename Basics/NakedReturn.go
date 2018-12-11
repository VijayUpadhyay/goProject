package Basics

import (
	"fmt"
)

func main() {
	fmt.Println(nakedReturn(100))
}

func nakedReturn(sum int) (x,y int){
	x=sum/10
	y=sum-10
	return
}