package main

// ISP (Interface Segregation Principle): Clients should not be forced to depend on methods they do not use.
// It's better to have many specific interfaces than one general-purpose interface.

// Document is a placeholder for a document.
type Document struct{}

// Machine is a "fat" interface that forces implementers to define methods they might not support.
// This violates ISP.
type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

// MultiFunctionPrinter implements Machine, supporting all operations.
type MultiFunctionPrinter struct{}

func (m MultiFunctionPrinter) Print(d Document) {
	// Implementation for printing
}

func (m MultiFunctionPrinter) Fax(d Document) {
	// Implementation for faxing
}

func (m MultiFunctionPrinter) Scan(d Document) {
	// Implementation for scanning
}

// OldFashionedPrinter only supports printing.
// Forced to implement Fax and Scan because of the Machine interface, which violates ISP.
type OldFashionedPrinter struct{}

func (o OldFashionedPrinter) Print(d Document) {
	// Implementation for printing
}

// Fax panics because OldFashionedPrinter cannot fax.
func (o OldFashionedPrinter) Fax(d Document) {
	panic("operation not supported")
}

// Scan panics because OldFashionedPrinter cannot scan.
func (o OldFashionedPrinter) Scan(d Document) {
	panic("operation not supported")
}

// Printer is a segregated interface for printing.
type Printer interface {
	Print(d Document)
}

// Scanner is a segregated interface for scanning.
type Scanner interface {
	Scan(d Document)
}

// MyPrinter implements only the Printer interface, adhering to ISP.
type MyPrinter struct{}

func (m MyPrinter) Print(d Document) {
	// Implementation for printing
}

// Photocopier implements both Printer and Scanner interfaces.
type Photocopier struct{}

func (p Photocopier) Scan(d Document) {
	// Implementation for scanning
}

func (p Photocopier) Print(d Document) {
	// Implementation for printing
}

// MultiFunctionDevice composes Printer and Scanner interfaces.
type MultiFunctionDevice interface {
	Printer
	Scanner
}

// MultiFunctionMachine uses composition/decoration to combine capabilities.
type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
}

func (m MultiFunctionMachine) Print(d Document) {
	m.printer.Print(d)
}

func (m MultiFunctionMachine) Scan(d Document) {
	m.scanner.Scan(d)
}

func main() {
	// Example usage
	ofp := OldFashionedPrinter{}
	ofp.Print(Document{})
	// ofp.Scan(Document{}) // This would panic

	p := MyPrinter{}
	p.Print(Document{})
}