package main

import "fmt"

// Function to calculate the square of a number
func square(number float64, squareResult chan float64) {
	result := number * number
	squareResult <- result
}

// Function to calculate the cube of a number
func cube(number float64, cubeResult chan float64) {
	result := number * number * number
	cubeResult <- result
}

// Function to calculate and print square and cube of a number
func calculateResult() {
	fmt.Println("Calculating square and cube of a number...")

	// Declaring buffered channels
	squareChan := make(chan float64, 1)
	cubeChan := make(chan float64, 1)

	// Launching Goroutines
	go square(1.5, squareChan)
	go cube(1.5, cubeChan)

	// Receiving results from channels
	squareResult := <-squareChan
	cubeResult := <-cubeChan

	// Printing results
	fmt.Println("Square:", squareResult)
	fmt.Println("Cube:  ", cubeResult)
}
