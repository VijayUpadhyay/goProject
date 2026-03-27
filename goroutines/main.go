package main

import (
	"fmt"
	"strings"
	"time"
)

// ============================================================================
// Example 1: bufChannel.go
// ============================================================================

func example1_bufchannel() {
	ch := make(chan int, 2)
	go func(chan int) {
		for i := 0; i < 2; i++ {
			ch <- 10
			time.Sleep(3 * time.Second)
			ch <- 20
		}
	}(ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	//TODO: check why below 2 lines are working fine?
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	//fmt.Println(<-ch) // this will cause deadlock
}

// ============================================================================
// Example 2: ex1.go
// ============================================================================

func SumOfSquares1(c, quit chan int) {
	y := 1
	for {
		select {
		case c <- y * y:
			y++
		case <-quit:
			return
		}
	}
}

func example2_ex1() {
	mychannel := make(chan int)
	quitchannel := make(chan int)
	sum := 0
	go func() {
		for i := 1; i <= 5; i++ {
			sum += <-mychannel
		}
		fmt.Println(sum)
		quitchannel <- 0
	}()
	SumOfSquares1(mychannel, quitchannel)
}

// ============================================================================
// Example 3: simpleGoroutine.go
// ============================================================================

func sayHello(str string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("str at %d is: %s \n", i+1, str)
	}
}
func example3_simplegoroutine() {
	fmt.Println()
	go sayHello("hello")
	sayHello("hi")
}

// ============================================================================
// Example 4: summation.go
// ============================================================================

// ============================================================================
// Example 5: sumOfSquares.go
// ============================================================================

func SumOfSquares(c, quit chan int) {
	y := 1
	for {
		select {
		case val := <-c:
			c <- y * y
			y++
			fmt.Println(val)
		case <-quit:
			return
		}
	}
}

func example5_sumofsquares() {
	mychannel := make(chan int)
	quitchannel := make(chan int)
	sum := 0
	go func() {
		for i := 1; i <= 5; i++ {
			sum += <-mychannel
		}
		fmt.Println(sum)
		quitchannel <- 0
	}()
	SumOfSquares(mychannel, quitchannel)
}

// ============================================================================
// Example 6: unbufChannel.go
// ============================================================================

func example6_unbufchannel() {
	ch := make(chan int)
	go func(chan int) {
		ch <- 10
	}(ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch) // passed once, read once
	/*fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)*/
}

// ============================================================================
// Main Function - Run All Examples
// ============================================================================

func main() {
	fmt.Println("Running 5 examples from goroutines")
	fmt.Println(strings.Repeat("=", 60))

	fmt.Println("\n--- example1_bufchannel ---")
	example1_bufchannel()

	fmt.Println("\n--- example2_ex1 ---")
	example2_ex1()

	fmt.Println("\n--- example3_simplegoroutine ---")
	example3_simplegoroutine()

	fmt.Println("\n--- example5_sumofsquares ---")
	example5_sumofsquares()

	fmt.Println("\n--- example6_unbufchannel ---")
	example6_unbufchannel()

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("All examples completed!")
}
