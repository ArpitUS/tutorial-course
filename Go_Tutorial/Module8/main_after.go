package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Number:", i)
		time.Sleep(200 * time.Millisecond)
	}
}

func printLetters() {
	for _, letter := range []string{"A", "B", "C", "D", "E"} {
		fmt.Println("Letter:", letter)
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	go printNumbers()
	go printLetters()

	// Using a channel to block main until completion
	time.Sleep(2 * time.Second)
	fmt.Println("Main function finished.")
}
