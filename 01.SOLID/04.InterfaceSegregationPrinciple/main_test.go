package main

import "testing"

func TestOldFashionedPrinter_Print(t *testing.T) {
	ofp := OldFashionedPrinter{}
	// Should not panic
	ofp.Print(Document{})
}

func TestOldFashionedPrinter_Scan_Panic(t *testing.T) {
	ofp := OldFashionedPrinter{}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	ofp.Scan(Document{})
}

func TestMyPrinter(t *testing.T) {
	p := MyPrinter{} // implements only Printer
	var _ Printer = p
	// var _ Scanner = p // compilation error if uncommented, proving segregation
}

func TestPhotocopier(t *testing.T) {
	p := Photocopier{}
	var _ Printer = p
	var _ Scanner = p
	var _ MultiFunctionDevice = p
}

type spyPrinter struct {
	printed bool
}

func (s *spyPrinter) Print(d Document) {
	s.printed = true
}

func TestMultiFunctionMachine(t *testing.T) {
	spy := &spyPrinter{}
	mfm := MultiFunctionMachine{
		printer: spy,
		scanner: nil, 
	}
	mfm.Print(Document{})

	if !spy.printed {
		t.Error("MultiFunctionMachine did not delegate Print call")
	}
}
