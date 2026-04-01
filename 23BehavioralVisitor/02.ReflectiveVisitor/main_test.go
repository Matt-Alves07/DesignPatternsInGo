package main

import (
	"strings"
	"testing"
)

func TestReflectiveVisitorDoubleExpression(t *testing.T) {
	expr := &DoubleExpression{value: 5}
	sb := &strings.Builder{}

	Print(expr, sb)

	output := sb.String()
	if !strings.Contains(output, "5") {
		t.Errorf("Expected 5 in output, got %s", output)
	}
}

func TestReflectiveVisitorAdditionExpression(t *testing.T) {
	expr := &AdditionExpression{
		left:  &DoubleExpression{3},
		right: &DoubleExpression{4},
	}
	sb := &strings.Builder{}

	Print(expr, sb)

	output := sb.String()
	if !strings.Contains(output, "(") || !strings.Contains(output, "+") {
		t.Errorf("Expected formatted output, got %s", output)
	}
}

func TestReflectiveVisitorNestedExpression(t *testing.T) {
	// 1+(2+3)
	expr := &AdditionExpression{
		left: &DoubleExpression{1},
		right: &AdditionExpression{
			left:  &DoubleExpression{2},
			right: &DoubleExpression{3},
		},
	}
	sb := &strings.Builder{}

	Print(expr, sb)

	output := sb.String()
	if output == "" {
		t.Error("Expected output for nested expression")
	}
}

func TestReflectiveVisitorTypeAssertion(t *testing.T) {
	var expr Expression

	expr = &DoubleExpression{10}
	if _, ok := expr.(*DoubleExpression); !ok {
		t.Error("Expected type assertion to work for DoubleExpression")
	}

	expr = &AdditionExpression{
		left:  &DoubleExpression{1},
		right: &DoubleExpression{2},
	}
	if _, ok := expr.(*AdditionExpression); !ok {
		t.Error("Expected type assertion to work for AdditionExpression")
	}
}

func TestReflectiveVisitorComplexTree(t *testing.T) {
	//   +
	//  / \
	// 1   +
	//    / \
	//   2   +
	//      / \
	//     3   4

	expr := &AdditionExpression{
		left: &DoubleExpression{1},
		right: &AdditionExpression{
			left: &DoubleExpression{2},
			right: &AdditionExpression{
				left:  &DoubleExpression{3},
				right: &DoubleExpression{4},
			},
		},
	}
	sb := &strings.Builder{}

	Print(expr, sb)

	output := sb.String()
	if output == "" {
		t.Error("Expected output for complex tree")
	}

	if !strings.Contains(output, "1") || !strings.Contains(output, "4") {
		t.Error("Expected all values in output")
	}
}

func TestReflectiveVisitorBreaksOCP(t *testing.T) {
	// Reflective visitor violates Open-Closed Principle
	// Because adding new expression types requires modifying Print function
	// We test the current implementation works

	expr := &AdditionExpression{
		left:  &DoubleExpression{10},
		right: &DoubleExpression{20},
	}
	sb := &strings.Builder{}

	Print(expr, sb)

	if sb.String() == "" {
		t.Error("Expected reflective visitor to work")
	}
}

func TestExpressionInterface(t *testing.T) {
	// Expression interface is empty in reflective visitor
	var expr Expression

	expr = &DoubleExpression{1}
	if expr == nil {
		t.Error("Expected DoubleExpression to implement Expression")
	}

	expr = &AdditionExpression{}
	if expr == nil {
		t.Error("Expected AdditionExpression to implement Expression")
	}
}
