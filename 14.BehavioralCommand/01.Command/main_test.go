package main

import (
	"testing"
)

func TestBankAccountCreation(t *testing.T) {
	ba := &BankAccount{}

	if ba.balance != 0 {
		t.Errorf("Expected initial balance 0, got %d", ba.balance)
	}
}

func TestBankAccountDeposit(t *testing.T) {
	ba := &BankAccount{}
	ba.Deposit(100)

	if ba.balance != 100 {
		t.Errorf("Expected balance 100 after deposit, got %d", ba.balance)
	}
}

func TestBankAccountWithdraw(t *testing.T) {
	ba := &BankAccount{balance: 200}

	ok := ba.Withdraw(50)

	if !ok {
		t.Error("Expected withdrawal to succeed")
	}

	if ba.balance != 150 {
		t.Errorf("Expected balance 150 after withdrawal, got %d", ba.balance)
	}
}

func TestBankAccountWithdrawInsufficient(t *testing.T) {
	ba := &BankAccount{balance: 20}

	// Overdraft limit is -500, so we need to exceed that limit
	// 20 - 525 = -505 which exceeds the -500 limit
	ok := ba.Withdraw(525)

	if ok {
		t.Error("Expected withdrawal to fail when exceeding overdraft limit")
	}

	if ba.balance != 20 {
		t.Errorf("Expected balance to remain 20, got %d", ba.balance)
	}
}

func TestBankAccountCommand(t *testing.T) {
	ba := &BankAccount{}
	cmd := NewBankAccountCommand(ba, Deposit, 100)

	cmd.Call()

	if ba.balance != 100 {
		t.Errorf("Expected balance 100 after command, got %d", ba.balance)
	}
}

func TestBankAccountCommandUndo(t *testing.T) {
	ba := &BankAccount{balance: 100}
	cmd := NewBankAccountCommand(ba, Withdraw, 50)

	cmd.Call()
	if ba.balance != 50 {
		t.Errorf("Expected balance 50 after withdrawal, got %d", ba.balance)
	}

	cmd.Undo()
	if ba.balance != 100 {
		t.Errorf("Expected balance 100 after undo, got %d", ba.balance)
	}
}

func TestBankAccountCommandSequence(t *testing.T) {
	ba := &BankAccount{}

	cmd1 := NewBankAccountCommand(ba, Deposit, 100)
	cmd1.Call()

	cmd2 := NewBankAccountCommand(ba, Withdraw, 50)
	cmd2.Call()

	if ba.balance != 50 {
		t.Errorf("Expected balance 50, got %d", ba.balance)
	}

	cmd2.Undo()
	if ba.balance != 100 {
		t.Errorf("Expected balance 100 after undo, got %d", ba.balance)
	}

	cmd1.Undo()
	if ba.balance != 0 {
		t.Errorf("Expected balance 0 after undo, got %d", ba.balance)
	}
}

func TestBankAccountCommandOverdraft(t *testing.T) {
	ba := &BankAccount{}
	cmd := NewBankAccountCommand(ba, Withdraw, 100)

	cmd.Call()

	if ba.balance > 100 { // Should be within overdraft limit
		t.Errorf("Expected balance to be within overdraft limit, got %d", ba.balance)
	}
}

func TestCommandUndoFailedCommand(t *testing.T) {
	ba := &BankAccount{balance: 10}
	cmd := NewBankAccountCommand(ba, Withdraw, 600)

	cmd.Call()

	// Command should fail because 10 - 600 = -610 exceeds overdraft limit of -500
	if ba.balance != 10 {
		t.Errorf("Expected balance to remain 10 after failed withdrawal, got %d", ba.balance)
	}

	// Undo on a failed command - balance should not change
	cmd.Undo()

	if ba.balance != 10 {
		t.Errorf("Expected balance to remain 10 after undo of failed command, got %d", ba.balance)
	}
}
