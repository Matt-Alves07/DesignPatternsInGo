package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

var entryCount = 0

// Journal represents a diary aimed at storing string entries.
// It complies with the Single Responsibility Principle by only handling
// the management of entries.
type Journal struct {
	entries []string
}

// String returns the string representation of the journal entries.
func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

// AddEntry adds a new text entry to the journal and returns the entry count.
func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s",
		entryCount,
		text)
	j.entries = append(j.entries, entry)
	return entryCount
}

// RemoveEntry removes an entry at a given index.
func (j *Journal) RemoveEntry(index int) {
	// ... implementation omitted
}

// Save is a method that breaks the Single Responsibility Principle (SRP).
// The Journal struct should not be responsible for persistence.
func (j *Journal) Save(filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(j.String()), 0644)
}

// Load is a method that breaks the SRP.
func (j *Journal) Load(filename string) {
	// ... implementation omitted
}

// LoadFromWeb is a method that breaks the SRP.
func (j *Journal) LoadFromWeb(url *url.URL) {
	// ... implementation omitted
}

var lineSeparator = "\n"

// SaveToFile saves the journal to a file.
// This function handles persistence, adhering to SRP by separating concerns.
func SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(strings.Join(j.entries, lineSeparator)), 0644)
}

// Persistence handles the saving of objects to files.
// It acts as a separate concern from the Journal itself.
type Persistence struct {
	lineSeparator string
}

// saveToFile saves the journal to a file using the Persistence configuration.
func (p *Persistence) saveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(strings.Join(j.entries, p.lineSeparator)), 0644)
}

func main() {
	j := Journal{}
	j.AddEntry("I cried today.")
	j.AddEntry("I ate a bug")
	fmt.Println(strings.Join(j.entries, "\n"))

	// Save using a separate function (better for SRP)
	SaveToFile(&j, "journal.txt")

	// Save using a Persistence struct (also follows SRP)
	p := Persistence{"\n"}
	p.saveToFile(&j, "journal.txt")
}