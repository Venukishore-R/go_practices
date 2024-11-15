// STRUCTS IN GO

// Structs are user-defined data types that allow grouping of different fields.
// It is useful for creating complex data structures by combining different types.

package main

import "fmt"

// Declaring a struct type using `type` keyword
type Person struct {
	name   string
	age    int
	marks  int
	gender string
}

// Embedded Structs Example
type Address struct {
	city    string
	state   string
	pincode int
}

// Employee struct with embedded Person and Address structs
type Employee struct {
	Person  // Embedded Person struct
	Address // Embedded Address struct
	salary  float64
}

// Method with value receiver
func (p Person) display() {
	fmt.Printf("Display Method - Name: %s, Age: %d, Marks: %d, Gender: %s\n", p.name, p.age, p.marks, p.gender)

	// Changes made here do not affect the original struct because it's passed by value.
	p.name = "Changed Name"
	fmt.Println("Inside display method:", p.name)
}

// Method with pointer receiver
func (p *Person) show() {
	fmt.Printf("Show Method - Name: %s, Age: %d, Marks: %d, Gender: %s\n", p.name, p.age, p.marks, p.gender)

	// Changes made here affect the original struct because it's passed by reference.
	p.name = "Updated Name"
	fmt.Println("Inside show method:", p.name)
}

func structExample() {
	fmt.Println("\n\nSTRUCTS EXAMPLE")

	// 1. Initializing a struct using field names
	person1 := Person{
		name:   "Alice",
		age:    20,
		marks:  85,
		gender: "Female",
	}

	// Accessing struct fields
	fmt.Println("Person 1:", person1)
	fmt.Println("Name:", person1.name)
	fmt.Println("Marks:", person1.marks)

	// Calling a method with value receiver
	person1.display()
	fmt.Println("Person 1 after display method:", person1)

	// 2. Array of structs
	persons := []Person{
		{name: "Bob", age: 21, marks: 90, gender: "Male"},
		{name: "Charlie", age: 22, marks: 75, gender: "Male"},
	}

	// Iterating over an array of structs
	for _, person := range persons {
		person.display()
	}

	// 3. Anonymous struct
	car := struct {
		name  string
		model string
	}{
		name:  "Tesla",
		model: "Model 3",
	}

	fmt.Println("\nAnonymous Struct:")
	fmt.Println("Car Name:", car.name)
	fmt.Println("Car Model:", car.model)

	// 4. Embedded structs
	employee := Employee{
		Person:  Person{name: "John", age: 30, marks: 88, gender: "Male"},
		Address: Address{city: "New York", state: "NY", pincode: 10001},
		salary:  75000.50,
	}

	fmt.Println("\nEmployee Details:")
	employee.display()
	fmt.Printf("City: %s, State: %s, Pincode: %d, Salary: $%.2f\n",
		employee.city, employee.state, employee.pincode, employee.salary)

	// 5. Pointer receiver method
	person2 := Person{
		name:   "Vk",
		age:    25,
		marks:  80,
		gender: "Male",
	}

	person2.show()
	fmt.Println("Person 2 after show method:", person2)
}
