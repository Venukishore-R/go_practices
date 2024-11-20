package main

import "fmt"

// uniDirectionalChannel demonstrates the use of a unidirectional channel.
// Unidirectional channels can either send or receive data, but not both.
func uniDirectionalChannel() {
	// Informing the start of the channel writing process
	fmt.Println("Writing to uniDirectionalChannel started")

	// Create a unidirectional channel for sending integers
	// The channel can hold one integer value
	uDChannel := make(chan<- int, 1)

	// Send the integer value 190 to the unidirectional channel
	uDChannel <- 190

	// Informing the completion of the channel writing process
	fmt.Println("Writing to uniDirectionalChannel completed")

	// Note: Receiving from a unidirectional channel is not allowed.
	// The following line is commented out to prevent compilation error:
	// fmt.Println(<-uDChannel)
}