package main

import (
	"fmt"
)

const MAX = 100

func main() {
	work := make(chan int, MAX)
	result := make(chan int)

	// 1. Create channel of multiples of 3 and 5
	// concurrently using goroutine
	go func() {
		defer close(work)
		for i := 0; i < MAX; i++ {
			if (i%3) == 0 || (i%5) == 0 {
				work <- i
			}
		}
	}()

	// 2. Concurrently sum up work and put result
	// in channel result
	go func() {
		r := 0
		for i := range work {
			r = r + i
		}
		result <- r
	}()

	// 3. Wait for result, then print
	fmt.Println("Total : ", <-result)
}
