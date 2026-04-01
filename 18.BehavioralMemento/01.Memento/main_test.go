package main

import (
	"testing"
)

func TestBankAccountMemento(t *testing.T) {
	ba := &BankAccount{balance: 100}

	if ba.balance != 100 {
		t.Errorf("Expected balance 100, got %d", ba.balance)
	}
}

func TestMementoDeposit(t *testing.T) {
	ba := &BankAccount{balance: 100}

	memento := ba.Deposit(50)

	if ba.balance != 150 {
		t.Errorf("Expected balance 150, got %d", ba.balance)
	}

	if memento.Balance != 150 {
		t.Errorf("Expected memento to capture balance 150, got %d", memento.Balance)
	}
}

func TestMementoRestore(t *testing.T) {
	ba := &BankAccount{balance: 100}
	m1 := ba.Deposit(50)

	if ba.balance != 150 {
		t.Errorf("Expected balance 150 before restore, got %d", ba.balance)
	}

	ba.Restore(m1)

	if ba.balance != 150 {
		t.Errorf("Expected balance 150 after restore to m1, got %d", ba.balance)
	}
}

func TestMementoMultipleSnapshots(t *testing.T) {
	ba := &BankAccount{balance: 100}

	m1 := ba.Deposit(50)  // balance = 150
	m2 := ba.Deposit(25)  // balance = 175

	if ba.balance != 175 {
		t.Errorf("Expected balance 175, got %d", ba.balance)
	}

	ba.Restore(m1)

	if ba.balance != 150 {
		t.Errorf("Expected balance 150 after restore to m1, got %d", ba.balance)
	}

	ba.Restore(m2)

	if ba.balance != 175 {
		t.Errorf("Expected balance 175 after restore to m2, got %d", ba.balance)
	}
}

func TestMementoCaptureState(t *testing.T) {
	initial := &BankAccount{balance: 200}
	snapshot := initial.Deposit(100)

	// Create new account and restore to snapshot
	recovered := &BankAccount{balance: 0}
	recovered.Restore(snapshot)

	if recovered.balance != 300 {
		t.Errorf("Expected recovered balance 300, got %d", recovered.balance)
	}
}

func TestMementoSequence(t *testing.T) {
	ba := &BankAccount{balance: 0}

	m1 := ba.Deposit(100)  // 100
	m2 := ba.Deposit(50)   // 150
	m3 := ba.Deposit(25)   // 175

	ba.Restore(m1)
	if ba.balance != 100 {
		t.Errorf("Expected balance 100, got %d", ba.balance)
	}

	ba.Restore(m3)
	if ba.balance != 175 {
		t.Errorf("Expected balance 175, got %d", ba.balance)
	}

	ba.Restore(m2)
	if ba.balance != 150 {
		t.Errorf("Expected balance 150, got %d", ba.balance)
	}
}

func TestMementoIndependent(t *testing.T) {
	ba := &BankAccount{balance: 100}

	m1 := ba.Deposit(50)  // m1 captures 150

	// Further changes after memento creation
	ba.Deposit(25)

	if ba.balance != 175 {
		t.Errorf("Expected balance 175, got %d", ba.balance)
	}

	// Restore to m1
	ba.Restore(m1)

	if ba.balance != 150 {
		t.Errorf("Expected memento to preserve state at capture time, balance 150, got %d", ba.balance)
	}
}

func TestMementoZeroBalance(t *testing.T) {
	ba := &BankAccount{balance: 100}
	_ = ba.Deposit(100) // 200
	m2 := ba.Deposit(-200) // 0

	ba.Restore(m2)

	if ba.balance != 0 {
		t.Errorf("Expected balance 0, got %d", ba.balance)
	}
}
