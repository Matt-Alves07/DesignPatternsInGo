package main

import "testing"

func TestPersonBuilder(t *testing.T) {
	pb := NewPersonBuilder()
	person := pb.
		Lives().
			At("123 Test St").
			In("Test City").
			WithPostcode("12345").
		Works().
			At("Tech Corp").
			AsA("Engineer").
			Earning(100000).
		Build()

	// Verify Address
	if person.StreetAddress != "123 Test St" {
		t.Errorf("Expected StreetAddress '123 Test St', got '%s'", person.StreetAddress)
	}
	if person.City != "Test City" {
		t.Errorf("Expected City 'Test City', got '%s'", person.City)
	}
	if person.Postcode != "12345" {
		t.Errorf("Expected Postcode '12345', got '%s'", person.Postcode)
	}

	// Verify Job
	if person.CompanyName != "Tech Corp" {
		t.Errorf("Expected CompanyName 'Tech Corp', got '%s'", person.CompanyName)
	}
	if person.Position != "Engineer" {
		t.Errorf("Expected Position 'Engineer', got '%s'", person.Position)
	}
	if person.AnnualIncome != 100000 {
		t.Errorf("Expected AnnualIncome 100000, got %d", person.AnnualIncome)
	}
}
