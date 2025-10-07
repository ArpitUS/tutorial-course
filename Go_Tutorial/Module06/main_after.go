package main

import "fmt"

// Person defines a simple struct
type Person struct {
	Name string
	Age  int
	City string
}

// Greet is a method that belongs to the Person type
func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old from %s.\n", p.Name, p.Age, p.City)
}

func main() {
	// Creating a struct instance
	person1 := Person{Name: "Arpit", Age: 29, City: "Lucknow"}

	// Calling the method
	person1.Greet()
}
