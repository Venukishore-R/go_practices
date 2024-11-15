package main

import "fmt"

// **FUNCTIONS**
// Functions are used to execute a set of statements or code blocks without rewriting them each time.
// Once a code block is written inside a function, it can be reused anywhere by calling the function by its name.
// You can create your own function with or without a return type.
// Additionally, you can pass multiple parameters to the function.
// Since Go is a statically typed language, you need to specify the data type of both the parameters and the return value.

func main() {
	// Calling deferExample with the defer keyword.
	// This ensures that the function is executed after the surrounding function (main) finishes execution.
	defer deferExample()

	// Calling a function with only a return type and printing the result
	resultFunc1 := funcWithOnlyReturn()
	fmt.Println(resultFunc1)

	// Calling a function that returns multiple values (string and error)
	resultFunc2, resultFunc2Error := funcWithMultipleReturn()
	fmt.Println(resultFunc2, resultFunc2Error)

	// Calling a function with named return values and printing them
	greetings, err := funcWithNamedReturns()
	fmt.Println(greetings, " ", err)

	// Calling a function that takes a single parameter
	funcWithSingleParam("Venukishore")

	// Calling a function that takes multiple parameters
	funcWithMultipleParams("Hello", "Harish")
}

// funcWithOnlyReturn demonstrates a function that returns a single string.
func funcWithOnlyReturn() string {
	return "Hello from Venukishore"
}

// funcWithMultipleReturn demonstrates a function that returns multiple values (string and error).
func funcWithMultipleReturn() (string, error) {
	return "No error while calling", nil
}

// funcWithNamedReturns demonstrates a function that uses named return values.
func funcWithNamedReturns() (greetings string, err error) {
	greetings = "Function with named return"
	return
}

// funcWithSingleParam demonstrates a function that accepts a single string parameter.
func funcWithSingleParam(name string) {
	fmt.Println("Hello ", name)
}

// funcWithMultipleParams demonstrates a function that accepts two string parameters.
func funcWithMultipleParams(greetings, name string) {
	fmt.Println(greetings + " " + name)
}

// DEFER EXAMPLE
// The defer keyword postpones the execution of the function until the surrounding function finishes execution.
func deferExample() {
	fmt.Println("This function is called first. This function will be executed last")
}
