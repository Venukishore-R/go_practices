package main

import "fmt"

func boolExample() {
	fmt.Print("\n\nBOOLEAN - DATA TYPE EXAMPLE\n")

	// The "bool" type in Go:
	// -------------------------
	// - The data type is "bool" and it can only have two possible values: `true` or `false`.
	// - Booleans are commonly used for conditional logic and comparisons.
	// - The default value of a boolean variable is `false`.

	// Example 1: Boolean with default value
	var myBool bool
	fmt.Println("Default value of myBool:", myBool)

	// Example 2: Explicitly assigning a boolean value
	var myBool2 bool = true
	fmt.Println("Value of myBool2:", myBool2)

	// Example 3: Another boolean assignment
	var myBool3 bool = false
	fmt.Println("Value of myBool3:", myBool3)

	// Example 4: Using boolean expressions
	isEven := 10%2 == 0
	fmt.Println("Is 10 even?", isEven)
}
