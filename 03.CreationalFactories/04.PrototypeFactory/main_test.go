package main

import "testing"

func TestNewEmployee(t *testing.T) {
	// Test Developer
	dev := NewEmployee(Developer)
	if dev.Position != "Developer" {
		t.Errorf("Expected Developer position, got %s", dev.Position)
	}
	if dev.AnnualIncome != 60000 {
		t.Errorf("Expected 60000 income, got %d", dev.AnnualIncome)
	}

	// Test Manager
	mgr := NewEmployee(Manager)
	if mgr.Position != "Manager" {
		t.Errorf("Expected Manager position, got %s", mgr.Position)
	}
	if mgr.AnnualIncome != 80000 {
		t.Errorf("Expected 80000 income, got %d", mgr.AnnualIncome)
	}
}

func TestNewEmployee_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for unsupported role")
		}
	}()

	NewEmployee(999)
}
