// MAPS

// Maps are used to store values as key-value pairs.
// They provide flexibility for adding and removing elements and do not allow duplicate keys.

// Valid map key types include:
// - Booleans, Numbers, Strings, Arrays, Pointers, Structs, Interfaces (if the dynamic type supports equality)

// Invalid key types include:
// - Slices, Maps, Functions (because they do not support equality)

// Maps are reference types, so changes to one map variable will affect others if they reference the same underlying data.

package main

import "fmt"

func mapExample() {
	fmt.Println("\n\nMAPS - DATA TYPE EXAMPLES")

	// There are two ways to create maps:

	// Method 1: Using a map literal

	// Syntax: mapName := map[keyType]valueType{key: value}

	cars := map[string]int{
		"Benz":  10000,
		"Volvo": 20000,
		"Tata":  30000,
	}
	fmt.Println("Cars:", cars)

	// Method 2: Using the `make()` function

	// Syntax: mapName := make(map[keyType]valueType)

	marks := make(map[string]int)

	// Adding or updating elements

	// Syntax: mapName[key] = value

	marks["Vk"] = 96
	marks["G"] = 98
	fmt.Println("Marks:", marks)

	// Removing an element

	// Syntax: delete(mapName, key)

	delete(marks, "Vk")
	fmt.Println("Marks after deletion:", marks)

	// Maps can store values of any type (using `interface{}` or `any` type for values)

	newMap := map[string]any{
		"Name":  "John",
		"Age":   30,
		"City":  "New York",
		"Grade": 95.5,
	}
	fmt.Println("NewMap:", newMap)

	// Checking for a specific key in the map

	// Syntax: value, exists := mapName[key]

	name, ok := newMap["Name"]
	if ok {
		fmt.Println("Name:", name)
	} else {
		fmt.Println("Name not found")
	}

	// Checking a non-existent key
	address, ok := newMap["Address"]
	if !ok {
		fmt.Println("Address not found")
	}
	fmt.Println("Address:", address)

	// Demonstrating that maps are reference types

	originalMap := map[string]int{
		"Benz":  10000,
		"Volvo": 20000,
		"Tata":  30000,
	}

	// Assigning originalMap to duplicateMap (both reference the same underlying map)
	duplicateMap := originalMap

	fmt.Println("Original map before changes:", originalMap)
	fmt.Println("Duplicate map before changes:", duplicateMap)

	// Modifying the duplicate map
	duplicateMap["Benz"] = 15000
	duplicateMap["Toyota"] = 40000

	fmt.Println("Original map after changes:", originalMap)
	fmt.Println("Duplicate map after changes:", duplicateMap)
}
