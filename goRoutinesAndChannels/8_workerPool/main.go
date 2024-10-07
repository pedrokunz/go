package main

import (
	"fmt"
	"time"
)

// main
// Objective: Implement a worker pool using goroutines and channels.
// Task: Write a program that creates a pool of worker goroutines to process tasks from a channel.
func main() {
	const numWorkers = 3

	// Create channels for tasks and results
	// The tasks channel will be used to send tasks to the workers
	// The results channel will be used to receive results from the workers
	tasks := make(chan int, 10)
	results := make(chan int, 10)

	// Create worker goroutines
	// Each worker will process tasks from the tasks channel
	// and send the results to the results channel
	for workerID := 1; workerID <= numWorkers; workerID++ {
		// This is a non-blocking operation
		go worker(workerID, tasks, results)
	}

	// Send tasks to the workers
	// We send 5 tasks to the workers that will process the tasks concurrently
	// and send the results back to the results channel
	for taskID := 1; taskID <= 5; taskID++ {
		tasks <- taskID
	}

	// Close the tasks channel
	// This is important to signal that no more tasks will be sent
	close(tasks)

	// Receive results from the workers
	// We receive 5 results from the workers
	// and print them to the console
	for resultID := 1; resultID <= 5; resultID++ {
		// Receive a result from the results channel
		result := <-results
		// Print the result to the console
		fmt.Printf("Received result: %d\n", result)
	}
}

// worker processes tasks from a channel
// It receives tasks from the tasks channel
// and sends the results to the results channel
// It simulates processing a task by sleeping for 1 second
// and doubling the task value
// <-chan is a receive-only channel
// chan<- is a send-only channel
func worker(id int, tasks <-chan int, results chan<- int) {
	for task := range tasks {
		fmt.Printf("Worker %d started task %d\n", id, task)
		time.Sleep(1 * time.Second)
		results <- task * 2
	}
}
