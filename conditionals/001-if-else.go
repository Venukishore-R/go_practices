// IF-ELSE EXAMPLE IN GO

// The `if-else` statement is used for conditional execution of code based on whether a condition is true or false.

package main

import "fmt"

func ifelseExample() {
	fmt.Println("\nIF-ELSE EXAMPLE")

	someNumber := 5

	// Basic `if-else` syntax
	if someNumber%2 == 0 {
		fmt.Printf("Given number %d is even\n", someNumber)
	} else {
		fmt.Printf("Given number %d is odd\n", someNumber)
	}

	// Using an initializer in `if` statement
	if num := 10; num > 0 {
		fmt.Printf("Number %d is positive\n", num)
	} else {
		fmt.Printf("Number %d is non-positive\n", num)
	}
}
