// SELECT

// Select statement is also like switch
// But, it is used to select multiple channel operations

package main

import (
	"fmt"
	"time"
)

func main() {
	// Create two channels for string messages
	c1 := make(chan string)
	c2 := make(chan string)

	// Start a goroutine that sends a message to c1 after 1 second
	go func() {
		time.Sleep(1 * time.Second) // Simulate some work with sleep
		c1 <- "One"                  // Send message to channel 1
	}()

	// Start another goroutine that sends a message to c2 after 1 second
	go func() {
		time.Sleep(1 * time.Second) // Simulate some work with sleep
		c2 <- "Two"                  // Send message to channel 2
	}()

	// Loop to receive messages from both channels
	for i := 0; i < 2; i++ {
		select {
		// Case when a message is received from channel 1
		case msg1 := <-c1:
			fmt.Println("Message from Channel 1:", msg1)
		// Case when a message is received from channel 2
		case msg2 := <-c2:
			fmt.Println("Message from Channel 2:", msg2)
		}
	}
}