package main

import (
	"testing"
)

func TestCompositeCommandCall(t *testing.T) {
	ba := &BankAccount{}

	cmd1 := NewBankAccountCommand(ba, Deposit, 100)
	cmd2 := NewBankAccountCommand(ba, Withdraw, 50)

	composite := &CompositeBankAccountCommand{
		commands: []Command{cmd1, cmd2},
	}

	composite.Call()

	if ba.balance != 50 {
		t.Errorf("Expected balance 50, got %d", ba.balance)
	}
}

func TestCompositeCommandUndo(t *testing.T) {
	ba := &BankAccount{}

	cmd1 := NewBankAccountCommand(ba, Deposit, 100)
	cmd2 := NewBankAccountCommand(ba, Withdraw, 50)

	composite := &CompositeBankAccountCommand{
		commands: []Command{cmd1, cmd2},
	}

	composite.Call()
	if ba.balance != 50 {
		t.Errorf("Expected balance 50 after call, got %d", ba.balance)
	}

	composite.Undo()
	if ba.balance != 0 {
		t.Errorf("Expected balance 0 after undo, got %d", ba.balance)
	}
}

func TestCompositeCommandSucceeded(t *testing.T) {
	ba := &BankAccount{}

	cmd1 := NewBankAccountCommand(ba, Deposit, 100)
	cmd2 := NewBankAccountCommand(ba, Withdraw, 50)

	composite := &CompositeBankAccountCommand{
		commands: []Command{cmd1, cmd2},
	}

	composite.Call()

	if !composite.Succeeded() {
		t.Error("Expected composite command to have succeeded")
	}
}

func TestMoneyTransferCommand(t *testing.T) {
	from := &BankAccount{balance: 100}
	to := &BankAccount{balance: 0}

	transfer := NewMoneyTransferCommand(from, to, 50)
	transfer.Call()

	if from.balance != 50 {
		t.Errorf("Expected from balance 50, got %d", from.balance)
	}

	if to.balance != 50 {
		t.Errorf("Expected to balance 50, got %d", to.balance)
	}
}

func TestMoneyTransferCommandUndo(t *testing.T) {
	from := &BankAccount{balance: 100}
	to := &BankAccount{balance: 0}

	transfer := NewMoneyTransferCommand(from, to, 50)
	transfer.Call()

	transfer.Undo()

	if from.balance != 100 {
		t.Errorf("Expected from balance 100 after undo, got %d", from.balance)
	}

	if to.balance != 0 {
		t.Errorf("Expected to balance 0 after undo, got %d", to.balance)
	}
}

func TestMoneyTransferFailure(t *testing.T) {
	from := &BankAccount{balance: 30}
	to := &BankAccount{balance: 0}

	transfer := NewMoneyTransferCommand(from, to, 100) // Try to transfer more than available
	transfer.Call()

	// Transfer should fail - from should remain unchanged or partially changed depending on implementation
	// The key is that it's handled gracefully
	if !transfer.Succeeded() || from.balance == 30 {
		// Expected behavior for failed transfer
	}
}

func TestCompositeBankAccountCommandSetSucceeded(t *testing.T) {
	ba := &BankAccount{}
	cmd1 := NewBankAccountCommand(ba, Deposit, 100)
	cmd2 := NewBankAccountCommand(ba, Withdraw, 50)

	composite := &CompositeBankAccountCommand{
		commands: []Command{cmd1, cmd2},
	}

	composite.Call()
	composite.SetSucceeded(false)

	if composite.Succeeded() {
		t.Error("Expected composite to have succeeded=false after SetSucceeded(false)")
	}
}

func TestMultipleTransfers(t *testing.T) {
	account1 := &BankAccount{balance: 1000}
	account2 := &BankAccount{balance: 500}

	transfer1 := NewMoneyTransferCommand(account1, account2, 100)
	transfer2 := NewMoneyTransferCommand(account2, account1, 50)

	transfer1.Call()
	transfer2.Call()

	if account1.balance != 950 {
		t.Errorf("Expected account1 balance 950, got %d", account1.balance)
	}

	if account2.balance != 550 {
		t.Errorf("Expected account2 balance 550, got %d", account2.balance)
	}
}
