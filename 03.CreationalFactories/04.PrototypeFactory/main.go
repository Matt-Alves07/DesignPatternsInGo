package main

import "fmt"

// Employee represents an employee with name, position, and income.
type Employee struct {
	Name, Position string
	AnnualIncome   int
}

// Role constants
const (
	Developer = iota
	Manager
)

// NewEmployee creates a new Employee based on a role prototype.
// This acts as a factory using pre-configured prototypes.
func NewEmployee(role int) *Employee {
	switch role {
	case Developer:
		return &Employee{"", "Developer", 60000}
	case Manager:
		return &Employee{"", "Manager", 80000}
	default:
		panic("unsupported role")
	}
}

func main() {
	m := NewEmployee(Manager)
	m.Name = "Sam"
	fmt.Println(m)
}