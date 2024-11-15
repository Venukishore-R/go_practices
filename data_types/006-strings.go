package main

import "fmt"

func stringExample() {
	fmt.Print("\n\nSTRING - DATA TYPE EXAMPLES\n")

	// The "string" type in Go:
	// -------------------------
	// - A string is a read-only slice of bytes.
	// - It is used to represent a sequence of characters (text).
	// - Strings are immutable, meaning once created, they cannot be modified.

	// Strings can be initialized in two ways:

	// **METHOD 1: Using Double Quotes (" ")**
	// - This method honors escape sequences like:
	//   - \n (new line)
	//   - \t (tab)

	var myName string = "\tMy Name is Venukishore\n"
	fmt.Print("Output with escape sequences:\n", myName)

	// **METHOD 2: Using Back Quotes (` `)**
	// - This method treats the string as a raw literal and ignores escape sequences.
	// - Useful when you want to include raw text, such as file paths or multiline strings.

	var myCar string = `\tThis is a raw string literal\nIt does not interpret escape sequences.`
	fmt.Print("Output with raw string literal:\n", myCar, "\n")

	// Multiline strings using back quotes:
	var multiLineString string = `
	This is a multiline
	string using back quotes.
	It preserves the formatting.
	`
	fmt.Print("Multiline string example:\n", multiLineString)
}
