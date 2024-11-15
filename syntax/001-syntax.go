// Package declaration is required at the top of every Go file.
// `package main` is a special package that indicates an executable program.
package main

// Importing the "fmt" package, which provides functions for formatted I/O operations.
import "fmt"

// The `main` function is the entry point of a Go program.
// It is mandatory for standalone programs, and it gets executed first.
func main() {
	// Using the `fmt` package to print a message to the console.
	// `fmt.Println()` prints the output followed by a newline.
	fmt.Println("Hello, world!")
}
