package main

import (
	"strings"
	"testing"
)

func TestHtmlBuilder(t *testing.T) {
	b := NewHtmlBuilder("ul")
	b.AddChild("li", "hello")
	b.AddChild("li", "world")

	result := b.String()
	expectedContains := []string{
		"<ul>",
		"  <li>",
		"    hello",
		"  </li>",
		"  <li>",
		"    world",
		"  </li>",
		"</ul>",
	}

	for _, exp := range expectedContains {
		if !strings.Contains(result, exp) {
			t.Errorf("Builder output missing expected part: %s. Got:\n%s", exp, result)
		}
	}
}

func TestHtmlBuilder_Fluent(t *testing.T) {
	b := NewHtmlBuilder("ul")
	b.AddChildFluent("li", "fluent").
		AddChildFluent("li", "chain")

	result := b.String()
	
	if !strings.Contains(result, "fluent") {
		t.Error("Missing 'fluent' in output")
	}
	if !strings.Contains(result, "chain") {
		t.Error("Missing 'chain' in output")
	}
}
