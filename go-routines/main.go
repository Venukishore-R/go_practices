// GO ROUTINES

// GO routines are used to execute functions or set of codes concurrently(at the same time) while other codes being executed

package main

import (
	"fmt"
	"time"
)

func doSomething() {
	fmt.Println("Hello, I'm a Go routine")
}

func printSomething() {
	for i := 0; i < 10; i++ {
		fmt.Println("Printed some thing: ", i)
	}
}

func main() {
	fmt.Print("\nGO ROUTINES\n\n")

	// called first
	go printSomething()

	doSomething()

	go func(msg string) {
		// fmt.Println(msg)
		for i := 0; i < 10; i++ {
			fmt.Println(msg)
		}
	}("Hello world")

	// need to wait for few seconds to get the output from go routines
	time.Sleep(time.Second)
	fmt.Println("Done")
}
