package main

import "fmt"

// Employee represents an employee with name, position, and income.
type Employee struct {
	Name, Position string
	AnnualIncome   int
}

// functional approach

// NewEmployeeFactory creates a factory function for a specific role and income.
// It returns a function that creates an Employee with the given name.
func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, annualIncome}
	}
}

// structural approach

// EmployeeFactory struct encapsulates data to create Employees.
type EmployeeFactory struct {
	Position     string
	AnnualIncome int
}

// NewEmployeeFactory2 creates a new EmployeeFactory struct.
func NewEmployeeFactory2(position string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{position, annualIncome}
}

// Create generates a new Employee using the factory's settings.
func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.Position, f.AnnualIncome}
}

func main() {
	// 1. Functional Approach
	developerFactory := NewEmployeeFactory("Developer", 60000)
	managerFactory := NewEmployeeFactory("Manager", 80000)

	developer := developerFactory("Adam")
	fmt.Println(developer)

	manager := managerFactory("Jane")
	fmt.Println(manager)

	// 2. Structural Approach
	bossFactory := NewEmployeeFactory2("CEO", 100000)
	// can modify post-hoc
	bossFactory.AnnualIncome = 110000
	boss := bossFactory.Create("Sam")
	fmt.Println(boss)
}