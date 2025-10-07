package main

import "fmt"

func main() {
	// IF-ELSE example
	age := 20
	if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a minor.")
	}

	// SWITCH example
	day := "Saturday"
	switch day {
	case "Saturday", "Sunday":
		fmt.Println("It’s the weekend!")
	default:
		fmt.Println("It’s a weekday.")
	}

	// FOR loop example
	fmt.Println("Counting from 1 to 5:")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// RANGE loop example
	fruits := []string{"apple", "banana", "mango"}
	fmt.Println("Favorite fruits:")
	for index, fruit := range fruits {
		fmt.Printf("%d: %s\n", index+1, fruit)
	}
}
