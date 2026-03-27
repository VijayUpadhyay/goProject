package main

import (
	"fmt"
	"strings"
	"time"
)

// ============================================================================
// Example 1: BufferedChannelsEx.go
// ============================================================================

func example1_bufferedchannelsex() {

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	x, y := <-ch, <-ch
	/*fmt.Println(ch)
	fmt.Println(ch)*/

	/*fmt.Println(<-ch)
	fmt.Println(<-ch)*/

	fmt.Printf("X is %d and Y is %d", x, y)
}

// ============================================================================
// Example 2: ChannelDirections.go
// ============================================================================

func example2_channeldirections() {
	fmt.Println("Creating channels")
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pongs, pings)

	fmt.Println(<-pongs)
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}
func pong(pongs chan<- string, pings <-chan string) {
	msg := <-pings
	pongs <- msg
}

// ============================================================================
// Example 3: ChannelSynchronization.go
// ============================================================================

func example3_channelsynchronization() {
	fmt.Println("Starting sync")
	done := make(chan bool)
	go worker(done)
	<-done

}

func worker(c chan bool) {
	fmt.Println("Inside method....")
	time.Sleep(time.Millisecond)
	fmt.Println("done")

	c <- true
}

// ============================================================================
// Example 4: concat.go
// ============================================================================

func example4_concat() {
	fmt.Println("Hi")

}

// ============================================================================
// Example 5: receiveOnlyChannel.go
// ============================================================================

func example5_receiveonlychannel() {
	// Create a channel with a capacity of 3
	ch := make(chan int, 3)
	// Send the value 2 to the channel
	ch <- 2
	// Call the process function with the channel as an argument
	process(ch)
	// Print a new line
	fmt.Println()
}

func process(ch <-chan int) {
	// Receive a value from the channel and assign it to the variable s
	s := <-ch
	// Print the value of s
	fmt.Println(s)
	// Comment out the line below to prevent sending it to a read-only channel
	//ch <- 2 // ===========================================================================================> NOT ALLOWED
}

// ============================================================================
// Example 6: SelectOverChannel.go
// ============================================================================

func example6_selectoverchannel() {
	fmt.Println("Starting go func")

	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		time.Sleep(time.Millisecond * 3)
		c1 <- "Hi"
	}()

	go func() {
		time.Sleep(time.Millisecond * 2)
		c2 <- "Vijay"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)

		}
	}

}

// ============================================================================
// Example 7: SelectStatement.go
// ============================================================================

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func example7_selectstatement() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

// ============================================================================
// Example 8: SelectUsingChannels.go
// ============================================================================

func example8_selectusingchannels() {
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)
	//c3 <- "test"
	fmt.Println(c3)
	go func() {
		c1 <- "inside channel 1"
	}()

	go func() {
		c2 <- "inside channel 2"
		c3 <- "test"
	}()
	fmt.Println(c3)
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		case msg3 := <-c3:
			fmt.Println(msg3)

			/*default:
			fmt.Println("No data in channels")*/
		}
	}
	fmt.Println("End")
}

// ============================================================================
// Example 9: SelectWithIf.go
// ============================================================================

func example9_selectwithif() {
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}
}

// ============================================================================
// Example 10: SumValueInChannel.go
// ============================================================================

func example10_sumvalueinchannel() {
	s := []int{2, 23, 0, -24, 3, -2}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c //first <-c will have sum of the last three index values i.e recent value
	fmt.Println(x, y, x+y)
}

func sum(arr []int, c chan int) {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	c <- sum
}

// ============================================================================
// Example 11: unidirectionalchannel.go
// ============================================================================

func example11_unidirectionalchannel() {

}

// ============================================================================
// Main Function - Run All Examples
// ============================================================================

func main() {
	fmt.Println("Running 11 examples from channel")
	fmt.Println(strings.Repeat("=", 60))

	fmt.Println("\n--- example1_bufferedchannelsex ---")
	example1_bufferedchannelsex()

	fmt.Println("\n--- example2_channeldirections ---")
	example2_channeldirections()

	fmt.Println("\n--- example3_channelsynchronization ---")
	example3_channelsynchronization()

	fmt.Println("\n--- example4_concat ---")
	example4_concat()

	fmt.Println("\n--- example5_receiveonlychannel ---")
	example5_receiveonlychannel()

	fmt.Println("\n--- example6_selectoverchannel ---")
	example6_selectoverchannel()

	fmt.Println("\n--- example7_selectstatement ---")
	example7_selectstatement()

	fmt.Println("\n--- example8_selectusingchannels ---")
	example8_selectusingchannels()

	fmt.Println("\n--- example9_selectwithif ---")
	example9_selectwithif()

	fmt.Println("\n--- example10_sumvalueinchannel ---")
	example10_sumvalueinchannel()

	fmt.Println("\n--- example11_unidirectionalchannel ---")
	example11_unidirectionalchannel()

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("All examples completed!")
}
