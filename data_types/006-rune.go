package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func runeExample() {
	fmt.Print("\nRUNE - DATA TYPE EXAMPLE\n")

	// The "rune" type in Go:
	// -------------------------
	// - "rune" is an alias for "int32" (32-bit integer).
	// - It is used to represent Unicode code points (characters).
	// - A Unicode code point is a unique identifier for every character, including letters, digits, symbols, and punctuation from different languages.

	// **UNICODE:**
	// - Unicode is a universal character encoding standard used for text representation.
	// - It covers almost all characters in the world, including special symbols.

	// **UTF-8 Encoding:**
	// - UTF-8 is a variable-length encoding system for Unicode characters.
	// - It uses 1 to 4 bytes to store characters:
	//   - 1 byte (8 bits) for ASCII characters (e.g., 'a', '$').
	//   - Up to 4 bytes (32 bits) for more complex Unicode characters (e.g., emojis).

	// Example of a rune (Unicode character)
	rPound := '$'

	fmt.Printf("Unicode Code Point: %U\n", rPound) // Prints the Unicode code point
	fmt.Println("Rune value:", rune(rPound))
	fmt.Println("Type of rPound:", reflect.TypeOf(rPound))
	fmt.Printf("Character representation: %c\n", rPound)
	fmt.Println("Size of rPound in bytes:", unsafe.Sizeof(rPound))

	// Note:
	// - You should use the "rune" type when working with Unicode characters,
	//   especially when handling multilingual text or special symbols.
}
