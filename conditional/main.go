package main

import (
	"fmt"
	"math"
	"runtime"
	"strings"
	"sync"
)

// ============================================================================
// Example 1: ExpressionSwitch.go
// ============================================================================

func example1_expressionswitch() {
	var grade = ""
	var marks = 80
	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 70:
		grade = "C"
	case 50, 60:
		grade = "D"
	default:
		grade = "E"
	}
	fmt.Println("Going for grade check with mark and values as: ", marks, grade)
	switch {
	case grade == "A":
		fmt.Println("Excellent")
	case grade == "B":
		fmt.Println("Good")
	case grade == "C":
		fmt.Println("Passed")
	case grade == "D":
		fmt.Println("Improve")
	case grade == "E":
		fmt.Println("Failed")

	}
	fmt.Println(marks, grade)
}

// ============================================================================
// Example 2: ForLoopTest.go
// ============================================================================

func example2_forlooptest() {
	var sum = 0
	fmt.Println(sum)
	for sum < 1000 {
		sum += 100
		fmt.Println(sum)
	}
}

// ============================================================================
// Example 3: IfStatement.go
// ============================================================================

func example3_ifstatement() {
	var sum int = 100

	if p := math.Pow10(sum); p < 1000 {
		fmt.Println("p is less than 1000")
	} else {
		fmt.Println("p is equalTo or greater than 1000")
	}

	fmt.Println()

	os := runtime.GOOS
	fmt.Println("os is: ", os)
}

// ============================================================================
// Example 4: syncConditional.go
// ============================================================================

var sharedRsc = make(map[string]interface{})

func example4_syncconditional() {
	var wg sync.WaitGroup
	wg.Add(2)
	m := sync.Mutex{}
	c := sync.NewCond(&m)
	go func() {
		// this go routine wait for changes to the sharedRsc
		c.L.Lock()
		for len(sharedRsc) == 0 {
			c.Wait()
		}
		fmt.Println("goroutine1", sharedRsc["rsc1"])
		c.L.Unlock()
		wg.Done()
	}()

	go func() {
		// this go routine wait for changes to the sharedRsc
		c.L.Lock()
		for len(sharedRsc) == 0 {
			c.Wait()
		}
		fmt.Println("goroutine2", sharedRsc["rsc2"])
		c.L.Unlock()
		wg.Done()
	}()

	// this one writes changes to sharedRsc
	c.L.Lock()
	sharedRsc["rsc1"] = "foo"
	sharedRsc["rsc2"] = "bar"
	c.Broadcast()
	c.L.Unlock()
	wg.Wait()
}

// ============================================================================
// Main Function - Run All Examples
// ============================================================================

func main() {
	fmt.Println("Running 4 examples from conditional")
	fmt.Println(strings.Repeat("=", 60))

	fmt.Println("\n--- example1_expressionswitch ---")
	example1_expressionswitch()

	fmt.Println("\n--- example2_forlooptest ---")
	example2_forlooptest()

	fmt.Println("\n--- example3_ifstatement ---")
	example3_ifstatement()

	fmt.Println("\n--- example4_syncconditional ---")
	example4_syncconditional()

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("All examples completed!")
}
