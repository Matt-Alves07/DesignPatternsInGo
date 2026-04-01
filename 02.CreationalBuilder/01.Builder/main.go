package main

import (
	"fmt"
	"strings"
)

const (
	indentSize = 2
)

// HtmlElement represents a single HTML element.
type HtmlElement struct {
	name, text string
	elements   []HtmlElement
}

// String returns the string representation of the HTML element.
func (e *HtmlElement) String() string {
	return e.string(0)
}

// string is a recursive helper to build the HTML string with indentation.
func (e *HtmlElement) string(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", indentSize*indent)
	sb.WriteString(fmt.Sprintf("%s<%s>\n", i, e.name))
	if len(e.text) > 0 {
		sb.WriteString(strings.Repeat(" ", indentSize*(indent+1)))
		sb.WriteString(e.text)
		sb.WriteString("\n")
	}

	for _, el := range e.elements {
		sb.WriteString(el.string(indent + 1))
	}
	sb.WriteString(fmt.Sprintf("%s</%s>\n", i, e.name))
	return sb.String()
}

// HtmlBuilder allows constructing HTML elements step-by-step.
type HtmlBuilder struct {
	rootName string
	root     HtmlElement
}

// NewHtmlBuilder creates a new HtmlBuilder with a root element name.
func NewHtmlBuilder(rootName string) *HtmlBuilder {
	b := HtmlBuilder{rootName,
		HtmlElement{rootName, "", []HtmlElement{}}}
	return &b
}

// String returns the string representation of the built HTML.
func (b *HtmlBuilder) String() string {
	return b.root.String()
}

// AddChild adds a child element to the root.
func (b *HtmlBuilder) AddChild(childName, childText string) {
	e := HtmlElement{childName, childText, []HtmlElement{}}
	b.root.elements = append(b.root.elements, e)
}

// AddChildFluent adds a child element and returns the builder for chaining.
func (b *HtmlBuilder) AddChildFluent(childName, childText string) *HtmlBuilder {
	e := HtmlElement{childName, childText, []HtmlElement{}}
	b.root.elements = append(b.root.elements, e)
	return b
}

func main() {
	// Simple string builder approach (non-structured)
	hello := "hello"
	sb := strings.Builder{}
	sb.WriteString("<p>")
	sb.WriteString(hello)
	sb.WriteString("</p>")
	fmt.Printf("%s\n", sb.String())

	// Example usage of simple string building
	words := []string{"hello", "world"}
	sb.Reset()
	sb.WriteString("<ul>")
	for _, v := range words {
		sb.WriteString("<li>")
		sb.WriteString(v)
		sb.WriteString("</li>")
	}
	sb.WriteString("</ul>")
	fmt.Println(sb.String())

	// Using the HtmlBuilder
	b := NewHtmlBuilder("ul")
	b.AddChildFluent("li", "hello").
		AddChildFluent("li", "world")
	fmt.Println(b.String())
}