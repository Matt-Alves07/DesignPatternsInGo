package main

import "fmt"

// OCP (Open-Closed Principle): Classes should be open for extension but closed for modification.
// This example demonstrates how to apply OCP using the Specification pattern.

// Color represents a product color.
type Color int

// Color constants
const (
	red Color = iota
	green
	blue
)

// Size represents a product size.
type Size int

// Size constants
const (
	small Size = iota
	medium
	large
)

// Product defines a product with properties.
type Product struct {
	name  string
	color Color
	size  Size
}

// Filter is a legacy struct that violates OCP because we have to add new methods
// every time we want to filter by a new criteria.
type Filter struct {
}

// filterByColor filters products by color.
func (f *Filter) filterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

// filterBySize filters products by size.
func (f *Filter) filterBySize(products []Product, size Size) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.size == size {
			result = append(result, &products[i])
		}
	}
	return result
}

// filterBySizeAndColor filters by both size and color. This already causes explosion of methods.
func (f *Filter) filterBySizeAndColor(products []Product, size Size, color Color) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.size == size && v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

// Specification is an interface that allows defining filtering criteria.
// This supports OCP because we can add new specifications without modifying existing code.
type Specification interface {
	IsSatisfied(p *Product) bool
}

// ColorSpecification defines a filter for a specific color.
type ColorSpecification struct {
	color Color
}

// IsSatisfied satisfies the Specification interface for ColorSpecification.
func (spec ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == spec.color
}

// SizeSpecification defines a filter for a specific size.
type SizeSpecification struct {
	size Size
}

// IsSatisfied satisfies the Specification interface for SizeSpecification.
func (spec SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == spec.size
}

// AndSpecification is a combinator that allows combining two specifications.
type AndSpecification struct {
	first, second Specification
}

// IsSatisfied satisfies the Specification interface for AndSpecification.
func (spec AndSpecification) IsSatisfied(p *Product) bool {
	return spec.first.IsSatisfied(p) && spec.second.IsSatisfied(p)
}

// BetterFilter implements filtering using the Specification interface.
// It complies with OCP as it doesn't need modification to support new filters.
type BetterFilter struct{}

// Filter filters products based on the provided specification.
func (f *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}

	fmt.Print("Green products (old):\n")
	f := Filter{}
	for _, v := range f.filterByColor(products, green) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	fmt.Print("Green products (new):\n")
	greenSpec := ColorSpecification{green}
	bf := BetterFilter{}
	for _, v := range bf.Filter(products, greenSpec) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	largeSpec := SizeSpecification{large}
	largeGreenSpec := AndSpecification{largeSpec, greenSpec}
	fmt.Print("Large green items (new):\n") // Changed text to match logic (large & green)
	for _, v := range bf.Filter(products, largeGreenSpec) {
		fmt.Printf(" - %s is large and green\n", v.name)
	}
}






