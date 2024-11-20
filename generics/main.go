// GENERICS
// Go is a statically typed language, meaning we need to define specific data types for functions, methods, variables, and data structures.
// In some cases, we may not know which data type to use (e.g., string or int).
// Generics allow us to use different data types for a function, variable, or data structure.
// Syntax for defining a generic function:
//	func funcName[genericName any](parameters) {
//		code
//	}

package main

import (
	"fmt"
	"reflect"
)

// SomeInterface defines an interface with a single method Work
type SomeInterface interface {
	Work()
}

// Worker struct implements the SomeInterface
type Worker struct {
	name string
}

// Work method for the Worker type
func (w Worker) Work() {
	fmt.Println("Now it is working")
}

func main() {
	// Sample data of various types
	cars := []string{"Volvo", "Benz", "Audi"}
	numbers := []int{1, 2, 3, 4, 5}

	// Displaying different types of data using the show function
	show(cars)      // Show string slice
	show(numbers)   // Show int slice
	show([]bool{true, false}) // Show bool slice

	// Show Worker structs
	show([]Worker{
		{name: "Vl"},
		{name: "Kishore"},
	})

	// Create generic slices for integers and strings
	gNum := GenericSlice[int]{1, 2, 3}
	gString := GenericSlice[string]{"Volvo", "Audi"}

	// Print the contents of the generic slices
	gNum.Print()
	gString.Print()

	// Output the contents of the generic slices using the Out function
	Out(gNum)
	Out(gString)
}

// Print1 and Print2 are specific functions for printing slices of different types
// These functions demonstrate the need for generics to avoid code duplication
func Print1(s []string) {
	for _, v := range s {
		fmt.Println(v) // Print each string in the slice
	}
}

func Print2(s []int) {
	for _, v := range s {
		fmt.Println(v) // Print each int in the slice
	}
}

// show is a generic function that accepts a slice of any type and prints its elements
func show[T any](s []T) {
	for _, v := range s {
		// Print the type of the element and its value
		fmt.Println(reflect.TypeOf(v)) // Print the type of the element
		fmt.Println(v)                 // Print the value of the element
	}
}

// GenericSlice is a type that represents a slice of any type T
type GenericSlice[T any] []T

// Print method for GenericSlice that prints each element in the slice
func (s GenericSlice[T]) Print() {
	for _, v := range s {
		fmt.Println(v) // Print each element in the generic slice
	}
}

// Out function that takes a GenericSlice of any type and prints its elements prefixed with "Out"
func Out[T any](g GenericSlice[T]) {
	for _, v := range g {
		fmt.Println("Out", v) // Print each element with a prefix
	}
}