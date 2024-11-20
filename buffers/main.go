// BUFFERS
// Buffers belong to the bytes package and are used to manipulate byte data.
// While strings allow us to calculate their length, buffers are more efficient 
// for handling large strings and performing multiple write operations.

package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Declare a Buffer to hold string data
	var strBuffer bytes.Buffer
	
	// Write a string to the buffer
	strBuffer.WriteString("Hello world\n")
	
	// Read and print the current contents of the buffer
	fmt.Println(strBuffer.String())
	
	// Format and add another string to the buffer
	fmt.Fprintf(&strBuffer, "Adding %s\n", "another text")
	
	// Write a byte slice to the buffer
	strBuffer.Write([]byte("[Done]"))
	
	// Print the final contents of the buffer
	fmt.Println(strBuffer.String())
}