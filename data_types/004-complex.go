package main

import (
	"fmt"
	"unsafe"
)

func complexExample() {
	fmt.Println("\nCOMPLEX - DATA TYPE EXAMPLES\n")

	// Complex numbers are a combination of real and imaginary parts.
	// The real part is a floating-point number, and the imaginary part is followed by 'i'.
	// Complex numbers in Go are classified into two types:
	// 1) complex64
	// 2) complex128

	// 1) complex64
	// -------------------------
	// "complex64" consists of two "float32" values (real and imaginary parts).
	// It uses 32 bits for each part, so the total size is 64 bits (8 bytes).

	var complexNum64 complex64 = complex(3.5, 4.5)
	fmt.Println("complexNum64:", complexNum64)
	fmt.Println("Size of complexNum64 in bytes:", unsafe.Sizeof(complexNum64))

	// 2) complex128
	// -------------------------
	// "complex128" consists of two "float64" values (real and imaginary parts).
	// It uses 64 bits for each part, so the total size is 128 bits (16 bytes).
	// Typically used for high-precision calculations, such as 2D rotations or electrical engineering applications.

	complexNum1 := complex(5, 5.5) // Using the "complex" function
	complexNum2 := 5 + 6i          // Direct assignment

	fmt.Println("complexNum1:", complexNum1)
	fmt.Println("Size of complexNum1 in bytes:", unsafe.Sizeof(complexNum1))

	fmt.Println("complexNum2:", complexNum2)
	fmt.Println("Size of complexNum2 in bytes:", unsafe.Sizeof(complexNum2))
}