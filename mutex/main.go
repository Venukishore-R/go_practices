// MUTEX Example
// This program demonstrates the use of a mutex to prevent race conditions
// when multiple goroutines access shared data.

package main

import (
	"fmt"
	"sync"
	"time"
)

// isEven checks if a given number is even.
func isEven(number int) bool {
	return number%2 == 0
}

func main() {
	n := 0           // Shared variable to be accessed by goroutines
	var m sync.Mutex // Mutex to synchronize access to the shared variable

	// First goroutine to check if the number `n` is even or odd
	go func() {
		m.Lock()         // Lock the mutex before accessing shared variable `n`
		defer m.Unlock() // Ensure the mutex is unlocked after this function completes

		nIsEven := isEven(n)             // Check if `n` is even
		time.Sleep(5 * time.Millisecond) // Simulate some processing time

		if nIsEven {
			fmt.Println(n, "is even") // Print if `n` is even
		} else {
			fmt.Println(n, "is odd") // Print if `n` is odd
		}
	}()

	// Second goroutine to increment the shared variable `n`
	go func() {
		m.Lock()   // Lock the mutex before modifying shared variable `n`
		n++        // Increment `n`
		m.Unlock() // Unlock the mutex after modification
	}()

	// Wait for a second to allow goroutines to complete their execution
	time.Sleep(time.Second)
}
