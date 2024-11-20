// CHANNELS
// Channels in Go are like pipes used to communicate between Go routines.
// They allow data to be passed between routines and the main program.

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Print("\nCHANNELS\n\n")

	// Declaring a buffered channel with a capacity of 1
	ageChannel := make(chan int, 1)

	// Writing a value to the channel (sending data)
	ageChannel <- 10

	// Reading the value from the channel (receiving data)
	actualAge := <-ageChannel
	fmt.Println("Age:", actualAge)

	// Example - Synchronizing Go routines using an unbuffered channel
	done := make(chan bool) // Unbuffered channel for synchronization

	// Launching the `hello` Go routine
	go hello(done)

	// Wait for the `hello` function to signal completion
	<-done

	// Call function to calculate results (not defined in the original code)
	calculateResult()

	// Uncomment to see the deadlock example
	// deadlockExample()

	// Demonstrating a uni-directional channel
	uniDirectionalChannel()

	// Demonstrating channel closure
	channelWithClose()

	fmt.Println("Main function completed")
}

// Function executed as a Go routine
func hello(done chan bool) {
	fmt.Println("Hello from the hello program")
	fmt.Println("The Go routine is going to sleep for 4 seconds...")

	// Simulating a delay of 4 seconds
	time.Sleep(4 * time.Second)

	fmt.Println("The Go routine has awakened")

	// Sending a signal to the main function to indicate completion
	done <- true 
}