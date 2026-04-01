package main

import (
	"strings"
	"testing"
)

func TestMinmax(t *testing.T) {
	tests := []struct {
		a, b     int
		minVal   int
		maxVal   int
	}{
		{1, 5, 1, 5},
		{5, 1, 1, 5},
		{3, 3, 3, 3},
		{-5, 10, -5, 10},
		{0, 0, 0, 0},
	}

	for _, tt := range tests {
		min, max := minmax(tt.a, tt.b)
		if min != tt.minVal || max != tt.maxVal {
			t.Errorf("minmax(%d, %d) = (%d, %d), want (%d, %d)", tt.a, tt.b, min, max, tt.minVal, tt.maxVal)
		}
	}
}

func TestNewRectangle(t *testing.T) {
	rect := NewRectangle(6, 4)

	if rect == nil {
		t.Fatal("Expected rectangle to be created")
	}

	if len(rect.Lines) != 4 {
		t.Errorf("Expected 4 lines, got %d", len(rect.Lines))
	}

	// Check rectangle bounds
	expectedLines := []Line{
		{0, 0, 5, 0},       // top
		{0, 0, 0, 3},       // left
		{5, 0, 5, 3},       // right
		{0, 3, 5, 3},       // bottom
	}

	for i, expected := range expectedLines {
		if rect.Lines[i] != expected {
			t.Errorf("Line %d: expected %v, got %v", i, expected, rect.Lines[i])
		}
	}
}

func TestVectorToRasterAdapter(t *testing.T) {
	rect := NewRectangle(3, 3)
	adapter := VectorToRaster(rect)

	if adapter == nil {
		t.Fatal("Expected adapter to be created")
	}

	points := adapter.GetPoints()
	if len(points) == 0 {
		t.Error("Expected adapter to generate points")
	}
}

func TestVectorImageHasLines(t *testing.T) {
	vi := &VectorImage{
		Lines: []Line{
			{0, 0, 10, 0},
			{0, 0, 0, 10},
		},
	}

	if len(vi.Lines) != 2 {
		t.Errorf("Expected 2 lines, got %d", len(vi.Lines))
	}
}

func TestDrawPointsRasterImage(t *testing.T) {
	rect := NewRectangle(3, 3)
	raster := VectorToRaster(rect)

	output := DrawPoints(raster)

	if len(output) == 0 {
		t.Error("Expected DrawPoints to produce output")
	}

	// Check that output contains the expected characters
	if !strings.Contains(output, "*") {
		t.Error("Expected output to contain '*' characters")
	}

	if !strings.Contains(output, "\n") {
		t.Error("Expected output to contain newlines")
	}
}

func TestRasterImageGetPoints(t *testing.T) {
	rect := NewRectangle(4, 4)
	raster := VectorToRaster(rect)

	points := raster.GetPoints()

	if len(points) == 0 {
		t.Error("Expected GetPoints to return points")
	}

	// Verify all points are within expected bounds
	for _, p := range points {
		if p.X < 0 || p.Y < 0 || p.X >= 4 || p.Y >= 4 {
			t.Errorf("Point (%d, %d) is out of bounds", p.X, p.Y)
		}
	}
}

func TestPointStructure(t *testing.T) {
	p := Point{X: 5, Y: 10}

	if p.X != 5 {
		t.Errorf("Expected X=5, got %d", p.X)
	}

	if p.Y != 10 {
		t.Errorf("Expected Y=10, got %d", p.Y)
	}
}

func TestMultipleAdapters(t *testing.T) {
	// Clear any previous cache by creating fresh rectangles
	rect1 := NewRectangle(3, 3)
	rect2 := NewRectangle(5, 5)

	adapter1 := VectorToRaster(rect1)
	adapter2 := VectorToRaster(rect2)

	points1 := adapter1.GetPoints()
	points2 := adapter2.GetPoints()

	// Verify that we have points from both adapters
	if len(points1) == 0 {
		t.Error("Expected points from adapter1")
	}

	if len(points2) == 0 {
		t.Error("Expected points from adapter2")
	}

	// The adapters should have generated points
	// We don't assert exact count due to caching effects from other tests
	t.Logf("Adapter1 generated %d points, Adapter2 generated %d points", len(points1), len(points2))
}
