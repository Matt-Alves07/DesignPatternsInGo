package main

import "testing"

func TestRectangle(t *testing.T) {
	r := &Rectangle{width: 2, height: 3}
	if r.GetWidth() != 2 || r.GetHeight() != 3 {
		t.Errorf("Expected 2x3, got %dx%d", r.GetWidth(), r.GetHeight())
	}

	r.SetWidth(4)
	if r.GetWidth() != 4 || r.GetHeight() != 3 {
		t.Errorf("Expected 4x3 (independent dimensions), got %dx%d", r.GetWidth(), r.GetHeight())
	}
}

func TestSquare(t *testing.T) {
	s := NewSquare(5)
	if s.GetWidth() != 5 || s.GetHeight() != 5 {
		t.Errorf("Expected 5x5, got %dx%d", s.GetWidth(), s.GetHeight())
	}

	s.SetWidth(6)
	if s.GetWidth() != 6 || s.GetHeight() != 6 {
		t.Errorf("Expected 6x6 (linked dimensions for Square), got %dx%d", s.GetWidth(), s.GetHeight())
	}

	s.SetHeight(7)
	if s.GetWidth() != 7 || s.GetHeight() != 7 {
		t.Errorf("Expected 7x7 (linked dimensions for Square), got %dx%d", s.GetWidth(), s.GetHeight())
	}
}

// TestUseItLogic verifies the logic inside UseIt without printing.
// This confirms that Square violates the expectation of independent dimensions.
func TestUseItLogic(t *testing.T) {
	// Test Rectangle (Should Pass LSP expectation)
	rc := &Rectangle{2, 3}
	initialWidth := rc.GetWidth()
	rc.SetHeight(10)
	expectedArea := 10 * initialWidth
	actualArea := rc.GetWidth() * rc.GetHeight()
	if expectedArea != actualArea {
		t.Errorf("Rectangle failed LSP check: Expected area %d, got %d", expectedArea, actualArea)
	}

	// Test Square (Should Fail LSP expectation)
	sq := NewSquare(5)
	initialWidthSq := sq.GetWidth()
	sq.SetHeight(10)
	expectedAreaSq := 10 * initialWidthSq // Expect 10 * 5 = 50
	actualAreaSq := sq.GetWidth() * sq.GetHeight() // Actual 10 * 10 = 100
	
	// Note: In a real test suite ensuring correctness of *implementations*, we might not assert failure.
	// But here we assert the *property* that it behaves like a Square, thus failing the Rectangle assumption.
	if expectedAreaSq == actualAreaSq {
		t.Errorf("Square behaved like a Rectangle? It shouldn't if it's a true square. Expected %d != %d", expectedAreaSq, actualAreaSq)
	}
}
