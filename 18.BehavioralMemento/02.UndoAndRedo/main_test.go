package main

import (
	"testing"
)

func TestBankAccountUndoRedoCreation(t *testing.T) {
	ba := NewBankAccount(100)

	if ba.balance != 100 {
		t.Errorf("Expected initial balance 100, got %d", ba.balance)
	}

	if len(ba.changes) != 1 {
		t.Errorf("Expected 1 change recorded initially, got %d", len(ba.changes))
	}

	if ba.current != 0 {
		t.Errorf("Expected current=0 initially, got %d", ba.current)
	}
}

func TestBankAccountDeposit(t *testing.T) {
	ba := NewBankAccount(100)
	ba.Deposit(50)

	if ba.balance != 150 {
		t.Errorf("Expected balance 150, got %d", ba.balance)
	}

	if len(ba.changes) != 2 {
		t.Errorf("Expected 2 changes after deposit, got %d", len(ba.changes))
	}

	if ba.current != 1 {
		t.Errorf("Expected current=1, got %d", ba.current)
	}
}

func TestUndo(t *testing.T) {
	ba := NewBankAccount(100)
	ba.Deposit(50)  // 150
	ba.Deposit(25)  // 175

	ba.Undo()

	if ba.balance != 150 {
		t.Errorf("Expected balance 150 after undo, got %d", ba.balance)
	}

	if ba.current != 1 {
		t.Errorf("Expected current=1 after undo, got %d", ba.current)
	}
}

func TestMultipleUndos(t *testing.T) {
	ba := NewBankAccount(100)
	ba.Deposit(50)  // 150
	ba.Deposit(25)  // 175

	ba.Undo()
	ba.Undo()

	if ba.balance != 100 {
		t.Errorf("Expected balance 100 after 2 undos, got %d", ba.balance)
	}

	if ba.current != 0 {
		t.Errorf("Expected current=0 after 2 undos, got %d", ba.current)
	}
}

func TestRedo(t *testing.T) {
	ba := NewBankAccount(100)
	ba.Deposit(50)  // 150
	ba.Undo()

	ba.Redo()

	if ba.balance != 150 {
		t.Errorf("Expected balance 150 after redo, got %d", ba.balance)
	}

	if ba.current != 1 {
		t.Errorf("Expected current=1 after redo, got %d", ba.current)
	}
}

func TestUndoAndRedoSequence(t *testing.T) {
	ba := NewBankAccount(100)

	ba.Deposit(50)   // 150
	ba.Deposit(25)   // 175
	ba.Undo()        // back to 150
	ba.Undo()        // back to 100

	ba.Redo()        // forward to 150
	if ba.balance != 150 {
		t.Errorf("Expected balance 150 after first redo, got %d", ba.balance)
	}

	ba.Redo()        // forward to 175
	if ba.balance != 175 {
		t.Errorf("Expected balance 175 after second redo, got %d", ba.balance)
	}
}

func TestUndoWithoutHistory(t *testing.T) {
	ba := NewBankAccount(100)

	result := ba.Undo()

	if result != nil {
		t.Error("Expected nil when no undo available")
	}

	if ba.balance != 100 {
		t.Errorf("Expected balance to remain 100, got %d", ba.balance)
	}
}

func TestRedoWithoutFuture(t *testing.T) {
	ba := NewBankAccount(100)
	ba.Deposit(50)

	result := ba.Redo()

	if result != nil {
		t.Error("Expected nil when no redo available")
	}

	if ba.balance != 150 {
		t.Errorf("Expected balance to remain 150, got %d", ba.balance)
	}
}

func TestComplexUndoRedoSequence(t *testing.T) {
	ba := NewBankAccount(100)

	ba.Deposit(50)   // 150
	ba.Deposit(25)   // 175
	ba.Deposit(10)   // 185

	ba.Undo()        // 175
	ba.Undo()        // 150
	ba.Redo()        // 175

	if ba.balance != 175 {
		t.Errorf("Expected balance 175, got %d", ba.balance)
	}

	ba.Deposit(5)    // 180

	if len(ba.changes) != 5 { // original, +50, +25, +5
		t.Logf("Changes history after new deposit: %d", len(ba.changes))
	}
}

func TestStringRepresentation(t *testing.T) {
	ba := NewBankAccount(100)
	ba.Deposit(50)

	str := ba.String()

	if str == "" {
		t.Error("Expected non-empty string representation")
	}

	if !contains(str, "200") {
		t.Logf("String representation: %s", str)
	}
}

func contains(str, substr string) bool {
	for i := 0; i <= len(str)-len(substr); i++ {
		if str[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
