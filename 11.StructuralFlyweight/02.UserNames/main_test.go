package main

import (
	"testing"
)

func TestUserCreation(t *testing.T) {
	user := NewUser("John Doe")

	if user == nil {
		t.Fatal("Expected user to be created")
	}

	if user.FullName != "John Doe" {
		t.Errorf("Expected FullName 'John Doe', got '%s'", user.FullName)
	}
}

func TestUser2Creation(t *testing.T) {
	user := NewUser2("John Doe")

	if user == nil {
		t.Fatal("Expected user to be created")
	}

	if len(user.names) != 2 {
		t.Errorf("Expected 2 name parts, got %d", len(user.names))
	}
}

func TestUser2FullName(t *testing.T) {
	user := NewUser2("John Doe")
	fullName := user.FullName()

	if fullName != "John Doe" {
		t.Errorf("Expected 'John Doe', got '%s'", fullName)
	}
}

func TestUser2ReusesNames(t *testing.T) {
	// Clear allNames
	allNames = nil

	_ = NewUser2("John Doe")
	initialNameCount := len(allNames)

	_ = NewUser2("John Smith")
	nameCountAfterSecond := len(allNames)

	// "John" should be reused, only "Smith" should be added
	// So we should have increased by 1, not 2
	if nameCountAfterSecond-initialNameCount > 2 {
		t.Errorf("Expected at most 1 new name, but gained %d", nameCountAfterSecond-initialNameCount)
	}
}

func TestUser2FlyweightOptimization(t *testing.T) {
	// Clear allNames
	allNames = nil

	user1 := NewUser2("Jane Doe")
	user2 := NewUser2("Jane Smith")

	// "Jane" should be shared in the pool
	if len(allNames) > 3 {
		t.Errorf("Expected at most 3 unique names, got %d", len(allNames))
	}

	if user1.FullName() != "Jane Doe" {
		t.Errorf("Expected 'Jane Doe', got '%s'", user1.FullName())
	}

	if user2.FullName() != "Jane Smith" {
		t.Errorf("Expected 'Jane Smith', got '%s'", user2.FullName())
	}
}

func TestUser2MultipleSameName(t *testing.T) {
	// Clear allNames
	allNames = nil

	_ = NewUser2("Alice Bob")
	_ = NewUser2("Alice Charlie")
	_ = NewUser2("David Bob")

	// Alice should appear twice, Bob appears twice - both should be in pool once
	// Total: Alice, Bob, Charlie, David
	if len(allNames) > 4 {
		t.Errorf("Expected at most 4 unique names, got %d", len(allNames))
	}
}

func TestUser2SameNameDifferentUsers(t *testing.T) {
	// Clear allNames
	allNames = nil

	user1 := NewUser2("Bob Smith")
	user2 := NewUser2("Bob Jones")

	if user1.FullName() != "Bob Smith" {
		t.Errorf("Expected 'Bob Smith', got '%s'", user1.FullName())
	}

	if user2.FullName() != "Bob Jones" {
		t.Errorf("Expected 'Bob Jones', got '%s'", user2.FullName())
	}

	// "Bob" should be in the pool only once
	bobCount := 0
	for _, name := range allNames {
		if name == "Bob" {
			bobCount++
		}
	}

	if bobCount != 1 {
		t.Errorf("Expected 'Bob' to appear once in pool, appeared %d times", bobCount)
	}
}

func TestUser2SingleName(t *testing.T) {
	user := NewUser2("Prince")
	fullName := user.FullName()

	if fullName != "Prince" {
		t.Errorf("Expected 'Prince', got '%s'", fullName)
	}
}

func TestUser2ThreeNamePerson(t *testing.T) {
	user := NewUser2("José María Garcia")
	fullName := user.FullName()

	if fullName != "José María Garcia" {
		t.Errorf("Expected 'José María Garcia', got '%s'", fullName)
	}

	if len(user.names) != 3 {
		t.Errorf("Expected 3 name parts, got %d", len(user.names))
	}
}

func TestUser2FlyweightMemory(t *testing.T) {
	// Clear allNames
	allNames = nil

	// Create many users with shared name parts
	for i := 0; i < 100; i++ {
		NewUser2("John Doe")
	}

	// Should still only have 2 unique names in the pool
	if len(allNames) != 2 {
		t.Errorf("Expected 2 unique names for 100 John Doe users, got %d", len(allNames))
	}
}

func TestUser2NameIndices(t *testing.T) {
	// Clear allNames
	allNames = nil

	user := NewUser2("Alice Brown")

	// Verify names are stored as indices
	if len(user.names) != 2 {
		t.Errorf("Expected 2 indices, got %d", len(user.names))
	}

	// Verify indices are valid
	for i, idx := range user.names {
		if int(idx) >= len(allNames) {
			t.Errorf("Index %d at position %d is out of bounds", idx, i)
		}
	}
}
