package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Print - used to print in the command line (console)
	fmt.Print("used to print in the command line (console)\n")
	fmt.Print("Hello world\n")

	// Printf - used to print formatted strings in the command line
	printf := "used to print the output with dynamic variable value in the command line"
	someNumber := 1
	fmt.Printf("Printf is %s\n", printf)
	fmt.Printf("Output with formatted number: %d\n", someNumber)

	// Println - prints the output with a newline
	fmt.Println("printed as a new line")
	fmt.Println("Next line is printed in a new line")

	// Sprint - similar to Print but returns the result as a string
	var newVar1 = fmt.Sprint("Returned to a new variable")
	fmt.Println(newVar1)

	// Sprintf - similar to Printf but returns the result as a formatted string
	formattedString := "SOME STRING"
	var newVar2 = fmt.Sprintf("Returned to a new variable with formatted string: %s", formattedString)
	fmt.Println(newVar2)

	// Sprintln - similar to Println but returns the result as a string
	var newVar3 = fmt.Sprintln("Printed")
	fmt.Print(newVar3)
	fmt.Println("Next line")

	// Ensure the "print" directory exists
	err := os.MkdirAll("print", os.ModePerm)
	if err != nil {
		log.Fatalf("Unable to create directory: %v", err)
	}

	// Open the file for writing, create it if it doesn't exist, and truncate it if it does
	f, err := os.OpenFile("print/test.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("Unable to open file: %v", err)
	}
	defer f.Close()

	// Fprint - write to a file
	nBytes, err := fmt.Fprint(f, "Hello world\n")
	if err != nil {
		log.Fatalf("Unable to write to file: %v", err)
	}
	fmt.Println("Number of bytes written:", nBytes)

	// Fprintln - write a line to a file with a newline
	_, err = fmt.Fprintln(f, "Next line will surely start in a new line")
	if err != nil {
		log.Fatalf("Unable to write to file: %v", err)
	}

	// Fprintf - write formatted output to a file
	myName := "VENUKISHORE R"
	_, err = fmt.Fprintf(f, "Formatted output: %s\n", myName)
	if err != nil {
		log.Fatalf("Unable to write to file: %v", err)
	}
}
