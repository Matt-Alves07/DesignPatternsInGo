package main

import (
	"testing"
)

func TestFormattedTextCreation(t *testing.T) {
	text := NewFormattedText("hello")

	if text == nil {
		t.Fatal("Expected FormattedText to be created")
	}

	if len(text.plainText) != 5 {
		t.Errorf("Expected plainText length 5, got %d", len(text.plainText))
	}

	if len(text.capitalize) != 5 {
		t.Errorf("Expected capitalize array length 5, got %d", len(text.capitalize))
	}
}

func TestFormattedTextString(t *testing.T) {
	text := NewFormattedText("hello")
	output := text.String()

	if output != "hello" {
		t.Errorf("Expected 'hello', got '%s'", output)
	}
}

func TestFormattedTextCapitalize(t *testing.T) {
	text := NewFormattedText("hello")
	text.Capitalize(0, 0)

	output := text.String()
	if output != "Hello" {
		t.Errorf("Expected 'Hello' with first char capitalized, got '%s'", output)
	}
}

func TestFormattedTextCapitalizeRange(t *testing.T) {
	text := NewFormattedText("hello")
	text.Capitalize(0, 4)

	output := text.String()
	if output != "HELLO" {
		t.Errorf("Expected 'HELLO' all caps, got '%s'", output)
	}
}

func TestBetterFormattedTextCreation(t *testing.T) {
	text := NewBetterFormattedText("hello")

	if text == nil {
		t.Fatal("Expected BetterFormattedText to be created")
	}

	if text.plainText != "hello" {
		t.Errorf("Expected plainText 'hello', got '%s'", text.plainText)
	}

	if len(text.formatting) != 0 {
		t.Errorf("Expected 0 formatting ranges initially, got %d", len(text.formatting))
	}
}

func TestBetterFormattedTextString(t *testing.T) {
	text := NewBetterFormattedText("hello")
	output := text.String()

	if output != "hello" {
		t.Errorf("Expected 'hello', got '%s'", output)
	}
}

func TestBetterFormattedTextRange(t *testing.T) {
	text := NewBetterFormattedText("hello")
	textRange := text.Range(0, 0)

	if textRange == nil {
		t.Fatal("Expected TextRange to be created")
	}

	if len(text.formatting) != 1 {
		t.Errorf("Expected 1 formatting range, got %d", len(text.formatting))
	}
}

func TestBetterFormattedTextRangeCapitalize(t *testing.T) {
	text := NewBetterFormattedText("hello")
	textRange := text.Range(0, 0)
	textRange.Capitalize = true

	output := text.String()
	if output != "Hello" {
		t.Errorf("Expected 'Hello' with first char capitalized, got '%s'", output)
	}
}

func TestBetterFormattedTextMultipleRanges(t *testing.T) {
	text := NewBetterFormattedText("hello world")

	r1 := text.Range(0, 4)
	r1.Capitalize = true

	r2 := text.Range(6, 10)
	r2.Capitalize = true

	output := text.String()
	if output != "HELLO WORLD" {
		t.Errorf("Expected 'HELLO WORLD', got '%s'", output)
	}
}

func TestTextRangeCoverage(t *testing.T) {
	tr := TextRange{Start: 2, End: 5}

	if !tr.Covers(2) {
		t.Error("Expected range to cover position 2")
	}

	if !tr.Covers(4) {
		t.Error("Expected range to cover position 4")
	}

	if !tr.Covers(5) {
		t.Error("Expected range to cover position 5")
	}

	if tr.Covers(1) {
		t.Error("Expected range not to cover position 1")
	}

	if tr.Covers(6) {
		t.Error("Expected range not to cover position 6")
	}
}

func TestBetterFormattedTextMemoryEfficiency(t *testing.T) {
	text := NewBetterFormattedText("hello world hello world")

	// BetterFormattedText uses TextRanges - should be more memory efficient
	// than FormattedText which uses a bool array for each character
	r1 := text.Range(0, 4)
	r1.Capitalize = true

	r2 := text.Range(18, 22)
	r2.Capitalize = true

	output := text.String()
	expected := "HELLO world hello WORLD"
	if output != expected {
		t.Errorf("Expected '%s', got '%s'", expected, output)
	}
}

func TestFormattedTextEmptyString(t *testing.T) {
	text := NewFormattedText("")

	if text.String() != "" {
		t.Error("Expected empty string for empty text")
	}
}

func TestBetterFormattedTextEmptyString(t *testing.T) {
	text := NewBetterFormattedText("")

	if text.String() != "" {
		t.Error("Expected empty string for empty text")
	}
}
