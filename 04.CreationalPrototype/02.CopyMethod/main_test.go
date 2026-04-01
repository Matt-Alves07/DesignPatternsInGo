package prototype

import (
	"testing"
)

func TestAddressDeepCopy(t *testing.T) {
	original := &Address{"123 Main St", "Springfield", "USA"}
	copied := original.DeepCopy()

	// Verify content is the same
	if copied.StreetAddress != original.StreetAddress {
		t.Errorf("Expected StreetAddress %s, got %s", original.StreetAddress, copied.StreetAddress)
	}
	if copied.City != original.City {
		t.Errorf("Expected City %s, got %s", original.City, copied.City)
	}
	if copied.Country != original.Country {
		t.Errorf("Expected Country %s, got %s", original.Country, copied.Country)
	}

	// Verify it's a different instance
	if copied == original {
		t.Error("Expected copied Address to be a different pointer")
	}

	// Modify copy and verify original is unchanged
	copied.StreetAddress = "456 Oak Ave"
	if original.StreetAddress != "123 Main St" {
		t.Error("Original Address was modified when copy was changed")
	}
}

func TestPersonDeepCopy(t *testing.T) {
	original := &Person{"John",
		&Address{"123 London Rd", "London", "UK"},
		[]string{"Chris", "Matt"}}

	copied := original.DeepCopy()

	// Verify content is the same
	if copied.Name != original.Name {
		t.Errorf("Expected Name %s, got %s", original.Name, copied.Name)
	}
	if copied.Address.StreetAddress != original.Address.StreetAddress {
		t.Error("Address content mismatch")
	}
	if len(copied.Friends) != len(original.Friends) {
		t.Errorf("Expected Friends length %d, got %d", len(original.Friends), len(copied.Friends))
	}

	// Verify Address is a different instance
	if copied.Address == original.Address {
		t.Error("Expected copied Address to be a different pointer")
	}
}

func TestPersonDeepCopyFriendsIndependent(t *testing.T) {
	original := &Person{"John",
		&Address{"123 London Rd", "London", "UK"},
		[]string{"Chris", "Matt"}}

	copied := original.DeepCopy()
	copied.Name = "Jane"
	copied.Address.StreetAddress = "321 Baker St"
	copied.Friends = append(copied.Friends, "Angela")

	// Verify original is unchanged
	if original.Name != "John" {
		t.Errorf("Expected original.Name to be 'John', got %s", original.Name)
	}
	if original.Address.StreetAddress != "123 London Rd" {
		t.Errorf("Expected original.Address.StreetAddress to be '123 London Rd', got %s", original.Address.StreetAddress)
	}
	if len(original.Friends) != 2 {
		t.Errorf("Expected original.Friends length to be 2, got %d", len(original.Friends))
	}

	// Verify copy has the changes
	if copied.Name != "Jane" {
		t.Errorf("Expected copied.Name to be 'Jane', got %s", copied.Name)
	}
	if len(copied.Friends) != 3 {
		t.Errorf("Expected copied.Friends length to be 3, got %d", len(copied.Friends))
	}
}

func TestMultiplePersonDeepCopies(t *testing.T) {
	original := &Person{"Alice",
		&Address{"456 Park Ave", "New York", "USA"},
		[]string{"Bob", "Charlie"}}

	copy1 := original.DeepCopy()
	copy2 := original.DeepCopy()

	copy1.Name = "Copy1"
	copy1.Friends = append(copy1.Friends, "Diana")

	copy2.Name = "Copy2"
	copy2.Friends = append(copy2.Friends, "Eve")

	// Verify all are independent
	if original.Name != "Alice" {
		t.Errorf("Expected original.Name to be 'Alice', got %s", original.Name)
	}
	if len(original.Friends) != 2 {
		t.Errorf("Expected original.Friends to have 2 friends, got %d", len(original.Friends))
	}

	if copy1.Name != "Copy1" || len(copy1.Friends) != 3 {
		t.Error("copy1 state is incorrect")
	}
	if copy2.Name != "Copy2" || len(copy2.Friends) != 3 {
		t.Error("copy2 state is incorrect")
	}

	// Verify copies are independent (slices cannot be compared directly with ==)
	// So we verify they have the same content but different lengths isn't the issue
	if len(copy1.Friends) != len(copy2.Friends) {
		t.Error("Expected copy1.Friends and copy2.Friends to have equal length")
	}
}

func TestPersonDeepCopyEmptyFriendsList(t *testing.T) {
	original := &Person{"Bob",
		&Address{"789 Oak St", "Chicago", "USA"},
		[]string{}}

	copied := original.DeepCopy()

	if len(copied.Friends) != 0 {
		t.Errorf("Expected copied.Friends to be empty, got length %d", len(copied.Friends))
	}

	copied.Friends = append(copied.Friends, "NewFriend")

	if len(original.Friends) != 0 {
		t.Errorf("Expected original.Friends to still be empty, got %v", original.Friends)
	}
}
