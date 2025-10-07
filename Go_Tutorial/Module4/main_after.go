package main

import "fmt"

// greet prints a greeting message
func greet(name string) {
	fmt.Println("Hello,", name)
}

// add returns the sum of two integers
func add(a int, b int) int {
	return a + b
}

// multiply returns the product of two integers
func multiply(a, b int) int {
	return a * b
}

func main() {
	greet("Arpit")

	sum := add(5, 7)
	product := multiply(3, 4)

	fmt.Println("Sum:", sum)
	fmt.Println("Product:", product)
}
