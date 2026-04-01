package main

import (
	"strings"
	"testing"
)

func TestDoubleExpressionPrint(t *testing.T) {
	expr := &DoubleExpression{value: 3.14}
	sb := &strings.Builder{}

	expr.Print(sb)

	output := sb.String()
	if !strings.Contains(output, "3.14") {
		t.Errorf("Expected 3.14 in output, got %s", output)
	}
}

func TestAdditionExpressionPrint(t *testing.T) {
	expr := &AdditionExpression{
		left:  &DoubleExpression{1},
		right: &DoubleExpression{2},
	}
	sb := &strings.Builder{}

	expr.Print(sb)

	output := sb.String()
	if !strings.Contains(output, "(") || !strings.Contains(output, "+") {
		t.Errorf("Expected format (1+2), got %s", output)
	}
}

func TestNestedExpression(t *testing.T) {
	// 1+(2+3)
	expr := &AdditionExpression{
		left: &DoubleExpression{1},
		right: &AdditionExpression{
			left:  &DoubleExpression{2},
			right: &DoubleExpression{3},
		},
	}
	sb := &strings.Builder{}

	expr.Print(sb)

	output := sb.String()
	if !strings.Contains(output, "(") {
		t.Error("Expected nested parentheses")
	}
}

func TestExpressionInterface(t *testing.T) {
	var expr Expression

	expr = &DoubleExpression{5}
	if expr == nil {
		t.Error("Expected DoubleExpression to implement Expression")
	}

	expr = &AdditionExpression{
		left:  &DoubleExpression{1},
		right: &DoubleExpression{2},
	}
	if expr == nil {
		t.Error("Expected AdditionExpression to implement Expression")
	}
}

func TestComplexExpression(t *testing.T) {
	// ((1+2)+(3+4))
	expr := &AdditionExpression{
		left: &AdditionExpression{
			left:  &DoubleExpression{1},
			right: &DoubleExpression{2},
		},
		right: &AdditionExpression{
			left:  &DoubleExpression{3},
			right: &DoubleExpression{4},
		},
	}
	sb := &strings.Builder{}

	expr.Print(sb)

	output := sb.String()
	if output == "" {
		t.Error("Expected output for complex expression")
	}

	if !strings.Contains(output, "(") || !strings.Contains(output, "+") {
		t.Error("Expected proper formatting")
	}
}

func TestDoubleExpressionValue(t *testing.T) {
	values := []float64{0, 1, -5, 3.14, 100.99}

	for _, val := range values {
		expr := &DoubleExpression{value: val}
		sb := &strings.Builder{}

		expr.Print(sb)

		output := sb.String()
		if output == "" {
			t.Errorf("Expected output for value %v", val)
		}
	}
}

func TestIntrusiveVisitorPattern(t *testing.T) {
	// Intrusive visitor adds methods to expression node
	expr := &DoubleExpression{42}

	sb := &strings.Builder{}
	expr.Print(sb)

	if sb.String() == "" {
		t.Error("Expected Print method to work on DoubleExpression")
	}
}
