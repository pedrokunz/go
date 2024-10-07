package main

import (
	"fmt"
	"time"
)

// main
// Objective: Learn how to create and manage basic goroutines.
// Task: Write a program that launches a goroutine to print numbers from 1 to 5 with a delay of 1 second between each print.
func main() {
	fmt.Println("STARTING")

	go printNumbers()

	time.Sleep(6 * time.Second) // Wait for the goroutine to finish

	fmt.Println("FINISHING")
}

// printNumbers prints numbers from 1 to 5
// It sleeps for 1 second between each number
// This is a non-blocking operation
// It will run in a separate goroutine
func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}
