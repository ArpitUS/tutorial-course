package main

import "fmt"

// Speaker defines an interface
type Speaker interface {
	Speak()
}

// Human implements the Speaker interface
type Human struct {
	Name string
}

func (h Human) Speak() {
	fmt.Println(h.Name, "says hello!")
}

// Robot implements the Speaker interface
type Robot struct {
	ID string
}

func (r Robot) Speak() {
	fmt.Println("Robot", r.ID, "beeps in greeting.")
}

func main() {
	var s Speaker

	s = Human{Name: "Arpit"}
	s.Speak()

	s = Robot{ID: "RX-78"}
	s.Speak()
}
