package main

// Import necessary packages
import (
	"fmt"
)

func main() {
	fmt.Println("Starting error handling and panic-recover examples...\n")

	// Call error handling examples
	errorExample()
	functionWithErrorExample()

	// Call panic and recover example
	panicRecoverExample()

	fmt.Println("\nProgram completed successfully.")
}
