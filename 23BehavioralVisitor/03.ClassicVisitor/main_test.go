package main

import (
	"strings"
	"testing"
)

func TestExpressionVisitorInterface(t *testing.T) {
	var visitor ExpressionVisitor

	visitor = NewExpressionPrinter()

	if visitor == nil {
		t.Error("Expected ExpressionPrinter to implement ExpressionVisitor")
	}
}

func TestExpressionPrinterCreation(t *testing.T) {
	printer := NewExpressionPrinter()

	if printer == nil {
		t.Fatal("Expected ExpressionPrinter to be created")
	}

	if printer.String() != "" {
		t.Error("Expected empty printer initially")
	}
}

func TestDoubleExpressionAccept(t *testing.T) {
	expr := &DoubleExpression{value: 42}
	printer := NewExpressionPrinter()

	expr.Accept(printer)

	output := printer.String()
	if !strings.Contains(output, "42") {
		t.Errorf("Expected 42 in output, got %s", output)
	}
}

func TestAdditionExpressionAccept(t *testing.T) {
	expr := &AdditionExpression{
		left:  &DoubleExpression{1},
		right: &DoubleExpression{2},
	}
	printer := NewExpressionPrinter()

	expr.Accept(printer)

	output := printer.String()
	if !strings.Contains(output, "(") || !strings.Contains(output, "+") {
		t.Errorf("Expected formatted addition, got %s", output)
	}
}

func TestNestedExpressionAccept(t *testing.T) {
	// 1+(2+3)
	expr := &AdditionExpression{
		left: &DoubleExpression{1},
		right: &AdditionExpression{
			left:  &DoubleExpression{2},
			right: &DoubleExpression{3},
		},
	}
	printer := NewExpressionPrinter()

	expr.Accept(printer)

	output := printer.String()
	if output == "" {
		t.Error("Expected output for nested expression")
	}

	// Should have proper formatting
	if !strings.Contains(output, "(") {
		t.Error("Expected parentheses for nested expressions")
	}
}

func TestVisitDoubleExpression(t *testing.T) {
	printer := NewExpressionPrinter()
	expr := &DoubleExpression{value: 3.14}

	printer.VisitDoubleExpression(expr)

	output := printer.String()
	if output == "" {
		t.Error("Expected output from visiting double expression")
	}
}

func TestVisitAdditionExpression(t *testing.T) {
	printer := NewExpressionPrinter()
	expr := &AdditionExpression{
		left:  &DoubleExpression{5},
		right: &DoubleExpression{10},
	}

	printer.VisitAdditionExpression(expr)

	output := printer.String()
	if output == "" {
		t.Error("Expected output from visiting addition expression")
	}
}

func TestClassicVisitorPattern(t *testing.T) {
	// Classic visitor pattern: separate visitor from expressions
	// Allows adding new operations without modifying expression classes

	expr := &AdditionExpression{
		left:  &DoubleExpression{100},
		right: &DoubleExpression{200},
	}

	printer := NewExpressionPrinter()
	expr.Accept(printer)

	output := printer.String()
	if !strings.Contains(output, "100") || !strings.Contains(output, "200") {
		t.Error("Expected both values in output")
	}
}

func TestComplexExpressionVisitor(t *testing.T) {
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

	printer := NewExpressionPrinter()
	expr.Accept(printer)

	output := printer.String()
	if len(output) == 0 {
		t.Error("Expected output for complex expression")
	}

	// All values should be present
	if !strings.Contains(output, "1") {
		t.Error("Expected 1 in output")
	}
}

func TestExpressionAcceptMethod(t *testing.T) {
	// Both expression types should implement Expression interface
	var expr Expression

	expr = &DoubleExpression{1}
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

	// Both should have Accept method
	printer := NewExpressionPrinter()
	expr.Accept(printer)
}
