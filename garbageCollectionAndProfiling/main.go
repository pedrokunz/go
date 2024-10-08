package main

import (
	"fmt"
	"math/rand"
	"net/http"
	_ "net/http/pprof" // Import pprof for profiling
	"runtime"
	"time"
)

const memoryThreshold = 100 // Memory threshold in MiB

// Global slice to hold references to allocated memory
var allocatedMemory [][]byte

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Allocate 100 MiB in chunks of 1 MiB to 10 MiB
		// This will cause the memory usage to increase
		// and trigger garbage collection when the memory threshold is exceeded
		// 100 << 20 is equivalent to 100 * 2^20 or 100 * 1024 * 1024 (104,857,600 bytes)
		// the << operator is the bit shift left operator
		// 20 is the number of bits to shift left
		totalSize := 100 << 20 // 100 MiB
		data := make([]byte, totalSize)
		allocatedMemory = append(allocatedMemory, data) // Keep reference to allocated memory

		fmt.Printf("Allocated %v MiB\n", bytesToMiB(uint64(totalSize)))

		// Simulate resource cleanup after some time
		go func() {
			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
			allocatedMemory = nil // Clear the allocated memory

			fmt.Println("Deallocated memory")
		}()

		// Write a response
		_, _ = w.Write([]byte(fmt.Sprintf("Allocated %v MiB\n", bytesToMiB(uint64(totalSize)))))
	})

	fmt.Println("Starting server on :8080")

	go func() {
		for {
			var memoryStats runtime.MemStats

			runtime.ReadMemStats(&memoryStats)

			allocMiB := bytesToMiB(memoryStats.Alloc)
			fmt.Printf(
				"Alloc: %v MiB\tTotalAlloc: %v MiB\tSys: %v MiB\tNumGC: %v\n",
				allocMiB,
				bytesToMiB(memoryStats.TotalAlloc),
				bytesToMiB(memoryStats.Sys),
				memoryStats.NumGC,
			)

			// Trigger garbage collection if memory usage exceeds the threshold
			if allocMiB >= memoryThreshold {
				runtime.GC()
				fmt.Println("Manually triggered garbage collection")
			}

			time.Sleep(10 * time.Second)
		}
	}()

	// Start the pprof server
	go func() {
		fmt.Println("Starting pprof server on :6060")
		http.ListenAndServe(":6060", nil)
	}()

	http.ListenAndServe(":8080", nil)
}

// bytesToMiB converts bytes to megabytes
// 1 MiB = 1024 * 1024 bytes
// Example: bytesToMiB(1048576) => 1
func bytesToMiB(b uint64) uint64 {
	return b / 1024 / 1024
}

// Output:
// Starting server on :8080
// Alloc: 0 MiB		TotalAlloc: 0 MiB	Sys: 6 MiB		NumGC: 0
// Allocated 100 MiB
// Deallocated memory
// Alloc: 100 MiB	TotalAlloc: 100 MiB	Sys: 107 MiB	NumGC: 1
// Manually triggered garbage collection
// Alloc: 0 MiB		TotalAlloc: 100 MiB	Sys: 107 MiB	NumGC: 2

// Visit http://localhost:6060/debug/pprof/heap to view the heap profile
