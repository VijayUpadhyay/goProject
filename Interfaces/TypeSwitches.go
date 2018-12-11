package main

import "fmt"

func main()  {
	var i interface{}
	typeCheck(i)
	i=10
	typeCheck(i)
	i = "vijay"
	typeCheck(i)
	i=79.00
	typeCheck(i)

}

func typeCheck(i interface{})  {
	switch val:=i.(type) {
	case int:
		fmt.Printf("Type is int for: %v\n",val)
	case string:
		fmt.Println("Type is string for",val)
	case float64:
		fmt.Printf("Type is %T for: %v\n",val,val)
	default:
		fmt.Printf("Type for value %v is : %T\n",val,val)
	}
}