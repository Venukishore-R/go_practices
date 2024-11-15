package main

import (
	"fmt"
	"unsafe"
)

func floatExample() {
	fmt.Println("\nFLOAT - DATA TYPE EXAMPLES\n")

	// Floating-point numbers represent real numbers with decimal points.
	// Keyword: "float"
	// Floating-point types in Go are classified into two types:
	// 1) float32
	// 2) float64

	// 1) float32
	// -------------------------
	// "float32" has a size of 32 bits (4 bytes).
	// It provides approximately 6-7 decimal digits of precision.
	// Use "float32" when memory is a constraint and less precision is acceptable.

	var numFloat32 float32
	fmt.Println("Size of numFloat32 in bytes:", unsafe.Sizeof(numFloat32))

	// 2) float64
	// -------------------------
	// "float64" has a size of 64 bits (8 bytes).
	// It provides approximately 15-16 decimal digits of precision.
	// Use "float64" when higher precision is needed (default type for floating-point literals).

	var numFloat64 float64
	b := 2.5 // The default type for floating-point numbers is "float64".

	fmt.Println("Size of numFloat64 in bytes:", unsafe.Sizeof(numFloat64))
	fmt.Println("Size of b in bytes:", unsafe.Sizeof(b))
}
