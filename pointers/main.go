// POINTERS
// Pointers are used to pass the references (addresses) of a variable.
package main

func main() {
	// Declare a string variable 'a' and initialize it with a value.
	a := "Venukishore"

	// Create a pointer 'b' that holds the address of variable 'a'.
	b := &a

	// Print the address stored in pointer 'b'.
	println(b)

	// Dereference pointer 'b' to get the value of 'a' and print it.
	println(*b)

	// Change the value of 'a' using the pointer 'b'.
	*b = "Kishore"

	// Print the address of pointer 'b' again.
	println(b)

	// Print the updated value of 'a' to confirm the change.
	println(a)
}