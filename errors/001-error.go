// ERRORS

// Error is a value returned when an unusual or unexpected situation occurs in the program.

// In Go, we can create our own custom errors for user-defined errors.

// Unlike Java, JavaScript, or Flutter, where errors are handled using try/catch blocks,
// Go treats errors like other variables and handles them by returning error values from functions.

// Errors can be handled by returning a message to the user or shutting down the program using `log.Fatal()`.

package main

import (
	"errors"
	"fmt"
	"log"
)

func errorExample() {
	fmt.Println("\nCUSTOM ERROR EXAMPLES")

	// Creating a custom error using `errors.New()`
	// Error string should start with a lowercase letter as per Go conventions
	newError := errors.New("user-defined error occurred")
	fmt.Println(newError.Error())

	// Creating a custom error using `fmt.Errorf()`
	var someError error = fmt.Errorf("an example of a formatted error")
	fmt.Println(someError.Error())
}

// Example of a function that returns an error
func divide(a, b int) (int, error) {
	if b == 0 {
		// Return a custom error if division by zero is attempted
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func functionWithErrorExample() {
	fmt.Println("\nFUNCTION ERROR HANDLING EXAMPLE")

	// Calling a function that may return an error
	result, err := divide(10, 0)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Division result: %d\n", result)

	// Successful function call
	result, err = divide(10, 2)
	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}
	fmt.Printf("Successful division result: %d\n", result)
}
