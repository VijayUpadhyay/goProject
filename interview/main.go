package main

import (
	"fmt"
	"strings"
	"sync"
)

// ============================================================================
// Example 1: compositionTest.go
// ============================================================================

type Staff struct {
	Name string
	Id   int32
}

type Student struct {
	Staff
	Subjects []string
}

type Faculty struct {
	Staff
	Subjects []string
}

func example1_compositiontest() {

	s := Student{Staff{Name: "Ram", Id: 100}, []string{"Maths", "Science"}}
	fmt.Println(s)

}

// ============================================================================
// Example 2: concurencyTest.go
// ============================================================================

var wg sync.WaitGroup

func example2_concurencytest() {
	ch := make(chan int)
	wg.Add(2)
	go f1(ch)
	go f2(ch)

	ch <- 100
	ch <- 200
	wg.Wait()
	close(ch)
}

func f1(ch chan int) {
	fmt.Println(<-ch)
	wg.Done()
}

func f2(ch chan int) {
	fmt.Println(<-ch)
	wg.Done()
}

// ============================================================================
// Main Function - Run All Examples
// ============================================================================

func main() {
	fmt.Println("Running 2 examples from interview")
	fmt.Println(strings.Repeat("=", 60))

	fmt.Println("\n--- example1_compositiontest ---")
	example1_compositiontest()

	fmt.Println("\n--- example2_concurencytest ---")
	example2_concurencytest()

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("All examples completed!")
}
