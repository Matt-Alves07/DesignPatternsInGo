package prototype

import (
	"testing"
)

func TestPersonDeepCopyThroughSerialization(t *testing.T) {
	original := &Person{"John",
		&Address{"123 London Rd", "London", "UK"},
		[]string{"Chris", "Matt", "Sam"}}

	copied := original.DeepCopy()

	// Verify content is the same
	if copied.Name != original.Name {
		t.Errorf("Expected Name %s, got %s", original.Name, copied.Name)
	}
	if copied.Address.StreetAddress != original.Address.StreetAddress {
		t.Error("Address.StreetAddress mismatch")
	}
	if copied.Address.City != original.Address.City {
		t.Error("Address.City mismatch")
	}
	if copied.Address.Country != original.Address.Country {
		t.Error("Address.Country mismatch")
	}

	// Verify Friends slice is copied
	if len(copied.Friends) != len(original.Friends) {
		t.Errorf("Expected Friends length %d, got %d", len(original.Friends), len(copied.Friends))
	}
	for i := range original.Friends {
		if copied.Friends[i] != original.Friends[i] {
			t.Errorf("Friend at index %d mismatch", i)
		}
	}
}

func TestPersonDeepCopyThroughSerializationIsDifferentInstance(t *testing.T) {
	original := &Person{"Jane",
		&Address{"321 Baker St", "London", "UK"},
		[]string{"Angela"}}

	copied := original.DeepCopy()

	// Verify Address is a different instance
	if copied.Address == original.Address {
		t.Error("Expected copied Address to be a different pointer")
	}

	// Modify copy and verify original is unchanged
	copied.Name = "Janet"
	copied.Address.StreetAddress = "999 Abbey Rd"
	copied.Friends = append(copied.Friends, "Brenda")

	if original.Name != "Jane" {
		t.Errorf("Expected original.Name to be 'Jane', got %s", original.Name)
	}
	if original.Address.StreetAddress != "321 Baker St" {
		t.Errorf("Expected original.Address.StreetAddress to be '321 Baker St', got %s", original.Address.StreetAddress)
	}
	if len(original.Friends) != 1 {
		t.Errorf("Expected original.Friends to have 1 friend, got %d", len(original.Friends))
	}
}

func TestMultiplePersonDeepCopiesThroughSerialization(t *testing.T) {
	original := &Person{"Alice",
		&Address{"456 Park Ave", "New York", "USA"},
		[]string{"Bob", "Charlie", "Diana"}}

	copy1 := original.DeepCopy()
	copy2 := original.DeepCopy()

	// Modify copies
	copy1.Name = "Copy1"
	copy1.Address.City = "Boston"
	copy1.Friends = append(copy1.Friends, "Eve")

	copy2.Name = "Copy2"
	copy2.Address.City = "Philadelphia"
	copy2.Friends = append(copy2.Friends, "Frank")

	// Verify original is unchanged
	if original.Name != "Alice" {
		t.Errorf("Expected original.Name to be 'Alice', got %s", original.Name)
	}
	if original.Address.City != "New York" {
		t.Errorf("Expected original.Address.City to be 'New York', got %s", original.Address.City)
	}
	if len(original.Friends) != 3 {
		t.Errorf("Expected original.Friends to have 3 friends, got %d", len(original.Friends))
	}

	// Verify copies are independent
	if copy1.Address.City != "Boston" {
		t.Errorf("Expected copy1.Address.City to be 'Boston', got %s", copy1.Address.City)
	}
	if copy2.Address.City != "Philadelphia" {
		t.Errorf("Expected copy2.Address.City to be 'Philadelphia', got %s", copy2.Address.City)
	}
}

func TestPersonDeepCopyEmptyFriendsList(t *testing.T) {
	original := &Person{"Bob",
		&Address{"789 Oak St", "Chicago", "USA"},
		[]string{}}

	copied := original.DeepCopy()

	// Verify Friends slice is properly copied
	if len(copied.Friends) != 0 {
		t.Errorf("Expected copied.Friends to be empty, got length %d", len(copied.Friends))
	}

	copied.Friends = append(copied.Friends, "NewFriend")

	if len(original.Friends) != 0 {
		t.Errorf("Expected original.Friends to still be empty, got length %d", len(original.Friends))
	}
}

func TestPersonDeepCopyComplexScenario(t *testing.T) {
	original := &Person{"Charlie",
		&Address{"111 Elm St", "Denver", "USA"},
		[]string{"David", "Eve", "Frank", "Grace"}}

	copied := original.DeepCopy()

	// Modify all fields of copied
	copied.Name = "Charlotte"
	copied.Address.StreetAddress = "222 Maple St"
	copied.Address.City = "Boulder"
	copied.Address.Country = "Canada"

	// Remove original friends and add new ones
	copied.Friends = []string{"Henry", "Iris"}

	// Verify original is completely unchanged
	if original.Name != "Charlie" ||
		original.Address.StreetAddress != "111 Elm St" ||
		original.Address.City != "Denver" ||
		original.Address.Country != "USA" ||
		len(original.Friends) != 4 {
		t.Error("Original was modified when copy was changed")
	}

	// Verify copy has all changes
	if copied.Name != "Charlotte" ||
		copied.Address.StreetAddress != "222 Maple St" ||
		copied.Address.City != "Boulder" ||
		copied.Address.Country != "Canada" ||
		len(copied.Friends) != 2 {
		t.Error("Copy does not have expected changes")
	}
}
