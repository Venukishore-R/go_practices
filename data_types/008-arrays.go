package main

import "fmt"

func arrayExample() {
	fmt.Print("\n\nARRAY - DATA TYPE EXAMPLE\n")

	// Arrays in Go:
	// --------------
	// - An array is a collection of elements with the **same data type**.
	// - The size of an array is fixed and cannot be changed once declared.

	// **Basic Syntax:**
	// arrayName := [size]DATA_TYPE{values}

	// **METHOD 1: Declaring an Array with a Specific Size**
	var marks [5]int    // Array of integers with size 5
	var names [5]string // Array of strings with size 5

	// Assigning values to array elements
	names[0] = "John"
	marks[0] = 95

	fmt.Println("Marks:", marks)
	fmt.Println("Names:", names)

	// **METHOD 2: Declaring and Initializing an Array**
	var cars = [5]string{"Toyota", "Honda"} // Partially initialized array

	cars[3] = "Ford" // Adding a value at index 3

	numbers := [3]float64{1.1, 2.2} // Array of float64 with 3 elements
	numbers[2] = 3.3

	fmt.Println("Cars:", cars)
	fmt.Println("Numbers:", numbers)

	// **METHOD 3: Declaring an Array without Specifying Size**
	// - Use `[...]` to let Go determine the size based on the number of elements provided.
	ages := [...]int{18, 21, 25, 30}
	fmt.Println("Ages:", ages)

	// **Built-in Functions for Arrays**

	// `len()` - Returns the length of an array
	lengthOfAgesArray := len(ages)
	fmt.Println("Length of ages array:", lengthOfAgesArray)

	// Accessing array elements using index
	fmt.Println("First element of ages array:", ages[0])

	// Iterating over an array using a loop
	fmt.Println("Iterating over the 'cars' array:")
	for i := 0; i < len(cars); i++ {
		fmt.Printf("cars[%d] = %s\n", i, cars[i])
	}
}
