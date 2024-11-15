package main

import (
	"fmt"
	"unsafe"
)

func uIntExample() {
	fmt.Println("\nUINT - DATA TYPE EXAMPLES\n")

	// uint - represents an unsigned integer data type
	// Keyword: "uint"
	// Unsigned integers can store only zero or positive values (no negative values).
	// Similar to "int", it is classified into five types: uint, uint8, uint16, uint32, uint64.

	// 1) uint
	// -------------------------
	// The size of "uint" depends on the platform:
	// - On 32-bit machines, the size of "uint" is 32 bits (4 bytes).
	// - On 64-bit machines, the size of "uint" is 64 bits (8 bytes).
	// - The range of "uint" on 32-bit machines is from 0 to 2^32 - 1.
	// - The range of "uint" on 64-bit machines is from 0 to 2^64 - 1.

	var numUint uint
	b := 2 // The default type for unsigned integers is "uint".

	fmt.Println("Size of numUint in bytes:", unsafe.Sizeof(numUint))
	fmt.Println("Size of b in bytes:", unsafe.Sizeof(b))

	// 2) uint8
	// -------------------------
	// "uint8" has a size of 8 bits (1 byte).
	// The range of "uint8" is from 0 to 2^8 - 1 (0 to 255).
	// Typically used for small integer values or loop counters.

	var numUInt8 uint8
	fmt.Println("Size of numUInt8 in bytes:", unsafe.Sizeof(numUInt8))

	// 3) uint16
	// -------------------------
	// "uint16" has a size of 16 bits (2 bytes).
	// The range of "uint16" is from 0 to 2^16 - 1 (0 to 65,535).

	var numUInt16 uint16
	fmt.Println("Size of numUInt16 in bytes:", unsafe.Sizeof(numUInt16))

	// 4) uint32
	// -------------------------
	// "uint32" has a size of 32 bits (4 bytes).
	// The range of "uint32" is from 0 to 2^32 - 1 (0 to 4,294,967,295).

	var numUInt32 uint32
	fmt.Println("Size of numUInt32 in bytes:", unsafe.Sizeof(numUInt32))

	// 5) uint64
	// -------------------------
	// "uint64" has a size of 64 bits (8 bytes).
	// The range of "uint64" is from 0 to 2^64 - 1.
	// Commonly used when a larger range is needed (e.g., time.Duration).

	var numUInt64 uint64
	fmt.Println("Size of numUInt64 in bytes:", unsafe.Sizeof(numUInt64))
}
