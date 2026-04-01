package main

import "testing"

func TestNewPerson(t *testing.T) {
	name := "Alice"
	age := 30
	p := NewPerson(name, age)

	if p.Name != name {
		t.Errorf("Expected name %s, got %s", name, p.Name)
	}
	if p.Age != age {
		t.Errorf("Expected age %d, got %d", age, p.Age)
	}
}
