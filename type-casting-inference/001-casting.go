// TYPE CASTING
//
// Type casting is used to change the data type of a variable to another data type.
//
// Basic syntax:
// result := dataType(originalValue)

package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func typeCastingExample() {
	// FLOAT TO INT
	var number float64 = 10.5
	var integer int = int(number) // Casting float64 to int (decimal part is truncated)

	fmt.Println("Original number:", number)
	fmt.Println("Converted integer:", integer)

	// INT TO STRING (Corrected)
	var mark int = 90
	var markInString = strconv.Itoa(mark) // Use strconv.Itoa() for integer to string conversion

	fmt.Println("Original mark:", mark, "type:", reflect.TypeOf(mark))
	fmt.Println("Converted mark:", markInString, "type:", reflect.TypeOf(markInString))

	// STRING TO INT
	var ageInString string = "21"
	ageInInt, err := strconv.Atoi(ageInString) // Converts string to int
	if err != nil {
		fmt.Println("Error converting string to int:", err)
	} else {
		fmt.Println("Original age:", ageInString, "type:", reflect.TypeOf(ageInString))
		fmt.Println("Converted age:", ageInInt, "type:", reflect.TypeOf(ageInInt))
	}

	// INT TO FLOAT
	result := float64(mark) // Casting int to float64
	fmt.Println("Converted float result:", result)

	// FLOAT TO INT
	result2 := int(number) // Casting float64 to int
	fmt.Println("Converted int result2:", result2)

	// STRING TO BYTE SLICE
	resultInByte := []byte("Some string") // Converts string to a byte slice
	fmt.Println("Converted byte slice:", resultInByte)

	// BYTE SLICE TO STRING
	byteToString := string(resultInByte) // Converts byte slice back to string
	fmt.Println("Converted back to string:", byteToString)
}