// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
)

func main() {
	route()
}

func route() {
	var i int64
	var wg sync.WaitGroup
	wg.Add(5)
	for i = 0; i < 5; i++ {
		go func(wg *sync.WaitGroup, j int64) {
			defer wg.Done()
			fmt.Printf("\n Value from counter: %d", j)
		}(&wg, i)
	}
	wg.Wait()
}
