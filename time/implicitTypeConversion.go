package main

import (
	"fmt"
	"time"
)

func main() {
	var timeout = 3
	fmt.Println(timeout)
	//	fmt.Println(timeout * time.Millisecond)
	//Millisecond's underlying type is an int64, which the compiler knows how to convert to.
	//Literals and constants are untyped until they are used, unless the type is explicitly declared.
	//timeout is an untyped constant in this example. Its type is implicitly converted to time.Millisecond
	const timeout1 = 10
	fmt.Println(timeout1)
	fmt.Println(timeout1 * time.Millisecond)
}
