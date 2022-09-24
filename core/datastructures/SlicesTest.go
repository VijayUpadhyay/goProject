package main

import "fmt"

func main() {
	var arr = [6]int{5,6,7,8,9}
	var sliceEx  =arr[1:4] // length is (4-1)
	fmt.Println("arr:",arr)
	fmt.Println("Slice:",sliceEx)
	fmt.Println("Slice length",len(sliceEx))
	p:= arr[:2]  // length is (2-0)
	fmt.Println("p:",p)
	q:= arr[3:]
	fmt.Println("q:",q)

	q[0]=100
	fmt.Println("--------------------------")
	fmt.Println("Updated arr:",arr)
	fmt.Println("Upadated sliceEx:",sliceEx)
	fmt.Println("Updated p:",p)
	fmt.Println("-------------Array of struct and implementing slice over that----------------")
	var r = []struct{a int
				b string}{
					{10,"ram"},
					{20,"sita"},
	}
	fmt.Println("r is:",r)
	fmt.Println("r is:",r[0])
}
