package prototype

import (
	"testing"
)

func TestDeepCopyingPerson(t *testing.T) {
	john := Person{"John",
		&Address{"123 London Rd", "London", "UK"}}

	// Deep copy: create new Address struct
	jane := john
	jane.Address = &Address{
		john.Address.StreetAddress,
		john.Address.City,
		john.Address.Country,
	}

	jane.Name = "Jane"
	jane.Address.StreetAddress = "321 Baker St"

	// Verify original is unchanged
	if john.Name != "John" {
		t.Errorf("Expected john.Name to be 'John', got %s", john.Name)
	}
	if john.Address.StreetAddress != "123 London Rd" {
		t.Errorf("Expected john.Address.StreetAddress to be '123 London Rd', got %s", john.Address.StreetAddress)
	}

	// Verify copy is changed
	if jane.Name != "Jane" {
		t.Errorf("Expected jane.Name to be 'Jane', got %s", jane.Name)
	}
	if jane.Address.StreetAddress != "321 Baker St" {
		t.Errorf("Expected jane.Address.StreetAddress to be '321 Baker St', got %s", jane.Address.StreetAddress)
	}

	// Verify addresses are different
	if john.Address == jane.Address {
		t.Error("Expected john.Address and jane.Address to be different pointers")
	}
}

func TestAddressFieldsAreIndependent(t *testing.T) {
	original := Person{"Alice",
		&Address{"456 Park Ave", "New York", "USA"}}

	// Create deep copy
	copy := original
	copy.Address = &Address{
		original.Address.StreetAddress,
		original.Address.City,
		original.Address.Country,
	}

	// Modify copy's address
	copy.Address.City = "Boston"
	copy.Address.Country = "USA"

	// Verify original is unchanged
	if original.Address.City != "New York" {
		t.Errorf("Expected original.Address.City to be 'New York', got %s", original.Address.City)
	}
}

func TestMultipleDeepCopies(t *testing.T) {
	original := Person{"Bob",
		&Address{"789 Oak St", "Chicago", "USA"}}

	// Create two independent copies
	copy1 := original
	copy1.Address = &Address{
		original.Address.StreetAddress,
		original.Address.City,
		original.Address.Country,
	}

	copy2 := original
	copy2.Address = &Address{
		original.Address.StreetAddress,
		original.Address.City,
		original.Address.Country,
	}

	// Modify both copies
	copy1.Name = "Charlie"
	copy1.Address.StreetAddress = "111 Elm St"

	copy2.Name = "Diana"
	copy2.Address.StreetAddress = "222 Maple St"

	// Verify all three are independent
	if original.Name != "Bob" {
		t.Errorf("Expected original.Name to be 'Bob', got %s", original.Name)
	}
	if original.Address.StreetAddress != "789 Oak St" {
		t.Errorf("Expected original.Address.StreetAddress to be '789 Oak St', got %s", original.Address.StreetAddress)
	}

	if copy1.Name != "Charlie" || copy1.Address.StreetAddress != "111 Elm St" {
		t.Error("copy1 was not modified correctly")
	}

	if copy2.Name != "Diana" || copy2.Address.StreetAddress != "222 Maple St" {
		t.Error("copy2 was not modified correctly")
	}

	// Verify copies are different from each other
	if copy1.Address == copy2.Address {
		t.Error("Expected copy1.Address and copy2.Address to be different pointers")
	}
}
