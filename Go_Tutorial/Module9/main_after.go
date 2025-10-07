package main

import (
	"errors"
	"fmt"
)

// divide returns the result of a/b or an error if b is zero
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// Example of an error case
	_, err = divide(5, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
