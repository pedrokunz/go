package main

import (
	"fmt"
	"time"
)

// main
// Objective: Implement a producer-consumer scenario using channels and goroutines.
// Task: Write a program where one goroutine produces numbers and sends them to a channel, and another goroutine consumes these numbers and prints them.
func main() {
	// Create a channel
	// This channel will be used to send integers
	ch := make(chan int)

	// These are non-blocking operations
	go producer(ch)
	go consumer(ch)

	// This is a blocking operation
	// Wait for the goroutines to finish
	time.Sleep(6 * time.Second)
}

// producer sends integers to a channel
// It sends numbers from 1 to 5 to the channel
// It sleeps for 1 second between each number
// It closes the channel after sending all numbers
func producer(ch chan int) {
	for i := 1; i <= 5; i++ {
		// Send the number to the channel
		ch <- i

		fmt.Printf("Produced number: %d\n", i)

		// Sleep for 1 second
		time.Sleep(1 * time.Second)
	}

	// Close the channel
	// This is important to avoid deadlocks
	close(ch)
}

// consumer receives integers from a channel
// It prints the numbers received from the channel
// It stops when the channel is closed
func consumer(ch chan int) {
	// Receive values from the channel
	// This is a blocking operation
	for number := range ch {
		fmt.Printf("Consumed number: %d\n", number)
	}
}
