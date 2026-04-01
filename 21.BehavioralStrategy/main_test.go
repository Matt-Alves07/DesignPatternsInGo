package main

import (
	"strings"
	"testing"
)

func TestTextProcessorCreation(t *testing.T) {
	tp := NewTextProcessor(&MarkdownListStrategy{})

	if tp == nil {
		t.Fatal("Expected TextProcessor to be created")
	}
}

func TestMarkdownStrategy(t *testing.T) {
	strategy := &MarkdownListStrategy{}

	builder := &strings.Builder{}
	strategy.Start(builder)
	strategy.AddListItem(builder, "Item 1")
	strategy.AddListItem(builder, "Item 2")
	strategy.End(builder)

	output := builder.String()

	if !strings.Contains(output, "* Item 1") {
		t.Error("Expected markdown bullet points")
	}
}

func TestHtmlStrategy(t *testing.T) {
	strategy := &HtmlListStrategy{}

	builder := &strings.Builder{}
	strategy.Start(builder)
	strategy.AddListItem(builder, "Item 1")
	strategy.End(builder)

	output := builder.String()

	if !strings.Contains(output, "<ul>") {
		t.Error("Expected HTML ul tag")
	}

	if !strings.Contains(output, "<li>Item 1</li>") {
		t.Error("Expected HTML li tag")
	}
}

func TestSetOutputFormat(t *testing.T) {
	tp := NewTextProcessor(&MarkdownListStrategy{})

	tp.SetOutputFormat(Html)

	_, isHtml := tp.listStrategy.(*HtmlListStrategy)
	if !isHtml {
		t.Error("Expected format to change to HTML")
	}
}

func TestAppendListMarkdown(t *testing.T) {
	tp := NewTextProcessor(&MarkdownListStrategy{})

	items := []string{"First", "Second", "Third"}
	tp.AppendList(items)

	output := tp.String()

	for _, item := range items {
		if !strings.Contains(output, item) {
			t.Errorf("Expected item '%s' in output", item)
		}
	}

	if !strings.Contains(output, "*") {
		t.Error("Expected markdown bullet points")
	}
}

func TestAppendListHtml(t *testing.T) {
	tp := NewTextProcessor(&HtmlListStrategy{})

	items := []string{"Apple", "Banana"}
	tp.AppendList(items)

	output := tp.String()

	if !strings.Contains(output, "<ul>") || !strings.Contains(output, "</ul>") {
		t.Error("Expected HTML list tags")
	}

	for _, item := range items {
		liTag := "<li>" + item + "</li>"
		if !strings.Contains(output, liTag) {
			t.Errorf("Expected HTML li tag for '%s'", item)
		}
	}
}

func TestReset(t *testing.T) {
	tp := NewTextProcessor(&MarkdownListStrategy{})

	tp.AppendList([]string{"Item"})

	if tp.String() == "" {
		t.Error("Expected content after AppendList")
	}

	tp.Reset()

	if tp.String() != "" {
		t.Error("Expected empty content after Reset")
	}
}

func TestStrategySwitch(t *testing.T) {
	tp := NewTextProcessor(&MarkdownListStrategy{})

	tp.AppendList([]string{"Item 1"})
	markdownOutput := tp.String()

	tp.Reset()
	tp.SetOutputFormat(Html)
	tp.AppendList([]string{"Item 1"})
	htmlOutput := tp.String()

	if markdownOutput == htmlOutput {
		t.Error("Expected different output for different strategies")
	}

	if !strings.Contains(markdownOutput, "*") {
		t.Error("Expected markdown format in first output")
	}

	if !strings.Contains(htmlOutput, "<") {
		t.Error("Expected HTML format in second output")
	}
}

func TestEmptyListMarkdown(t *testing.T) {
	tp := NewTextProcessor(&MarkdownListStrategy{})

	tp.AppendList([]string{})

	output := tp.String()

	if output != "" && !strings.Contains(output, "* ") == false {
		t.Log("Empty list markdown: " + output)
	}
}

func TestEmptyListHtml(t *testing.T) {
	tp := NewTextProcessor(&HtmlListStrategy{})

	tp.AppendList([]string{})

	output := tp.String()

	if !strings.Contains(output, "<ul>") || !strings.Contains(output, "</ul>") {
		t.Error("Expected HTML tags even for empty list")
	}
}
