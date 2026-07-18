package main

// WARNING : THE LOCK AND UNLOCK METHOD IS ALWAYS WRITTEN INSIDE THE GOROUTINE FUNCTION. THIS CODE DOESN'T DEMONSTRATE THE BEST PRACTICES.

import (
	"fmt"
	"runtime"
	"sync"
)

var count = 0             // Shared counter accessed by multiple goroutines.
var wg = sync.WaitGroup{} // Add() increments the counter, Done() decrements it, and Wait() blocks until the counter reaches zero.
var rwm = sync.RWMutex{}  // It allows multiple readers (RLock) or one writer (Lock), but never both at the same time.

func main() {

	for range 10 {
		wg.Add(2) // Two goroutines will be started, so increment the WaitGroup counter by 2.

		rwm.RLock() // Acquire a read lock before reading the shared counter.
		go sayHello()

		rwm.Lock() // Acquire a write lock before modifying the shared counter.
		go counter()
	}
	wg.Wait()

}

func sayHello() {
	fmt.Printf("Hello i am counter number : %v \n", count)
	rwm.RUnlock() // Release the read lock.
	wg.Done()     // Signal that this goroutine has finished.
}

func counter() {
	count++
	rwm.Unlock() // Release the write lock.
	wg.Done()    // Signal that this goroutine has finished.
}

var Threads = runtime.GOMAXPROCS(-1) // Returns the number of CPU cores Go is currently allowed to use in parallel
