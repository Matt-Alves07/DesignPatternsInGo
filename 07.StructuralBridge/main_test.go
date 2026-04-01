package main

import (
	"testing"
)

func TestVectorRendererRenderCircle(t *testing.T) {
	renderer := &VectorRenderer{}

	// Verify renderer implements Renderer interface
	var r Renderer = renderer
	if r == nil {
		t.Error("VectorRenderer should implement Renderer interface")
	}

	renderer.RenderCircle(5.0)
	// No error should occur
}

func TestRasterRendererRenderCircle(t *testing.T) {
	renderer := &RasterRenderer{Dpi: 300}

	// Verify renderer implements Renderer interface
	var r Renderer = renderer
	if r == nil {
		t.Error("RasterRenderer should implement Renderer interface")
	}

	renderer.RenderCircle(10.0)
	// No error should occur
}

func TestCircleCreation(t *testing.T) {
	renderer := &VectorRenderer{}
	circle := NewCircle(renderer, 5.0)

	if circle == nil {
		t.Fatal("Expected circle to be created")
	}

	if circle.radius != 5.0 {
		t.Errorf("Expected radius 5.0, got %f", circle.radius)
	}

	if circle.renderer != renderer {
		t.Error("Expected renderer to be set correctly")
	}
}

func TestCircleDraw(t *testing.T) {
	renderer := &VectorRenderer{}
	circle := NewCircle(renderer, 7.5)

	// Should not panic
	circle.Draw()
}

func TestCircleResize(t *testing.T) {
	renderer := &VectorRenderer{}
	circle := NewCircle(renderer, 5.0)

	circle.Resize(2.0)

	if circle.radius != 10.0 {
		t.Errorf("Expected radius 10.0 after resize, got %f", circle.radius)
	}
}

func TestCircleWithVectorRenderer(t *testing.T) {
	vRenderer := &VectorRenderer{}
	circle := NewCircle(vRenderer, 3.0)

	if circle.renderer != vRenderer {
		t.Error("Expected VectorRenderer to be set")
	}

	circle.Draw()
}

func TestCircleWithRasterRenderer(t *testing.T) {
	rRenderer := &RasterRenderer{Dpi: 300}
	circle := NewCircle(rRenderer, 4.0)

	if circle.renderer != rRenderer {
		t.Error("Expected RasterRenderer to be set")
	}

	circle.Draw()
}

func TestBridgeDecoupling(t *testing.T) {
	// Vector rendering
	vectorCircle := NewCircle(&VectorRenderer{}, 5.0)
	vectorCircle.Draw()

	// Raster rendering with same circle properties
	rasterCircle := NewCircle(&RasterRenderer{Dpi: 300}, 5.0)
	rasterCircle.Draw()

	// Both circles should have same radius
	if vectorCircle.radius != rasterCircle.radius {
		t.Error("Expected circles to have same radius")
	}
}

func TestMultipleCirclesWithDifferentRenderers(t *testing.T) {
	vectorRenderer := &VectorRenderer{}
	rasterRenderer := &RasterRenderer{Dpi: 72}

	circle1 := NewCircle(vectorRenderer, 2.0)
	circle2 := NewCircle(vectorRenderer, 3.0)
	circle3 := NewCircle(rasterRenderer, 4.0)

	circle1.Draw()
	circle2.Draw()
	circle3.Draw()

	if circle1.radius != 2.0 || circle2.radius != 3.0 || circle3.radius != 4.0 {
		t.Error("Expected circles to maintain their radii")
	}
}

func TestRasterRendererWithDifferentDpi(t *testing.T) {
	r72 := &RasterRenderer{Dpi: 72}
	r300 := &RasterRenderer{Dpi: 300}

	circle72 := NewCircle(r72, 5.0)
	circle300 := NewCircle(r300, 5.0)

	if circle72.radius != circle300.radius {
		t.Error("Expected circles to have same radius despite different DPI")
	}
}

func TestCircleResizeMultipleTimes(t *testing.T) {
	renderer := &VectorRenderer{}
	circle := NewCircle(renderer, 2.0)

	circle.Resize(2.0) // 4.0
	circle.Resize(1.5) // 6.0
	circle.Resize(0.5) // 3.0

	if circle.radius != 3.0 {
		t.Errorf("Expected radius 3.0 after multiple resizes, got %f", circle.radius)
	}
}
