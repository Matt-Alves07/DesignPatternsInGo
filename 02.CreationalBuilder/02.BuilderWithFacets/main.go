package main

import "fmt"

// Person represents a detailed person struct with address and job info.
type Person struct {
	StreetAddress, Postcode, City string
	CompanyName, Position         string
	AnnualIncome                  int
}

// PersonBuilder is the starting point for building a Person.
// It maintains a reference to the person object being built.
type PersonBuilder struct {
	person *Person // needs to be inited
}

// NewPersonBuilder creates a new PersonBuilder.
func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

// Build returns the built Person.
func (it *PersonBuilder) Build() *Person {
	return it.person
}

// Works switches context to PersonJobBuilder.
func (it *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*it}
}

// Lives switches context to PersonAddressBuilder.
func (it *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*it}
}

// PersonJobBuilder is a faceted builder for job-related information.
type PersonJobBuilder struct {
	PersonBuilder
}

// At sets the company name.
func (pjb *PersonJobBuilder) At(companyName string) *PersonJobBuilder {
	pjb.person.CompanyName = companyName
	return pjb
}

// AsA sets the job position.
func (pjb *PersonJobBuilder) AsA(position string) *PersonJobBuilder {
	pjb.person.Position = position
	return pjb
}

// Earning sets the annual income.
func (pjb *PersonJobBuilder) Earning(annualIncome int) *PersonJobBuilder {
	pjb.person.AnnualIncome = annualIncome
	return pjb
}

// PersonAddressBuilder is a faceted builder for address-related information.
type PersonAddressBuilder struct {
	PersonBuilder
}

// At sets the street address.
func (it *PersonAddressBuilder) At(streetAddress string) *PersonAddressBuilder {
	it.person.StreetAddress = streetAddress
	return it
}

// In sets the city.
func (it *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	it.person.City = city
	return it
}

// WithPostcode sets the postcode.
func (it *PersonAddressBuilder) WithPostcode(postcode string) *PersonAddressBuilder {
	it.person.Postcode = postcode
	return it
}

func main() {
	pb := NewPersonBuilder()
	pb.
		Lives().
		At("123 London Road").
		In("London").
		WithPostcode("SW12BC").
		Works().
		At("Fabrikam").
		AsA("Programmer").
		Earning(123000)
	person := pb.Build()
	fmt.Println(*person)
}