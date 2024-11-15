// NESTED IF EXAMPLE IN GO

// Nested `if` statements allow you to place one `if` inside another `if`.
// This is useful for checking multiple dependent conditions.

package main

import "fmt"

func nestedIf() {
	fmt.Println("\nNESTED IF EXAMPLE")

	number1 := 1
	number2 := 2

	// Outer `if` condition
	if number1 == 1 {
		// Inner `if` condition
		if number2 == 2 {
			fmt.Println("Both number1 and number2 conditions are true.")
		} else {
			fmt.Println("Number1 is true, but number2 is false.")
		}
	} else {
		fmt.Println("Number1 is false, so the inner if is skipped.")
	}
}
