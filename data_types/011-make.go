// MAKE FUNCTION

// The `make()` function in Go is used to create or initialize slices, maps, and channels.
// It allocates memory and sets up the internal data structures needed for these types.

// Note: The `make()` function is only used for slices, maps, and channels.
// For other types (like arrays, structs), you can use the `new()` function or declare them directly.

package main

import "fmt"

func makeExample() {
	fmt.Println("\n\nMAKE FUNCTION EXAMPLES")

	// 1. Creating a map using make()

	// Syntax: mapName := make(map[keyType]valueType)

	myMap := make(map[string]interface{})
	fmt.Println("My map:", myMap)

	// Adding elements to the map
	myMap["Name"] = "Alice"
	myMap["Age"] = 25
	fmt.Println("My map after adding elements:", myMap)

	// 2. Creating a slice using make()

	// Syntax: sliceName := make([]dataType, length, capacity)
	// Length: Initial number of elements in the slice
	// Capacity: Maximum number of elements the slice can hold before resizing

	mySlice := make([]int, 5, 10)
	fmt.Println("My slice:", mySlice)
	fmt.Println("Length of my slice:", len(mySlice))
	fmt.Println("Capacity of my slice:", cap(mySlice))

	// Updating elements in the slice
	mySlice[0] = 100
	fmt.Println("My slice after updating:", mySlice)

	// Appending elements to the slice
	mySlice = append(mySlice, 200, 300)
	fmt.Println("My slice after appending:", mySlice)
	fmt.Println("Length of my slice after appending:", len(mySlice))
	fmt.Println("Capacity of my slice after appending:", cap(mySlice))

	// 3. Creating a channel using make()

	// Syntax: channelName := make(chan dataType, bufferSize)
	// BufferSize: Specifies the capacity of the channel (0 means unbuffered)

	myChannel := make(chan int, 3)
	fmt.Println("My channel:", myChannel)

	// Sending data to the channel (non-blocking because it has a buffer)
	myChannel <- 1
	myChannel <- 2
	myChannel <- 3

	// Receiving data from the channel
	fmt.Println("Received from channel:", <-myChannel)
	fmt.Println("Received from channel:", <-myChannel)
	fmt.Println("Received from channel:", <-myChannel)

	// Closing the channel
	close(myChannel)
}
