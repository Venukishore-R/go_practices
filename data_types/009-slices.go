// SLICES

// Slices are similar to arrays but are more powerful and flexible.
// They are used to store multiple values of the same type in a single variable.

package main

import (
	"fmt"
)

func sliceExample() {
	fmt.Print("\n\nSLICES - DATA TYPE EXAMPLES\n")

	// There are three ways to declare slices:

	// Method 1: Using a slice literal

	// Syntax: sliceName := []dataType{values}

	marks := []float64{97, 96.5, 100}
	fmt.Println("Marks:", marks)

	// Method 2: Creating a slice from an array

	// Syntax: sliceName := array[startIndex:endIndex]

	numbers := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numberSlice := numbers[0:5] // Slice from index 0 to 4 (exclusive)
	fmt.Println("Number Slice:", numberSlice)

	// Method 3: Using the `make()` function

	// Syntax: sliceName := make([]dataType, length, capacity)

	newSlice := make([]string, 5) // Slice of length 5 with zero values
	fmt.Println("New Slice:", newSlice)

	// Accessing elements

	// Syntax: variable = sliceName[index]

	fmt.Println("First element of marks slice:", marks[0])

	// Updating elements

	// Syntax: sliceName[index] = newValue

	names := []string{"a", "b", "c"}
	fmt.Println("Before updating:", names)
	names[0] = "f"
	fmt.Println("After updating:", names)

	// Appending elements

	// Syntax: sliceName = append(sliceName, values...)

	names = append(names, "l")
	fmt.Println("Names after append:", names)

	// Appending one slice to another

	newNames := []string{"z", "y", "x"}
	names = append(names, newNames...)
	fmt.Println("After adding new names:", names)

	// Copying slices

	// The `copy()` function is used to copy one slice to another.

	// Syntax: copy(targetSlice, sourceSlice)

	mySlice := make([]string, len(names))
	copy(mySlice, names)
	fmt.Println("My Slice:", mySlice)

	// Length and Capacity of a Slice

	// `len(slice)` gives the length (number of elements).
	// `cap(slice)` gives the capacity (total size of the underlying array).

	fmt.Println("Length of mySlice:", len(mySlice))
	fmt.Println("Capacity of mySlice:", cap(mySlice))
}
