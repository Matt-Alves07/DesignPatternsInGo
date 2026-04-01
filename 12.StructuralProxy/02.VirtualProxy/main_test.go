package main

import (
	"testing"
)

func TestBitmapCreation(t *testing.T) {
	bitmap := NewBitmap("test.png")

	if bitmap == nil {
		t.Fatal("Expected bitmap to be created")
	}

	if bitmap.filename != "test.png" {
		t.Errorf("Expected filename 'test.png', got '%s'", bitmap.filename)
	}
}

func TestBitmapDraw(t *testing.T) {
	bitmap := NewBitmap("test.png")
	bitmap.Draw()
	// Should not panic
}

func TestLazyBitmapCreation(t *testing.T) {
	lazyBitmap := NewLazyBitmap("test.png")

	if lazyBitmap == nil {
		t.Fatal("Expected LazyBitmap to be created")
	}

	if lazyBitmap.filename != "test.png" {
		t.Errorf("Expected filename 'test.png', got '%s'", lazyBitmap.filename)
	}

	// Bitmap should not be loaded yet
	if lazyBitmap.bitmap != nil {
		t.Error("Expected bitmap to not be loaded initially")
	}
}

func TestLazyBitmapDraw(t *testing.T) {
	lazyBitmap := NewLazyBitmap("test.png")

	// First Draw should load the bitmap
	lazyBitmap.Draw()

	// After first Draw, bitmap should be loaded
	if lazyBitmap.bitmap == nil {
		t.Error("Expected bitmap to be loaded after Draw()")
	}
}

func TestLazyBitmapDrawTwice(t *testing.T) {
	lazyBitmap := NewLazyBitmap("test.png")

	// Draw twice - should reuse loaded bitmap
	lazyBitmap.Draw()
	firstBitmap := lazyBitmap.bitmap

	lazyBitmap.Draw()
	secondBitmap := lazyBitmap.bitmap

	// Should be the same bitmap instance
	if firstBitmap != secondBitmap {
		t.Error("Expected LazyBitmap to reuse loaded bitmap")
	}
}

func TestLazyBitmapImplementsImage(t *testing.T) {
	lazyBitmap := NewLazyBitmap("test.png")

	// Verify LazyBitmap implements Image interface
	var image Image = lazyBitmap
	if image == nil {
		t.Error("Expected LazyBitmap to implement Image interface")
	}
}

func TestVirtualProxyDelaysLoading(t *testing.T) {
	// Create lazy bitmap - should not load yet
	lazyBitmap := NewLazyBitmap("demo.png")

	if lazyBitmap.bitmap != nil {
		t.Error("Expected loading to be deferred")
	}

	// Use it through DrawImage function
	DrawImage(lazyBitmap)

	// Now it should be loaded
	if lazyBitmap.bitmap == nil {
		t.Error("Expected bitmap to be loaded after DrawImage()")
	}
}

func TestBitmapEagerLoading(t *testing.T) {
	// Create bitmap directly - should load immediately
	bitmap := NewBitmap("demo.png")

	if bitmap == nil {
		t.Fatal("Expected bitmap to be created and loaded")
	}

	if bitmap.filename != "demo.png" {
		t.Errorf("Expected filename 'demo.png', got '%s'", bitmap.filename)
	}
}

func TestLazyBitmapFilenamePersistence(t *testing.T) {
	lazyBitmap := NewLazyBitmap("image.jpg")

	if lazyBitmap.filename != "image.jpg" {
		t.Errorf("Expected filename 'image.jpg', got '%s'", lazyBitmap.filename)
	}

	// Draw to load
	lazyBitmap.Draw()

	// Filename should still be preserved
	if lazyBitmap.filename != "image.jpg" {
		t.Error("Expected filename to be preserved after loading")
	}
}

func TestMultipleLazyBitmaps(t *testing.T) {
	lazy1 := NewLazyBitmap("image1.png")
	lazy2 := NewLazyBitmap("image2.png")

	if lazy1.bitmap != nil || lazy2.bitmap != nil {
		t.Error("Expected both bitmaps to be lazy-loaded")
	}

	lazy1.Draw()
	if lazy1.bitmap == nil {
		t.Error("Expected lazy1 bitmap to be loaded")
	}

	// lazy2 should still not be loaded
	if lazy2.bitmap != nil {
		t.Error("Expected lazy2 bitmap to still be lazy")
	}

	lazy2.Draw()
	if lazy2.bitmap == nil {
		t.Error("Expected lazy2 bitmap to be loaded after Draw()")
	}
}

func TestDrawImageFunction(t *testing.T) {
	lazyBitmap := NewLazyBitmap("test.png")

	// DrawImage should trigger loading
	DrawImage(lazyBitmap)

	// After DrawImage, bitmap should be loaded
	if lazyBitmap.bitmap == nil {
		t.Error("Expected bitmap to be loaded after DrawImage()")
	}
}
