package main

import (
	"testing"
)

func TestBufferCreation(t *testing.T) {
	buffer := NewBuffer(10, 10)

	if buffer == nil {
		t.Fatal("Expected buffer to be created")
	}

	if buffer.width != 10 || buffer.height != 10 {
		t.Errorf("Expected dimensions 10x10, got %dx%d", buffer.width, buffer.height)
	}

	if len(buffer.buffer) != 100 {
		t.Errorf("Expected 100 elements, got %d", len(buffer.buffer))
	}
}

func TestBufferAt(t *testing.T) {
	buffer := NewBuffer(5, 5)
	// Buffer is initialized with zero runes
	char := buffer.At(0)

	if char != 0 {
		t.Errorf("Expected zero rune at index 0, got %v", char)
	}
}

func TestViewportCreation(t *testing.T) {
	buffer := NewBuffer(10, 10)
	viewport := NewViewport(buffer)

	if viewport == nil {
		t.Fatal("Expected viewport to be created")
	}

	if viewport.buffer != buffer {
		t.Error("Expected viewport to reference the buffer")
	}

	if viewport.offset != 0 {
		t.Errorf("Expected initial offset 0, got %d", viewport.offset)
	}
}

func TestViewportGetCharacterAt(t *testing.T) {
	buffer := NewBuffer(10, 10)
	viewport := NewViewport(buffer)

	char := viewport.GetCharacterAt(0)
	if char != 0 {
		t.Errorf("Expected zero rune, got %v", char)
	}
}

func TestConsoleCreation(t *testing.T) {
	console := NewConsole()

	if console == nil {
		t.Fatal("Expected console to be created")
	}

	if len(console.buffers) != 1 {
		t.Errorf("Expected 1 buffer, got %d", len(console.buffers))
	}

	if len(console.viewports) != 1 {
		t.Errorf("Expected 1 viewport, got %d", len(console.viewports))
	}

	if console.offset != 0 {
		t.Errorf("Expected initial offset 0, got %d", console.offset)
	}
}

func TestConsoleGetCharacterAt(t *testing.T) {
	console := NewConsole()
	char := console.GetCharacterAt(0)

	if char != 0 {
		t.Errorf("Expected zero rune, got %v", char)
	}
}

func TestConsoleFacadeSimplification(t *testing.T) {
	console := NewConsole()

	// Console simplifies access to Buffer and Viewport
	// Users don't need to know about Buffer or Viewport internals
	char1 := console.GetCharacterAt(0)
	char2 := console.GetCharacterAt(5)
	char3 := console.GetCharacterAt(10)

	if char1 != 0 || char2 != 0 || char3 != 0 {
		t.Error("Expected all characters to be zero runes")
	}
}

func TestBufferAndViewportInteraction(t *testing.T) {
	buffer := NewBuffer(10, 5)
	viewport := NewViewport(buffer)

	// Viewport should access characters through buffer with offset
	char := viewport.GetCharacterAt(2)
	if char != 0 {
		t.Errorf("Expected zero rune at offset position, got %v", char)
	}
}

func TestMultipleViewports(t *testing.T) {
	buffer := NewBuffer(20, 10)
	viewport1 := NewViewport(buffer)
	viewport2 := NewViewport(buffer)

	char1 := viewport1.GetCharacterAt(0)
	char2 := viewport2.GetCharacterAt(0)

	if char1 != char2 {
		t.Error("Expected both viewports to see the same buffer content")
	}
}

func TestConsoleBufferDimensions(t *testing.T) {
	console := NewConsole()

	// Console should create a 10x10 buffer
	if console.buffers[0].width != 10 || console.buffers[0].height != 10 {
		t.Errorf("Expected 10x10 buffer, got %dx%d",
			console.buffers[0].width, console.buffers[0].height)
	}
}
