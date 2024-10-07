package main

import (
	"fmt"
	"sync"
)

// main
// Objective: Learn how to use mutexes to synchronize access to shared resources.
// Task: Write a program where multiple goroutines increment a shared counter using a mutex to ensure safe access.
func main() {
	// Create a counter variable
	// This variable will be shared between multiple goroutines
	var counter int

	// Create a mutex
	// This mutex will be used to synchronize access to the counter
	// Only one goroutine can hold the lock at a time
	var mu sync.Mutex

	// Create a WaitGroup
	// This WaitGroup will be used to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Create a function to increment the counter
	// This function will be called by multiple goroutines
	// It will lock the mutex, increment the counter, print the value, unlock the mutex, and notify the WaitGroup
	increment := func() {
		// Lock the mutex
		// This prevents other goroutines from accessing the shared resource
		mu.Lock()

		// Increment the counter
		counter++

		// Print the counter value
		fmt.Println("Counter value:", counter)

		// Unlock the mutex
		// This allows other goroutines to access the shared resource
		mu.Unlock()

		// Notify the WaitGroup that the goroutine is done
		wg.Done()
	}

	// Launch multiple goroutines to increment the counter
	for i := 0; i < 5; i++ {
		// Add 1 to the WaitGroup
		// This indicates that a new goroutine is starting
		wg.Add(1)

		// Launch a goroutine to increment the counter
		go increment()
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("Final counter:", counter)
}
