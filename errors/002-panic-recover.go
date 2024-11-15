// PANIC AND RECOVER

// `panic` signifies a critical situation that stops the normal execution flow of a program.
// It is different from errors because errors can be handled, while `panic` is executed
// when the program encounters an unrecoverable situation.

// Causes of `panic`:
// 1. Runtime errors (e.g., nil pointer dereference, out-of-bounds array access, division by zero).
// 2. Explicit calls to `panic()` made by the programmer to stop the execution.

// RECOVER
// `recover()` is a function used to handle `panic`. It allows the program to regain control.
// `recover()` can only be used inside a deferred function. If there is no deferred function, the `panic` cannot be handled.

package main

import (
	"fmt"
)

func panicRecoverExample() {
	fmt.Println("Starting panic and recover example")

	// Example with explicit call to panic
	explicitPanicExample()

	// Example with runtime panic (division by zero)
	runtimePanicExample()

	// Example with nil pointer dereference
	nilPointerDeferenceExample()

	fmt.Println("Program completed successfully")
}

func explicitPanicExample() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from explicit panic:", r)
		}
	}()

	fmt.Println("\nCalling explicit panic example")
	panic("This is an explicit panic")
	fmt.Println("This line will not be executed after panic")
}

func runtimePanicExample() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from runtime panic:", r)
		}
	}()

	fmt.Println("\nCalling runtime panic example")

	// Division by zero will cause a runtime panic
	num := 10
	denom := 0
	result := num / denom // This will trigger a panic
	fmt.Println("Division result:", result)
}

// Nil pointer deference panic will be caused when the pointer as a nil value in it
func nilPointerDeferenceExample() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from nil pointer dereference:", r)
		}
	}()

	type Person struct {
		name string
		age  int
	}

	var p *Person

	// Dereferencing a nil pointer will cause a panic
	// You can handle it like this

	// if p == nil {
	// 	fmt.Println("Warning: Attempt to access a nil pointer. Skipping operation.")
	// 	return
	// }

	fmt.Println("\nCalling nil pointer", p.name)

	fmt.Println("Rest of the program executed in nil pointer deference example")
}
