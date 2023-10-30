package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
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
