package main

import (
	"fmt"
	"sync"
)

// main
// Objective: Combine mutexes and channels to manage complex synchronization.
// Task: Write a program where multiple goroutines update a shared map concurrently. Use a mutex to protect the map and a channel to signal completion.
func main() {
	// Create a mutex
	// This mutex will be used to synchronize access to the map
	// Only one goroutine can hold the lock at a time
	var mu sync.Mutex

	// Create a WaitGroup
	// This WaitGroup will be used to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Create a map to store data
	// This map will be shared between multiple goroutines
	data := make(map[int]int)

	// Create a channel to signal completion
	// This channel will be used to notify the main goroutine that all updates are done
	done := make(chan bool)

	// Create a function to update the map
	// This function will be called by multiple goroutines
	// It will lock the mutex, update the map, unlock the mutex, and notify the WaitGroup
	updateMap := func(key, value int) {
		// Lock the mutex
		// This prevents other goroutines from accessing the shared resource
		mu.Lock()

		// Update the map
		data[key] = value

		// Print the updated map
		fmt.Printf("Updated map: %v\n", data)

		// Unlock the mutex
		mu.Unlock()
		// Notify the WaitGroup that the goroutine is done
		wg.Done()
	}

	// Launch multiple goroutines to update the map
	for i := 1; i <= 5; i++ {
		// Add 1 to the WaitGroup
		// This indicates that a new goroutine is starting
		wg.Add(1)
		// Launch a goroutine to update the map
		go updateMap(i, i*10)
	}

	// Launch a goroutine to wait for all updates to finish
	go func() {
		// Wait for all goroutines to finish
		wg.Wait()
		// Close the channel to signal completion
		done <- true
	}()

	// Wait for the channel to be closed
	// This is a blocking operation
	<-done

	// Print the final map
	fmt.Printf("Final map: %v\n", data)
}
