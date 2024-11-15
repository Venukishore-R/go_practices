package main

import "fmt"

// LOOPING USING RANGE
// The 'range' keyword is used to loop through arrays, maps, slices, and strings.

func rangeExample() {
	fmt.Print("\n\nLOOPS USING RANGE\n")

	// Basic FOR EACH loop (iterating through a slice)
	cars := []string{"Benz", "Audi", "Volvo", "Tata", "Tesla"}

	// Looping through the 'cars' slice. 'index' is the index and 'car' is the value at that index.
	for index, car := range cars {
		fmt.Println("Number of car:", index, "Car name:", car)
	}

	// String iteration using runes or bytes
	// In Go, strings are essentially sequences of bytes. We can convert them to byte slices for iteration.
	myName := "Venukishore"

	// Convert the string to a byte slice for iteration
	myNameInBytes := []byte(myName)

	// Iterate through each byte value of the string
	for i, ch := range myNameInBytes {
		// 'i' is the index, 'ch' is the byte value (ASCII) of the character
		// Using %U to show the Unicode value of the character
		fmt.Printf("Character index: %d Character in bytes (ASCII): %v Character in string: %s, Unicode: %U\n", i, ch, string(ch), rune(ch))
	}

	// Map iteration (iterating over a map's key-value pairs)
	// Maps are unordered collections of key-value pairs.
	studentMarks := map[string]int{
		"Venu":    100,
		"Kishore": 96,
	}

	// Looping through the map. 'key' is the map key and 'value' is the value associated with the key.
	for key, value := range studentMarks {
		fmt.Println("Student:", key, "Marks:", value)
	}

	// Channel iteration
	// Channels are used to communicate between goroutines. We can iterate over values received from a channel.
	numChannel := make(chan int)

	// Goroutine sending values to the channel
	go func() {
		numChannel <- 10
		numChannel <- 20
		close(numChannel) // Close the channel to indicate no more values will be sent
	}()

	// Iterating over the channel values
	// The loop will continue until the channel is closed and all values are received
	for num := range numChannel {
		fmt.Println("Channel value:", num)
	}
}
