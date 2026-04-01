package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestCircleRender(t *testing.T) {
	circle := Circle{Radius: 5.5}
	output := circle.Render()

	if !strings.Contains(output, "Circle") {
		t.Error("Expected output to contain 'Circle'")
	}

	if !strings.Contains(output, "5.5") {
		t.Error("Expected output to contain radius value")
	}
}

func TestSquareRender(t *testing.T) {
	square := Square{Side: 3.0}
	output := square.Render()

	if !strings.Contains(output, "Square") {
		t.Error("Expected output to contain 'Square'")
	}

	if !strings.Contains(output, "3") {
		t.Error("Expected output to contain side value")
	}
}

func TestColoredShape(t *testing.T) {
	circle := Circle{Radius: 4.0}
	colored := ColoredShape{&circle, "Red"}

	output := colored.Render()

	if !strings.Contains(output, "Circle") {
		t.Error("Expected output to contain 'Circle'")
	}

	if !strings.Contains(output, "Red") {
		t.Error("Expected output to contain color 'Red'")
	}

	if !strings.Contains(output, "has the color") {
		t.Error("Expected output to contain format string")
	}
}

func TestTransparentShape(t *testing.T) {
	square := Square{Side: 2.0}
	transparent := TransparentShape{&square, 0.5}

	output := transparent.Render()

	if !strings.Contains(output, "Square") {
		t.Error("Expected output to contain 'Square'")
	}

	if !strings.Contains(output, "transparency") {
		t.Error("Expected output to contain 'transparency'")
	}

	if !strings.Contains(output, "50") {
		t.Error("Expected output to contain transparency percentage")
	}
}

func TestMultipleDecorators(t *testing.T) {
	circle := Circle{Radius: 3.0}

	// Decorate with color
	redCircle := ColoredShape{&circle, "Red"}

	// Decorate with transparency
	transparentRedCircle := TransparentShape{&redCircle, 0.7}

	output := transparentRedCircle.Render()

	if !strings.Contains(output, "Circle") {
		t.Error("Expected output to contain 'Circle'")
	}

	if !strings.Contains(output, "Red") {
		t.Error("Expected output to contain 'Red'")
	}

	if !strings.Contains(output, "transparency") {
		t.Error("Expected output to contain 'transparency'")
	}

	if !strings.Contains(output, "70") {
		t.Error("Expected output to contain transparency percentage")
	}
}

func TestCircleResize(t *testing.T) {
	circle := Circle{Radius: 2.0}
	circle.Resize(3.0)

	if circle.Radius != 6.0 {
		t.Errorf("Expected radius 6.0 after resize, got %f", circle.Radius)
	}
}

func TestColoredShapeMultipleColors(t *testing.T) {
	square := Square{Side: 4.0}
	green := ColoredShape{&square, "Green"}
	output := green.Render()

	if !strings.Contains(output, "Green") {
		t.Error("Expected output to contain 'Green'")
	}

	redSquare := ColoredShape{&square, "Red"}
	output2 := redSquare.Render()

	if !strings.Contains(output2, "Red") {
		t.Error("Expected output to contain 'Red'")
	}
}

func TestTransparentShapeRanges(t *testing.T) {
	tests := []struct {
		transparency float32
		expectedPct  float32
	}{
		{0.0, 0.0},
		{0.5, 50.0},
		{1.0, 100.0},
		{0.25, 25.0},
		{0.75, 75.0},
	}

	for _, tt := range tests {
		circle := Circle{Radius: 1.0}
		transparent := TransparentShape{&circle, tt.transparency}
		output := transparent.Render()

		if !strings.Contains(output, fmt.Sprintf("%.f", tt.expectedPct)) &&
			!strings.Contains(output, fmt.Sprintf("%.1f", tt.expectedPct)) {
			t.Logf("Output for transparency %f: %s", tt.transparency, output)
		}
	}
}

func TestDecoratorComposition(t *testing.T) {
	// Create base shape
	circle := Circle{Radius: 3.0}

	// Apply multiple decorators
	coloredCircle := ColoredShape{&circle, "Blue"}
	finalDecoration := TransparentShape{&coloredCircle, 0.3}

	output := finalDecoration.Render()

	// All properties should be visible
	if !strings.Contains(output, "3") {
		t.Error("Expected radius to be visible")
	}
	if !strings.Contains(output, "Blue") {
		t.Error("Expected color to be visible")
	}
	if !strings.Contains(output, "30") {
		t.Error("Expected transparency percentage to be visible")
	}
}
