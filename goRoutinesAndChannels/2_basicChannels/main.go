package main

import "fmt"

// main
// Objective: Learn how to use channels to communicate between goroutines.
// Task: Write a program that launches a goroutine to send numbers from 1 to 5 to a channel, and the main function should receive and print these numbers.
func main() {
	// Create a channel
	// This channel will be used to send integers
	// from the sendNumbers function to the main function
	// The channel is unbuffered, which means it can only
	// hold one value at a time
	// We initialize the channel with the make function because it is a reference type
	ch := make(chan int)

	// This is a non-blocking operation
	// It will run in a separate goroutine
	// and will send values to the channel
	go sendNumbers(ch)

	// This is a blocking operation
	// It will wait until the channel is closed
	// and all values are received
	for number := range ch {
		fmt.Printf("Received number: %d\n", number)
	}
}

// sendNumbers sends integers to a channel
func sendNumbers(ch chan int) {
	// Send values from 1 to 5 to the channel
	for i := 1; i <= 5; i++ {
		// <- is the channel operator
		// It is used to send and receive values from channels
		// In this case, we are sending the value of i to the channel
		ch <- i
	}

	// Close the channel
	// This is important to avoid deadlocks
	// and to signal that no more values will be sent
	close(ch)
}
