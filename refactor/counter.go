package main

import (
	"fmt"
	"sync/atomic"
)

var n int64

func main() {
	n = 100
	atomic.AddInt64(&n, 1)
	fmt.Println(n)
}
