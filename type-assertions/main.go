// TYPE ASSERTIONS

// Type assertion is used to specify the data type that a Go variable (usually of type `interface{}`) can hold.
// When you have an interface and you don't know its concrete data type, you can use type assertion to extract
// the underlying value with the specific type.

// Syntax:
// newVariable, ok := variable.(dataType)

// If the assertion fails (i.e., the data type doesn't match), `ok` will be `false`.
// Type assertion on a `nil` interface will cause a panic if not handled properly.

package main

import (
	"fmt"
	"log"
)

type Person struct {
	name string
	age  int
}

func main() {
	// Example 1: Type assertion with a string
	var greetings interface{} = "Hello world"

	isInt, ok := greetings.(int)
	fmt.Printf("Value: %v, Type Assertion Success: %v (Expected: false)\n", isInt, ok)

	isString, ok := greetings.(string)
	fmt.Printf("Value: %v, Type Assertion Success: %v (Expected: true)\n", isString, ok)

	// If the given data type is not present in the Go variable, the result will be the data type's zero value.
	isBool, ok := greetings.(bool)
	fmt.Printf("Value: %v, Type Assertion Success: %v (Expected: false)\n", isBool, ok)

	// Example 2: Type assertion with a struct
	assertionWithStructExample()

	// Example 3: Type assertion with a nil interface
	typeAssertionWithNilInterface()

	// Example 4: Type assertion with mismatched struct type
	typeAssertionWithFalseValues()
}

// Example 2: Type assertion with a struct
func assertionWithStructExample() {
	var p interface{} = Person{
		name: "Vk",
		age:  18,
	}

	person1, ok := p.(Person)
	fmt.Printf("Struct Value: %+v, Type Assertion Success: %v (Expected: true)\n", person1, ok)
}

// Example 3: Type assertion with a nil interface
func typeAssertionWithNilInterface() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered from panic:", r)
		}
	}()

	var p interface{}

	// Attempting type assertion on a nil interface will cause a panic if not checked with `ok`.
	person1, ok := p.(*Person)
	fmt.Printf("Value: %v, Type Assertion Success: %v (Expected: false)\n", person1, ok)
}

// Example 4: Type assertion with mismatched struct type
func typeAssertionWithFalseValues() {
	var p interface{} = struct {
		name    string
		address string
	}{
		name:    "Vk",
		address: "Some address",
	}

	// Attempting to assert to a different struct type will return zero value and false.
	person1, ok := p.(Person)
	fmt.Printf("Value: %+v, Type Assertion Success: %v (Expected: false)\n", person1, ok)
}
