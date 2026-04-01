package main

import "fmt"

// DIP (Dependency Inversion Principle): High-level modules should not depend on low-level modules.
// Both should depend on abstractions.
// Abstractions should not depend on details. Details should depend on abstractions.

// Relationship defines the type of relationship between people.
type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

// Person represents a person with a name.
type Person struct {
	name string
}

// Info represents a single relationship entry.
type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

// RelationshipBrowser is an abstraction that allows researching relationships.
// High-level modules will depend on this interface.
type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

// Relationships is a low-level module that stores relationships.
// It implements RelationshipBrowser.
type Relationships struct {
	relations []Info
}

// FindAllChildrenOf finds all children of a given person.
func (rs *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)
	for i, v := range rs.relations {
		if v.relationship == Parent && v.from.name == name {
			result = append(result, rs.relations[i].to)
		}
	}
	return result
}

// AddParentAndChild adds a parent-child relationship.
func (rs *Relationships) AddParentAndChild(parent, child *Person) {
	rs.relations = append(rs.relations, Info{parent, Parent, child})
	rs.relations = append(rs.relations, Info{child, Child, parent})
}

// Research is a high-level module that performs operations on relationships.
// It depends on the RelationshipBrowser interface, not the concrete Relationships struct.
type Research struct {
	browser RelationshipBrowser
}

// Investigate performs the research operation.
func (r *Research) Investigate() {
	for _, p := range r.browser.FindAllChildrenOf("John") {
		fmt.Println("John has a child called", p.name)
	}
}

func main() {
	parent := Person{"John"}
	child1 := Person{"Chris"}
	child2 := Person{"Matt"}

	// low-level module
	relationships := Relationships{}
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)

	// high-level module
	// Research now depends on the RelationshipBrowser abstraction
	research := Research{&relationships}
	research.Investigate()
}