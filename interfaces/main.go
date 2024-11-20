// INTERFACES

// Interfaces in Go are collections of method signatures. Any type that implements
// all the methods of an interface is said to satisfy the interface.

package main

import "fmt"

// Declaring an interface named Area with method signatures
type Area interface {
	area()
}

// Circle struct
type Circle struct {
	radius float64
}

// Rectangle struct
type Rectangle struct {
	length float64
	width  float64
}

// Square struct
type Square struct {
	side float64
}

func (s *Square) area() {
	fmt.Printf("Area of the square: %.2f\n", s.side*s.side)
}

func (c *Circle) area() {
	fmt.Printf("Area of the circle: %.2f\n", 3.14*c.radius*c.radius)
}

func (r *Rectangle) area() {
	fmt.Printf("Area of the rectangle: %.2f\n", r.length*r.width)
}

// Main function
func main() {
	// Using interface type Area
	var shape Area

	// Assigning Circle to the interface and calling the method
	shape = &Circle{radius: 5}
	shape.area()

	// Assigning Rectangle to the interface and calling the method
	shape = &Rectangle{length: 5, width: 10}
	shape.area()

	// Assigning Square to the interface and calling the method
	shape = &Square{side: 7}
	shape.area()
}
