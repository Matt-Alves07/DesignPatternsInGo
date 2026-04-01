package main

import "testing"

func TestPersonBuilder_Functional(t *testing.T) {
	b := PersonBuilder{}
	p := b.Called("John").WorksAsA("Engineer").Build()

	if p.name != "John" {
		t.Errorf("Expected name 'John', got '%s'", p.name)
	}
	if p.position != "Engineer" {
		t.Errorf("Expected position 'Engineer', got '%s'", p.position)
	}
}

func TestPersonBuilder_SequentialOrder(t *testing.T) {
	// Verify that actions are applied in order
	b := PersonBuilder{}
	p := b.Called("Initial").Called("Final").Build()

	if p.name != "Final" {
		t.Errorf("Expected name 'Final', got '%s'", p.name)
	}
}
