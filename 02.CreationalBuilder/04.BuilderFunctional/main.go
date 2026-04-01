package main

import "fmt"

// Person represents a person with a name and a job position.
type Person struct {
	name, position string
}

// personMod is a function type that modifies a Person.
type personMod func(*Person)

// PersonBuilder builds a Person by applying a list of modifications.
type PersonBuilder struct {
	actions []personMod
}

// Called adds a modification to set the person's name.
func (b *PersonBuilder) Called(name string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.name = name
	})
	return b
}

// Build creates a Person by applying all accumulated modifications.
func (b *PersonBuilder) Build() *Person {
	p := Person{}
	for _, a := range b.actions {
		a(&p)
	}
	return &p
}

// WorksAsA adds a modification to set the person's position.
// This demonstrates adding extensions to the builder.
func (b *PersonBuilder) WorksAsA(position string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.position = position
	})
	return b
}

func main() {
	b := PersonBuilder{}
	p := b.Called("Dmitri").WorksAsA("dev").Build()
	fmt.Println(*p)
}