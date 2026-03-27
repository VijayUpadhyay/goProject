package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

// ============================================================================
// Example 1: paths.go
// ============================================================================

func example1_paths() {

}

// ============================================================================
// Example 2: sortByValFloat64.go
// ============================================================================

// round a float64 to 2 decimal places.
func round(n float64) float64 {
	return math.Round(n*100) / 100
}

type KV struct {
	Key   string
	Value float64
}

type KVList []KV

func (p KVList) Len() int           { return len(p) }
func (p KVList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p KVList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p KVList) Get() []KV          { return p }

func SortByValue(data map[string]float64) []KV {
	var i int
	sorted := make(KVList, len(data))
	for k, v := range data {
		sorted[i] = KV{k, v}
		i++
	}
	sort.Sort(sort.Reverse(sorted))
	return sorted.Get()
	/* for _, pair := range sorted {
		fmt.Printf("%s %v\n", pair.Key, round(pair.Value))
	} */
}

func example2_sortbyvalfloat64() {
	data := make(map[string]float64)
	data["red"] = 1.00
	data["green"] = 3.00
	data["blue"] = 2.00

	list := SortByValue(data)
	for k, v := range list {
		fmt.Println(k, v)
	}
}

// ============================================================================
// Main Function - Run All Examples
// ============================================================================

func main() {
	fmt.Println("Running 2 examples from udemy")
	fmt.Println(strings.Repeat("=", 60))

	fmt.Println("\n--- example1_paths ---")
	example1_paths()

	fmt.Println("\n--- example2_sortbyvalfloat64 ---")
	example2_sortbyvalfloat64()

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("All examples completed!")
}
