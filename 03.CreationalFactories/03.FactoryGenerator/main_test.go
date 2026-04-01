package main

import "testing"

func TestNewEmployeeFactory_Functional(t *testing.T) {
	devFactory := NewEmployeeFactory("Developer", 60000)
	
	e := devFactory("Alice")
	
	if e.Name != "Alice" {
		t.Errorf("Expected name 'Alice', got '%s'", e.Name)
	}
	if e.Position != "Developer" {
		t.Errorf("Expected position 'Developer', got '%s'", e.Position)
	}
	if e.AnnualIncome != 60000 {
		t.Errorf("Expected income 60000, got %d", e.AnnualIncome)
	}
}

func TestEmployeeFactory_Structural(t *testing.T) {
	managerFactory := NewEmployeeFactory2("Manager", 80000)
	
	e := managerFactory.Create("Bob")
	
	if e.Name != "Bob" {
		t.Errorf("Expected name 'Bob', got '%s'", e.Name)
	}
	if e.Position != "Manager" {
		t.Errorf("Expected position 'Manager', got '%s'", e.Position)
	}
	if e.AnnualIncome != 80000 {
		t.Errorf("Expected income 80000, got %d", e.AnnualIncome)
	}

	// Test modification
	managerFactory.AnnualIncome = 90000
	e2 := managerFactory.Create("Charlie")
	if e2.AnnualIncome != 90000 {
		t.Errorf("Expected income 90000 after update, got %d", e2.AnnualIncome)
	}
}
