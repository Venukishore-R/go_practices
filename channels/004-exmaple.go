package main

import "fmt"

// channelWithClose demonstrates sending values through a channel and closing it.
func channelWithClose() {
	// Create a channel of integers
	ch := make(chan int)

	// Start the producer goroutine to send values to the channel
	go producer(ch)

	// Continuously receive values from the channel
	for {
		// Receive value from the channel and check if the channel is closed
		v, ok := <-ch
		if !ok { // If ok is false, the channel is closed
			fmt.Println("Channel is closed")
			break // Exit the loop
		}
		// Print the received value
		fmt.Println("Value from channel:", v)
	}
}

// producer sends integers to the provided channel and then closes it
func producer(sendChan chan int) {
	// Send integers from 0 to 9 to the channel
	for i := 0; i < 10; i++ {
		sendChan <- i
	}
	// Close the channel to indicate no more values will be sent
	close(sendChan)
}