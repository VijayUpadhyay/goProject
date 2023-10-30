package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	sliceTest()
}

func sliceTest() {
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
