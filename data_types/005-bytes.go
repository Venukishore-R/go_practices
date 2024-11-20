package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func bytesExample() {
	fmt.Print("\nBYTES - DATA TYPE EXAMPLE\n\n")

	// The "byte" type in Go:
	// -------------------------
	// - "byte" is an alias for "uint8" (unsigned integer of 8 bits/1 byte size).
	// - It is used to represent ASCII characters or raw binary data.

	// **USE CASES:**
	// - Storing and handling raw data (e.g., images, audio, videos).
	// - Reading and writing files (especially in buffers).
	// - Encoding and decoding data (e.g., in network communication).
	// - String manipulation (e.g., converting between strings and byte slices).

	var newByte byte = 'a' // Storing an ASCII character as a byte.

	fmt.Println("Size of newByte in bytes:", unsafe.Sizeof(newByte))
	fmt.Println("Type of newByte:", reflect.TypeOf(newByte))
	fmt.Println("ASCII value of 'a':", newByte)
	fmt.Printf("Character representation: %c\n", newByte)
}
