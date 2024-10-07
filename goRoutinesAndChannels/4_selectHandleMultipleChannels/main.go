package main

import (
	"fmt"
	"time"
)

// main
// Objective: Learn how to use the select statement to handle multiple channels.
// Task: Write a program that uses two channels to send numbers and letters concurrently, and use select to print whichever is received first.
func main() {
	// Create two channels
	// One channel will be used to send integers
	// The other channel will be used to send strings
	numbersChannel := make(chan int)
	lettersChannel := make(chan string)

	// This is a non-blocking operation
	// It will run in a separate goroutine
	// and will send values to the channels
	go sendNumbers(numbersChannel)
	go sendLetters(lettersChannel)

	// Use the select statement to receive values from the channels
	// The select statement will block until one of the cases is ready
	// If multiple cases are ready, one will be chosen at random
	// If no cases are ready, the default case will be executed
	// In this case, we are using nil channels to check if the channels are closed
	// If a channel is closed, we set it to nil, and the select statement will ignore it
	// We keep looping until both channels are closed
	for {
		// Select the first channel that is ready
		select {
		// Receive a number from the numbers channel
		case number, ok := <-numbersChannel:
			// Check if the channel is open
			if ok {
				fmt.Printf("Received number: %d\n", number)
			} else {
				// If the channel is closed, set it to nil
				numbersChannel = nil
			}
		// Receive a letter from the letters channel
		case letter, ok := <-lettersChannel:
			// Check if the channel is open
			if ok {
				fmt.Printf("Received letter: %s\n", letter)
			} else {
				// If the channel is closed, set it to nil
				lettersChannel = nil
			}
		}

		// Check if both channels are closed
		if numbersChannel == nil && lettersChannel == nil {
			break
		}
	}
}

// sendNumbers sends integers to a channel
// It sends numbers from 1 to 5 to the channel
// It sleeps for 1 second between each number
func sendNumbers(ch chan int) {
	for i := 1; i <= 5; i++ {
		// Send the number to the channel
		ch <- i
		// Sleep for 1 second
		time.Sleep(1 * time.Second)
	}

	// Close the channel
	// ch will be nil after this
	close(ch)
}

// sendLetters sends letters to a channel
// It sends letters A to E to the channel
// It sleeps for 1 second between each letter
func sendLetters(ch chan string) {
	letters := []string{"A", "B", "C", "D", "E"}

	for _, letter := range letters {
		// Send the letter to the channel
		ch <- letter
		// Sleep for 1 second
		time.Sleep(1 * time.Second)
	}

	// Close the channel
	// ch will be nil after this
	close(ch)
}
