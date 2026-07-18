package main

import (
	"fmt"
	"sync"
)

// WaitGroup is used to wait for multiple goroutines to finish.
// Add(n)   -> Increases the goroutine counter by n.
// Done()   -> Decreases the counter by 1.
// Wait()   -> Blocks the main goroutine until the counter becomes 0.
var wg = sync.WaitGroup{}

func main() {

	// Create an unbuffered channel that can send and receive integers.
	// By default, channels are bidirectional.
	num := make(chan int)

	// We are about to create two goroutines.
	// Increase the WaitGroup counter to 2.
	wg.Add(2)

	// ---------------------- Sender Goroutine ----------------------
	go func(num chan<- int) {
		// chan<- means this goroutine can ONLY SEND values
		// through this channel. Reading is not allowed.

		// Send 42 to the channel.
		// Since the channel is unbuffered, this blocks until
		// another goroutine receives the value.
		num <- 42

		// Send another value.
		num <- 20

		// Close the channel.
		// Closing tells receivers:
		// "No more values will be sent."
		close(num)

		// Notify the WaitGroup that this goroutine has finished.
		wg.Done()

	}(num)

	// ---------------------- Receiver Goroutine ----------------------
	go func(num <-chan int) {

		// <-chan means this goroutine can ONLY RECEIVE values.
		// Sending to this channel would cause a compile-time error.

		// Range automatically:
		// 1. Receives every value from the channel.
		// 2. Stops when the channel is closed.
		for v := range num {
			fmt.Println(v)
		}

		// Notify the WaitGroup that this goroutine has finished.
		wg.Done()

	}(num)

	// Wait until both goroutines call Done().
	// This prevents main() from exiting early.
	wg.Wait()
}
