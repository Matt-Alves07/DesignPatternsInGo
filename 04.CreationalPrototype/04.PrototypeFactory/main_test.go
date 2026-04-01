package main

import (
	"testing"
)

func TestEmployeeDeepCopy(t *testing.T) {
	original := &Employee{"John",
		Address{100, "123 East Dr", "London"}}

	copied := original.DeepCopy()

	// Verify content is the same
	if copied.Name != original.Name {
		t.Errorf("Expected Name %s, got %s", original.Name, copied.Name)
	}
	if copied.Office.Suite != original.Office.Suite {
		t.Errorf("Expected Suite %d, got %d", original.Office.Suite, copied.Office.Suite)
	}
	if copied.Office.StreetAddress != original.Office.StreetAddress {
		t.Errorf("Expected StreetAddress %s, got %s", original.Office.StreetAddress, copied.Office.StreetAddress)
	}
	if copied.Office.City != original.Office.City {
		t.Errorf("Expected City %s, got %s", original.Office.City, copied.Office.City)
	}
}

func TestMainOfficeEmployeeFactory(t *testing.T) {
	emp1 := NewMainOfficeEmployee("John", 100)
	emp2 := NewMainOfficeEmployee("Jane", 200)

	// Verify both have main office address
	if emp1.Office.StreetAddress != "123 East Dr" {
		t.Errorf("Expected emp1 office 'East Dr', got %s", emp1.Office.StreetAddress)
	}
	if emp2.Office.StreetAddress != "123 East Dr" {
		t.Errorf("Expected emp2 office 'East Dr', got %s", emp2.Office.StreetAddress)
	}

	// Verify different suites
	if emp1.Office.Suite != 100 {
		t.Errorf("Expected emp1 Suite 100, got %d", emp1.Office.Suite)
	}
	if emp2.Office.Suite != 200 {
		t.Errorf("Expected emp2 Suite 200, got %d", emp2.Office.Suite)
	}

	// Verify names
	if emp1.Name != "John" {
		t.Errorf("Expected emp1.Name 'John', got %s", emp1.Name)
	}
	if emp2.Name != "Jane" {
		t.Errorf("Expected emp2.Name 'Jane', got %s", emp2.Name)
	}
}

func TestAuxOfficeEmployeeFactory(t *testing.T) {
	emp1 := NewAuxOfficeEmployee("Alice", 150)
	emp2 := NewAuxOfficeEmployee("Bob", 250)

	// Verify both have aux office address
	if emp1.Office.StreetAddress != "66 West Dr" {
		t.Errorf("Expected emp1 office 'West Dr', got %s", emp1.Office.StreetAddress)
	}
	if emp2.Office.StreetAddress != "66 West Dr" {
		t.Errorf("Expected emp2 office 'West Dr', got %s", emp2.Office.StreetAddress)
	}

	// Verify different suites
	if emp1.Office.Suite != 150 {
		t.Errorf("Expected emp1 Suite 150, got %d", emp1.Office.Suite)
	}
	if emp2.Office.Suite != 250 {
		t.Errorf("Expected emp2 Suite 250, got %d", emp2.Office.Suite)
	}

	// Verify names
	if emp1.Name != "Alice" {
		t.Errorf("Expected emp1.Name 'Alice', got %s", emp1.Name)
	}
	if emp2.Name != "Bob" {
		t.Errorf("Expected emp2.Name 'Bob', got %s", emp2.Name)
	}
}

func TestMainAndAuxOfficeAreIndependent(t *testing.T) {
	mainEmp := NewMainOfficeEmployee("John", 100)
	auxEmp := NewAuxOfficeEmployee("Jane", 100)

	// Verify they have different addresses
	if mainEmp.Office.StreetAddress == auxEmp.Office.StreetAddress {
		t.Error("Expected main and aux office employees to have different addresses")
	}

	// Verify both have same suite number but different streets
	if mainEmp.Office.Suite != 100 || auxEmp.Office.Suite != 100 {
		t.Error("Expected both employees to have suite 100")
	}
	if mainEmp.Office.StreetAddress != "123 East Dr" {
		t.Errorf("Expected main office to be 'East Dr', got %s", mainEmp.Office.StreetAddress)
	}
	if auxEmp.Office.StreetAddress != "66 West Dr" {
		t.Errorf("Expected aux office to be 'West Dr', got %s", auxEmp.Office.StreetAddress)
	}
}

func TestEmployeeFactoriesProduceIndependentCopies(t *testing.T) {
	emp1 := NewMainOfficeEmployee("Employee1", 100)
	emp2 := NewMainOfficeEmployee("Employee2", 100)

	// Modify emp1's office details (shouldn't affect emp2)
	emp1.Office.Suite = 999

	// Verify emp2 is unchanged
	if emp2.Office.Suite != 100 {
		t.Errorf("Expected emp2.Office.Suite to be 100, got %d", emp2.Office.Suite)
	}

	// Verify they are independent
	if &emp1.Office == &emp2.Office {
		t.Error("Expected emp1 and emp2 to have independent Office structures")
	}
}

func TestNewEmployeeUtilityFunction(t *testing.T) {
	proto := &Employee{"Proto",
		Address{0, "999 Test St", "TestCity"}}

	emp1 := newEmployee(proto, "John", 100)
	emp2 := newEmployee(proto, "Jane", 200)

	// Verify all are independent copies
	if emp1.Name != "John" || emp1.Office.Suite != 100 {
		t.Error("emp1 not created correctly")
	}
	if emp2.Name != "Jane" || emp2.Office.Suite != 200 {
		t.Error("emp2 not created correctly")
	}

	// Original proto should remain unchanged
	if proto.Name != "Proto" || proto.Office.Suite != 0 {
		t.Error("Original proto was modified")
	}
}
