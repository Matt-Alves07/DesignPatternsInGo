package main

import (
	"testing"
)

func TestPersonCreation(t *testing.T) {
	person := &Person{"Alexander", "Graham", "Bell"}

	if person.FirstName != "Alexander" {
		t.Errorf("Expected FirstName 'Alexander', got %s", person.FirstName)
	}

	if person.MiddleName != "Graham" {
		t.Errorf("Expected MiddleName 'Graham', got %s", person.MiddleName)
	}

	if person.LastName != "Bell" {
		t.Errorf("Expected LastName 'Bell', got %s", person.LastName)
	}
}

func TestPersonNames(t *testing.T) {
	person := &Person{"Alexander", "Graham", "Bell"}
	names := person.Names()

	if len(names) != 3 {
		t.Errorf("Expected 3 names, got %d", len(names))
	}

	if names[0] != "Alexander" {
		t.Errorf("Expected names[0]='Alexander', got '%s'", names[0])
	}

	if names[2] != "Bell" {
		t.Errorf("Expected names[2]='Bell', got '%s'", names[2])
	}
}

func TestNamesGeneratorChannel(t *testing.T) {
	person := &Person{"Alexander", "Graham", "Bell"}
	ch := person.NamesGenerator()

	name1 := <-ch
	if name1 != "Alexander" {
		t.Errorf("Expected 'Alexander', got '%s'", name1)
	}

	name2 := <-ch
	if name2 != "Graham" {
		t.Errorf("Expected 'Graham', got '%s'", name2)
	}

	name3 := <-ch
	if name3 != "Bell" {
		t.Errorf("Expected 'Bell', got '%s'", name3)
	}

	_, ok := <-ch
	if ok {
		t.Error("Expected channel to be closed")
	}
}

func TestNamesGeneratorWithoutMiddleName(t *testing.T) {
	person := &Person{"Alexander", "", "Bell"}
	ch := person.NamesGenerator()

	name1 := <-ch
	if name1 != "Alexander" {
		t.Errorf("Expected 'Alexander', got '%s'", name1)
	}

	name2 := <-ch
	if name2 != "Bell" {
		t.Errorf("Expected 'Bell', got '%s'", name2)
	}

	_, ok := <-ch
	if ok {
		t.Error("Expected channel to be closed after 2 names")
	}
}

func TestPersonNameIterator(t *testing.T) {
	person := &Person{"Alexander", "Graham", "Bell"}
	iterator := NewPersonNameIterator(person)

	if iterator.person != person {
		t.Error("Expected iterator to reference person")
	}

	if iterator.current != -1 {
		t.Errorf("Expected initial current=-1, got %d", iterator.current)
	}
}

func TestIteratorMoveNext(t *testing.T) {
	person := &Person{"Alexander", "Graham", "Bell"}
	iterator := NewPersonNameIterator(person)

	if !iterator.MoveNext() {
		t.Error("Expected first MoveNext to succeed")
	}

	if iterator.current != 0 {
		t.Errorf("Expected current=0, got %d", iterator.current)
	}
}

func TestIteratorTraversal(t *testing.T) {
	person := &Person{"Alexander", "Graham", "Bell"}
	iterator := NewPersonNameIterator(person)

	names := []string{}
	for iterator.MoveNext() {
		names = append(names, iterator.Current())
	}

	if len(names) != 3 {
		t.Errorf("Expected 3 names, got %d", len(names))
	}

	if names[0] != "Alexander" || names[1] != "Graham" || names[2] != "Bell" {
		t.Errorf("Expected names [Alexander Graham Bell], got %v", names)
	}
}

func TestIteratorWithoutMiddleName(t *testing.T) {
	person := &Person{"John", "", "Doe"}
	iterator := NewPersonNameIterator(person)

	names := []string{}
	for iterator.MoveNext() {
		names = append(names, iterator.Current())
	}

	// Should still iterate through all positions (including empty middle name)
	if len(names) != 3 {
		t.Errorf("Expected 3 iterations, got %d", len(names))
	}
}

func TestMultipleIterators(t *testing.T) {
	person := &Person{"Jane", "Mary", "Smith"}

	iter1 := NewPersonNameIterator(person)
	iter2 := NewPersonNameIterator(person)

	iter1.MoveNext()
	iter2.MoveNext()
	iter2.MoveNext()

	name1 := iter1.Current()
	name2 := iter2.Current()

	if name1 == name2 {
		t.Error("Expected different positions for different iterators")
	}
}

func TestIteratorBoundary(t *testing.T) {
	person := &Person{"Alice", "B", "Brown"}
	iterator := NewPersonNameIterator(person)

	// Move through all names
	iterator.MoveNext()
	iterator.MoveNext()
	iterator.MoveNext()

	// Should not move further
	if iterator.MoveNext() {
		t.Error("Expected MoveNext to fail at end")
	}
}
