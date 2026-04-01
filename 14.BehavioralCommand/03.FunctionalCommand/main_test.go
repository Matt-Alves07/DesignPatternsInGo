package main

import (
	"testing"
)

func TestBankAccountFunctional(t *testing.T) {
	ba := &BankAccount{Balance: 0}

	if ba.Balance != 0 {
		t.Errorf("Expected initial balance 0, got %d", ba.Balance)
	}
}

func TestDepositFunctional(t *testing.T) {
	ba := &BankAccount{Balance: 100}

	Deposit(ba, 50)

	if ba.Balance != 150 {
		t.Errorf("Expected balance 150, got %d", ba.Balance)
	}
}

func TestWithdrawFunctional(t *testing.T) {
	ba := &BankAccount{Balance: 100}

	Withdraw(ba, 30)

	if ba.Balance != 70 {
		t.Errorf("Expected balance 70, got %d", ba.Balance)
	}
}

func TestWithdrawInsufficientFunds(t *testing.T) {
	ba := &BankAccount{Balance: 20}

	Withdraw(ba, 100)

	if ba.Balance != 20 {
		t.Errorf("Expected balance to remain 20, got %d", ba.Balance)
	}
}

func TestFunctionalCommandsAsClosures(t *testing.T) {
	ba := &BankAccount{Balance: 0}

	var commands []func()

	commands = append(commands, func() {
		Deposit(ba, 100)
	})

	commands = append(commands, func() {
		Withdraw(ba, 50)
	})

	for _, cmd := range commands {
		cmd()
	}

	if ba.Balance != 50 {
		t.Errorf("Expected balance 50 after executing commands, got %d", ba.Balance)
	}
}

func TestMultipleFunctionalCommands(t *testing.T) {
	ba := &BankAccount{Balance: 0}

	commands := []func(){
		func() { Deposit(ba, 100) },
		func() { Deposit(ba, 50) },
		func() { Withdraw(ba, 30) },
		func() { Withdraw(ba, 20) },
	}

	for _, cmd := range commands {
		cmd()
	}

	if ba.Balance != 100 {
		t.Errorf("Expected balance 100, got %d", ba.Balance)
	}
}

func TestFunctionalCommandOrder(t *testing.T) {
	ba := &BankAccount{Balance: 0}

	var operations []func()

	operations = append(operations, func() { Deposit(ba, 100) })
	operations = append(operations, func() { Withdraw(ba, 25) })
	operations = append(operations, func() { Deposit(ba, 75) })

	for _, op := range operations {
		op()
	}

	if ba.Balance != 150 {
		t.Errorf("Expected balance 150, got %d", ba.Balance)
	}
}

func TestFunctionalCommandWithHigherOrder(t *testing.T) {
	ba := &BankAccount{Balance: 0}

	// Create a command generator
	makeDepositCommand := func(amount int) func() {
		return func() {
			Deposit(ba, amount)
		}
	}

	cmd1 := makeDepositCommand(100)
	cmd2 := makeDepositCommand(50)

	cmd1()
	cmd2()

	if ba.Balance != 150 {
		t.Errorf("Expected balance 150, got %d", ba.Balance)
	}
}
