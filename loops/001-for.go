package main

import "fmt"

// LOOPS
// Loops are used to execute a process or commands multiple times.

// FOR LOOP
// Used to execute a process for a fixed number of times, such as for 5 or 10 times.

func forExample() {
	fmt.Print("\n\nFOR LOOP\n")

	// THREE COMPONENT LOOP
	// This loop has three parts: initialization, condition, and increment/decrement.

	// Example: Looping from 0 to 9 (10 iterations)
	for i := 0; i < 10; i++ {
		fmt.Println("For loop: ", i)
	}

	// WHILE LOOP
	// Without the initialization and increment/decrement parts, you can simulate a while loop.

	// Example: While loop that prints numbers until n is less than 10
	n := 0
	for n < 10 {
		fmt.Println("While loop: ", n)
		n += 3  // Incrementing n by 3
	}

	// INFINITE LOOP
	// This loop runs indefinitely. Use with caution.
	// It will never exit unless you explicitly stop it with a condition, break, or other control flow.
	times := 0
	for {
		fmt.Println("Infinite loop: Hello, World!")
		if times == 10 {
			break // Exit the loop when times equals 10
		}

		times++
	}

	// EXITING EARLY
	// You can use 'break' to exit the loop before the condition ends,
	// or 'continue' to skip the current iteration and proceed with the next one.
}
