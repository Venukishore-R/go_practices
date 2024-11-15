package main

func main() {
	// Go has two main categories of data types:

	// 1) **Basic Types**: These include simple, primitive types.
	//    - Examples: int, float, complex, string, bool, rune, byte

	// 2) **Composite Types**: These are types that are made up of other types.
	//    - Composite types are further classified into:
	//      - **Non-reference types**: These hold the actual data.
	//        - Examples: Slices, Maps, Channels, Functions, Interfaces, Pointers
	//      - **Reference types**: These hold references to data.
	//        - Examples: Arrays, Structs

	// Calling functions for each data type examples
	intExample()     // Example of basic int data type
	uIntExample()    // Example of unsigned int data type
	floatExample()   // Example of float data type
	complexExample() // Example of complex data type
	bytesExample()   // Example of byte data type
	runeExample()    // Example of rune data type (alias for int32)
	stringExample()  // Example of string data type
	boolExample()    // Example of bool data type
	arrayExample()   // Example of array data type
	sliceExample()   // Example of slice data type
	mapExample()     // Example of map data type
	makeExample()    // Example of make data type
	structExample()  // Example of struct data type
}
