// ONLY IF EXAMPLE IN GO

// In Go, you can use the `if` statement independently without an `else` block.
// It allows conditional execution of code based on a boolean expression.

package main

import "fmt"

func onlyIfExample() {
	fmt.Println("\nONLY IF EXAMPLE")

	// Example with a boolean condition
	condition := true

	// `if` statement with a true condition
	if condition {
		fmt.Println("Condition is true, so this line is executed.")
	}

	// `if` statement with a negated condition
	if !condition {
		fmt.Println("Condition is false, so this line is NOT executed.")
	}

	// This line executes regardless of the condition
	fmt.Println("This line executes after the if statements.")
}
