package main

import "fmt"

// LSP (Liskov Substitution Principle): Objects of a superclass should be replaceable with objects
// of its subclasses without affecting the correctness of the program.
// In Go, this applies to interfaces: if a type implements an interface, it should behave
// in a way that is consistent with the expectations of that interface.

// Sized represents an object with dimensions.
type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

// Rectangle represents a rectangle shape.
type Rectangle struct {
	width, height int
}

// GetWidth returns the width of the rectangle.
func (r *Rectangle) GetWidth() int {
	return r.width
}

// SetWidth sets the width of the rectangle.
func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

// GetHeight returns the height of the rectangle.
func (r *Rectangle) GetHeight() int {
	return r.height
}

// SetHeight sets the height of the rectangle.
func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

// Square represents a square.
// It embeds Rectangle but enforces that width and height are always equal.
// This violates LSP when used as a Sized interface expecting independent dimensions.
type Square struct {
	Rectangle
}

// NewSquare creates a new Square.
func NewSquare(size int) *Square {
	sq := Square{}
	sq.width = size
	sq.height = size
	return &sq
}

// SetWidth sets both width and height to ensure square property.
func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = width
}

// SetHeight sets both width and height to ensure square property.
func (s *Square) SetHeight(height int) {
	s.width = height
	s.height = height
}

// UseIt is a function that demonstrates the behavior of Sized objects.
// It expects that setting the height does not affect the width.
func UseIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := 10 * width
	actualArea := sized.GetWidth() * sized.GetHeight()
	fmt.Print("Expected an area of ", expectedArea,
		", but got ", actualArea, "\n")
}

func main() {
	rc := &Rectangle{2, 3}
	UseIt(rc)

	sq := NewSquare(5)
	UseIt(sq)
}