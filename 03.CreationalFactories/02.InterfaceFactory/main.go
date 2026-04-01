package main

import "fmt"

// Person is the interface that our factory will return.
// Using an interface allows us to encapsulate the implementation details.
type Person interface {
	SayHello()
}

// person is the concrete implementation of the Person interface.
// It is unexported (lower case) to force usage of the factory function.
type person struct {
	name string
	age  int
}

// SayHello implements the Person interface.
func (p *person) SayHello() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.name, p.age)
}

// NewPerson is the factory function that returns a Person interface.
// This decouples the client code from the specific struct implementation.
func NewPerson(name string, age int) Person {
	return &person{
		name: name,
		age:  age,
	}
}

func main() {
	// We can't create 'person' struct directly here because it is private (if we were in a different package),
	// but within the same package main it works. Ideally factories are in their own packages.
	// For this example, we demonstrate using the factory.
	p := NewPerson("Sam", 25)
	p.SayHello()
}
