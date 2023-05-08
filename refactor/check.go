package main

import (
	"fmt"
	"time"
)

func main() {
	val1 := "2023-01-24T17:04:46.145306+05:30"
	kv := make(map[string]interface{})
	kv["CompletedDate"] = val1

	val, perr := time.Parse(time.RFC3339, kv["CompletedDate"].(string))
	fmt.Println(val)
	fmt.Println(perr)
}
