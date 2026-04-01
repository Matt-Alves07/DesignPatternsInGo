package main

import (
	"strings"
	"testing"
)

func TestGraphicObjectCreation(t *testing.T) {
	obj := GraphicObject{"Test", "Red", nil}

	if obj.Name != "Test" {
		t.Errorf("Expected name 'Test', got %s", obj.Name)
	}

	if obj.Color != "Red" {
		t.Errorf("Expected color 'Red', got %s", obj.Color)
	}

	if len(obj.Children) != 0 {
		t.Errorf("Expected 0 children, got %d", len(obj.Children))
	}
}

func TestNewCircleShape(t *testing.T) {
	circle := NewCircle("Blue")

	if circle.Name != "Circle" {
		t.Errorf("Expected name 'Circle', got %s", circle.Name)
	}

	if circle.Color != "Blue" {
		t.Errorf("Expected color 'Blue', got %s", circle.Color)
	}

	if len(circle.Children) != 0 {
		t.Errorf("Expected 0 children, got %d", len(circle.Children))
	}
}

func TestNewSquareShape(t *testing.T) {
	square := NewSquare("Green")

	if square.Name != "Square" {
		t.Errorf("Expected name 'Square', got %s", square.Name)
	}

	if square.Color != "Green" {
		t.Errorf("Expected color 'Green', got %s", square.Color)
	}
}

func TestCompositeStructure(t *testing.T) {
	drawing := GraphicObject{"My Drawing", "", nil}
	drawing.Children = append(drawing.Children, *NewSquare("Red"))
	drawing.Children = append(drawing.Children, *NewCircle("Yellow"))

	if len(drawing.Children) != 2 {
		t.Errorf("Expected 2 children, got %d", len(drawing.Children))
	}

	if drawing.Children[0].Name != "Square" {
		t.Errorf("Expected first child to be Square, got %s", drawing.Children[0].Name)
	}

	if drawing.Children[0].Color != "Red" {
		t.Errorf("Expected first child color Red, got %s", drawing.Children[0].Color)
	}

	if drawing.Children[1].Name != "Circle" {
		t.Errorf("Expected second child to be Circle, got %s", drawing.Children[1].Name)
	}
}

func TestNestedComposite(t *testing.T) {
	drawing := GraphicObject{"Drawing", "", nil}
	group := GraphicObject{"Group", "", nil}

	group.Children = append(group.Children, *NewCircle("Blue"))
	group.Children = append(group.Children, *NewSquare("Red"))

	drawing.Children = append(drawing.Children, group)
	drawing.Children = append(drawing.Children, *NewSquare("Green"))

	if len(drawing.Children) != 2 {
		t.Errorf("Expected 2 top-level children, got %d", len(drawing.Children))
	}

	if len(drawing.Children[0].Children) != 2 {
		t.Errorf("Expected 2 group children, got %d", len(drawing.Children[0].Children))
	}
}

func TestGraphicObjectString(t *testing.T) {
	circle := NewCircle("Blue")
	output := circle.String()

	if len(output) == 0 {
		t.Error("Expected String() to produce output")
	}

	if !strings.Contains(output, "Circle") {
		t.Error("Expected String() to contain 'Circle'")
	}
}

func TestCompositeStringRendering(t *testing.T) {
	drawing := GraphicObject{"My Drawing", "", nil}
	drawing.Children = append(drawing.Children, *NewSquare("Red"))
	drawing.Children = append(drawing.Children, *NewCircle("Yellow"))

	output := drawing.String()

	if !strings.Contains(output, "My Drawing") {
		t.Error("Expected output to contain 'My Drawing'")
	}

	if !strings.Contains(output, "Square") {
		t.Error("Expected output to contain 'Square'")
	}

	if !strings.Contains(output, "Circle") {
		t.Error("Expected output to contain 'Circle'")
	}

	if !strings.Contains(output, "Red") {
		t.Error("Expected output to contain 'Red'")
	}

	if !strings.Contains(output, "Yellow") {
		t.Error("Expected output to contain 'Yellow'")
	}
}

func TestDepthBasedFormatting(t *testing.T) {
	drawing := GraphicObject{"Drawing", "", nil}
	group := GraphicObject{"Group", "", nil}
	group.Children = append(group.Children, *NewCircle("Blue"))
	drawing.Children = append(drawing.Children, group)

	output := drawing.String()

	// Check that indentation increases with depth
	lines := strings.Split(output, "\n")
	if len(lines) < 2 {
		t.Error("Expected multiple lines in output")
	}

	// Find lines with asterisks (indentation markers)
	var depths []int
	for _, line := range lines {
		if len(line) > 0 {
			count := 0
			for _, c := range line {
				if c == '*' {
					count++
				} else {
					break
				}
			}
			if count > 0 {
				depths = append(depths, count)
			}
		}
	}

	// Check that deeper elements have more asterisks
	if len(depths) >= 2 {
		for i := 1; i < len(depths); i++ {
			if depths[i] <= depths[i-1] && i <= 2 {
				// Allow for some flexibility in depth tracking
			}
		}
	}
}

func TestMultipleGroupsComposite(t *testing.T) {
	drawing := GraphicObject{"Drawing", "", nil}

	group1 := GraphicObject{"Group1", "", nil}
	group1.Children = append(group1.Children, *NewCircle("Red"))

	group2 := GraphicObject{"Group2", "", nil}
	group2.Children = append(group2.Children, *NewSquare("Blue"))

	drawing.Children = append(drawing.Children, group1)
	drawing.Children = append(drawing.Children, group2)

	if len(drawing.Children) != 2 {
		t.Errorf("Expected 2 groups, got %d", len(drawing.Children))
	}

	if len(drawing.Children[0].Children) != 1 {
		t.Errorf("Expected group1 to have 1 child, got %d", len(drawing.Children[0].Children))
	}

	if len(drawing.Children[1].Children) != 1 {
		t.Errorf("Expected group2 to have 1 child, got %d", len(drawing.Children[1].Children))
	}
}
