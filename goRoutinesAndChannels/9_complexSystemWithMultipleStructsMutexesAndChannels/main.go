package main

import (
	"github.com/pedrokunz/go/goRoutinesAndChannels/9_complexSystemWithMultipleStructsMutexesAndChannels/internal"
	"time"
)

// main
// Objective: Implement a complex system with multiple structs, each having its own mutex and channels for communication.
// Specifications:
//  1. Define a Sensor struct that generates data and sends it to a channel.
//  2. Define a Processor struct that receives data from the Sensor and processes it.
//  3. Define a Logger struct that logs the processed data.
//  4. Use mutexes to protect shared resources within each struct.
//  5. Use channels to communicate between the structs.
func main() {
	// Create instances of the Sensor, Processor, and Logger
	sensor := internal.NewSensor(1)
	processor := internal.NewProcessor(1)
	logger := internal.NewLogger(1)

	// Connect the Sensor to the Processor
	processor.DataCh = sensor.DataCh
	// Connect the Processor to the Logger
	logger.ResultCh = processor.ResultCh

	// Start the Sensor, Processor, and Logger
	go sensor.GenerateData()
	go processor.ProcessData()
	go logger.LogResult()

	// Wait for all goroutines to finish
	time.Sleep(10 * time.Second)
}
