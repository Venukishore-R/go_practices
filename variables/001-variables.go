package main

import "fmt"

func main() {
	// First: Using the "var" keyword
	// Variables initialized without explicit data type are inferred from the value assigned
	var name = "name"   // inferred type is string
	var number = 3      // inferred type is int
	var isLogged = true // inferred type is bool

	// Explicit type declaration
	var name2 string = "name2"
	var number2 = 3 // inferred type is int
	var isLoggedIn bool = false

	// Second: Using "const" keyword
	// Constants cannot be changed after initialization
	const myName = "my name"
	const myNumber = 3
	const logged bool = true

	// Explicit type declaration for constants
	const myName2 = "my name 2"
	const myNumber2 = 3
	const loggedIn2 bool = false

	// Third: Using shorthand ":=" for declaration
	// This is the most concise way to declare and initialize variables in Go
	someName := "some name"
	someNumber := 4

	// Declaring variables with specific types and initializing later
	var age int
	var greetings string
	var married bool

	age = 10
	greetings = "Good morning, Venukishore R"
	married = true

	// Displaying variables to check their values
	fmt.Println("name:", name, "number:", number, "isLogged:", isLogged)
	fmt.Println("name2:", name2, "number2:", number2, "isLoggedIn:", isLoggedIn)
	fmt.Println("myName:", myName, "myNumber:", myNumber, "logged:", logged)
	fmt.Println("myName2:", myName2, "myNumber2:", myNumber2, "loggedIn2:", loggedIn2)
	fmt.Println("someName:", someName, "someNumber:", someNumber)
	fmt.Println("age:", age, "greetings:", greetings, "married:", married)
}
