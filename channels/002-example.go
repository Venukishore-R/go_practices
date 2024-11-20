package main

import "fmt"

// deadlockExample demonstrates a deadlock situation in Go.
func deadlockExample() {
	fmt.Println("\nDemonstrating Deadlock...")

	// Declaring an unbuffered channel that can hold integers
	deadlockChannel := make(chan int)

	// Attempting to send data into the unbuffered channel without a receiver
	// This will cause a deadlock because there is no goroutine ready to receive the data
	// Uncommenting the line below will result in a runtime deadlock
	deadlockChannel <- 1

	// This message will not print if a deadlock occurs
	fmt.Println("This message will not print if deadlock occurs!")
}