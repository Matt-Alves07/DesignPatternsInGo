package main

import (
	"testing"
)

func TestColorSpecification(t *testing.T) {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}
	products := []Product{apple, tree, house}

	bf := BetterFilter{}
	greenSpec := ColorSpecification{green}

	filtered := bf.Filter(products, greenSpec)
	if len(filtered) != 2 {
		t.Errorf("Expected 2 green products, got %d", len(filtered))
	}

	for _, p := range filtered {
		if p.color != green {
			t.Errorf("Expected green product, got %v", p.color)
		}
	}
}

func TestSizeSpecification(t *testing.T) {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}
	products := []Product{apple, tree, house}

	bf := BetterFilter{}
	largeSpec := SizeSpecification{large}

	filtered := bf.Filter(products, largeSpec)
	if len(filtered) != 2 {
		t.Errorf("Expected 2 large products, got %d", len(filtered))
	}

	for _, p := range filtered {
		if p.size != large {
			t.Errorf("Expected large product, got %v", p.size)
		}
	}
}

func TestAndSpecification(t *testing.T) {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}
	products := []Product{apple, tree, house}

	bf := BetterFilter{}
	largeSpec := SizeSpecification{large}
	greenSpec := ColorSpecification{green}
	largeGreenSpec := AndSpecification{largeSpec, greenSpec}

	filtered := bf.Filter(products, largeGreenSpec)
	if len(filtered) != 1 {
		t.Errorf("Expected 1 large green product, got %d", len(filtered))
	}

	if filtered[0].name != "Tree" {
		t.Errorf("Expected Tree, got %s", filtered[0].name)
	}
}
