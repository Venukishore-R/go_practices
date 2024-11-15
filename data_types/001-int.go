package main

import (
	"fmt"
	"unsafe"
)

func intExample() {
	fmt.Print("\n\nINT - DATA TYPE EXAMPLES\n")

	// int - represents an integer data type
	// An integer can hold negative, zero, or positive values.
	// It is classified into "signed int" and "unsigned int".

	// SIGNED INT
	// -------------------------
	// A signed integer can hold negative, zero, or positive values.
	// Example:

	// var number int = 12
	// var number2 int = -12
	// var number3 = 0

	// Signed integers are further classified into 5 types based on size:
	// 1) int (default), 2) int8, 3) int16, 4) int32, 5) int64

	// 1) int
	// -------------------------
	// The size of an "int" depends on the platform:
	// - On 32-bit machines, the size of an "int" is 32 bits (4 bytes).
	// - On 64-bit machines, the size of an "int" is 64 bits (8 bytes).
	// - The range of "int" on 32-bit machines is from -2^31 to 2^31 - 1.
	// - The range of "int" on 64-bit machines is from -2^63 to 2^63 - 1.

	var number3 int
	b := 2 // The default type for integers is "int".

	fmt.Println("Size of number3 in bytes:", unsafe.Sizeof(number3))
	fmt.Println("Size of b in bytes:", unsafe.Sizeof(b))

	// 2) int8
	// -------------------------
	// "int8" has a size of 8 bits (1 byte).
	// The range of "int8" is from -2^7 to 2^7 - 1 (-128 to 127).
	// Typically used for small integer values or loop counters.

	var numInt8 int8
	fmt.Println("Size of numInt8 in bytes:", unsafe.Sizeof(numInt8))

	// 3) int16
	// -------------------------
	// "int16" has a size of 16 bits (2 bytes).
	// The range of "int16" is from -2^15 to 2^15 - 1 (-32,768 to 32,767).

	var numInt16 int16
	fmt.Println("Size of numInt16 in bytes:", unsafe.Sizeof(numInt16))

	// 4) int32
	// -------------------------
	// "int32" has a size of 32 bits (4 bytes).
	// The range of "int32" is from -2^31 to 2^31 - 1 (-2,147,483,648 to 2,147,483,647).

	var numInt32 int32
	fmt.Println("Size of numInt32 in bytes:", unsafe.Sizeof(numInt32))

	// 5) int64
	// -------------------------
	// "int64" has a size of 64 bits (8 bytes).
	// The range of "int64" is from -2^63 to 2^63 - 1.
	// Commonly used when a larger range is needed (e.g., time.Duration).

	var numInt64 int64
	fmt.Println("Size of numInt64 in bytes:", unsafe.Sizeof(numInt64))
}
