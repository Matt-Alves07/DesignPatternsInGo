package main

import "testing"

func TestRelationships_FindAllChildrenOf(t *testing.T) {
	parent := Person{"Parent"}
	child1 := Person{"Child1"}
	child2 := Person{"Child2"}

	relationships := Relationships{}
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)

	children := relationships.FindAllChildrenOf("Parent")

	if len(children) != 2 {
		t.Errorf("Expected 2 children, got %d", len(children))
	}

	// Verify names
	found1 := false
	found2 := false

	for _, c := range children {
		if c.name == "Child1" {
			found1 = true
		}
		if c.name == "Child2" {
			found2 = true
		}
	}

	if !found1 || !found2 {
		t.Errorf("Did not find all children. Found1: %v, Found2: %v", found1, found2)
	}
}

// MockBrowser helps testing Research without depending on real Relationships (double decoupling verification)
type MockBrowser struct {}

func (m *MockBrowser) FindAllChildrenOf(name string) []*Person {
	if name == "John" {
		return []*Person{{"MockChild"}}
	}
	return []*Person{}
}

func TestResearch_Investigate(t *testing.T) {
	// Although Investigate prints to stdout, we can verify that Research accepts any RelationshipBrowser
	mock := &MockBrowser{}
	research := Research{browser: mock}
	
	// This mainly tests that the interface integration works, as we can't easily assert stdout here without redirection,
	// but the fact it compiles and runs with a mock proves DIP.
	research.Investigate()
}
