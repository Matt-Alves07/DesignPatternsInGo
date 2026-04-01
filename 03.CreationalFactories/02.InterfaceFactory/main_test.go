package main

import (
	"testing"
)

func TestNewPersonFactory(t *testing.T) {
	name := "Test User"
	age := 40

	// Test Factory returns the interface
	var p Person = NewPerson(name, age)

	if p == nil {
		t.Fatal("Factory returned nil")
	}

	// Verify concrete type details (possible since we are in the same package)
	concreteP, ok := p.(*person)
	if !ok {
		t.Fatal("Factory did not return *person type")
	}

	if concreteP.name != name {
		t.Errorf("Expected name '%s', got '%s'", name, concreteP.name)
	}
	if concreteP.age != age {
		t.Errorf("Expected age %d, got %d", age, concreteP.age)
	}
}
