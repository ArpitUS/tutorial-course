package main

import "fmt"

func main() {
	// Arrays
	var numbers [3]int = [3]int{10, 20, 30}
	fmt.Println("Array:", numbers)

	// Slices
	fruits := []string{"apple", "banana", "cherry"}
	fruits = append(fruits, "mango")
	fmt.Println("Slice:", fruits)

	// Maps
	capitals := map[string]string{
		"India":  "New Delhi",
		"France": "Paris",
		"Japan":  "Tokyo",
	}
	capitals["Germany"] = "Berlin"

	fmt.Println("Map of capitals:")
	for country, capital := range capitals {
		fmt.Printf("%s: %s\n", country, capital)
	}
}
