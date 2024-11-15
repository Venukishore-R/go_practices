// TYPE INFERENCE
//
// Type inference is when the compiler automatically determines the type of a variable
// based on the value assigned to it, without explicitly specifying the type.
//
// In Go, you can use either:
// - `var` keyword with `=`
// - Short variable declaration with `:=`
//
// The data type is inferred from the value on the right side of the assignment.

package main

import (
	"fmt"
	"reflect"
)

func typeInferenceExample() {
	// Using `var` keyword (type is inferred as float64)
	var someNumber = 4.5

	// Using `:=` (type is inferred as string)
	someName := "Venukishore"

	// Using `:=` with integer (type is inferred as int)
	count := 10

	// Using `:=` with boolean (type is inferred as bool)
	isActive := true

	// Printing the types of variables using reflect.TypeOf()
	fmt.Println("Type of someNumber:", reflect.TypeOf(someNumber))
	fmt.Println("Type of someName:", reflect.TypeOf(someName))
	fmt.Println("Type of count:", reflect.TypeOf(count))
	fmt.Println("Type of isActive:", reflect.TypeOf(isActive))
}
