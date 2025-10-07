package main

import "fmt"

func main() {
	// Declaring variables
	var name string = "Arpit"
	var age int = 29
	city := "Lucknow" // Type inferred automatically

	// Declaring constants
	const country = "India"

	// Displaying values
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("City:", city)
	fmt.Println("Country:", country)

	// Using formatted output
	height := 5.9
	fmt.Printf("%s is %d years old, %.1f feet tall, and lives in %s, %s.\n",
		name, age, height, city, country)
}
