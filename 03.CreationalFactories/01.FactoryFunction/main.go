package main

import "fmt"

// Person represents a person with name and age.
type Person struct {
	Name string
	Age  int
}

// NewPerson is a factory function that creates a new Person.
// It abstracts the creation logic, allowing for validation or default values if needed.
func NewPerson(name string, age int) *Person {
	return &Person{name, age}
}

func main() {
	// Initialize directly
	p := Person{"John", 22}
	fmt.Println(p)

	// Use a constructor/factory function
	p2 := NewPerson("Jane", 21)
	p2.Age = 30
	fmt.Println(p2)
}
