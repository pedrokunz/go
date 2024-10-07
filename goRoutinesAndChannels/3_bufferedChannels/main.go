package main

import "fmt"

// main
// Objective: Understand buffered channels and their behavior.
// Task: Write a program that uses a buffered channel to send and receive 3 numbers.
func main() {
	// Create a buffered channel
	// This channel can hold up to 3 values
	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3

	// Close the channel
	close(ch)

	for number := range ch {
		fmt.Println(number)
	}
}
