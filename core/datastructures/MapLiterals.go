package main

import "fmt"

type s struct {
	a int
	b float64
}

var ma = map[string]s{
	"key1": s{10, 10.89},
	"key2": {20, 30.90},
}

func main() {
	fmt.Println(ma)

	ele, ok := ma["key1"]
	fmt.Println("ele:", ele, "ok:", ok)
	ele1, ok1 := ma["key3"]
	fmt.Println("ele1:", ele1, "ok1:", ok1)
}
