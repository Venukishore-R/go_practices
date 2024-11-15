// IF-ELSE IF EXAMPLE IN GO

// The `if-else if` statement is used to check multiple conditions one after another.
// If one condition is true, the corresponding block of code is executed, and the rest are skipped.

package main

import "fmt"

func ifElseIfExample() {
	fmt.Println("\nIF-ELSE IF EXAMPLE")

	number := 1

	// First condition
	if number == 1 {
		fmt.Println("Number is 1")
	} else if number == 2 {
		// Second condition
		fmt.Println("Number is 2")
	} else if number == 3 {
		// Third condition
		fmt.Println("Number is 3")
	} else {
		// Default case when none of the above conditions are true
		fmt.Println("Number is neither 1, 2, nor 3")
	}
}
