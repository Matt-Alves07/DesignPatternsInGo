package main

import (
	"testing"
)

func TestFunctionalTemplateMethod(t *testing.T) {
	turn := 1
	maxTurns := 10
	currentPlayer := 0

	start := func() {
		// Initialize
	}

	takeTurn := func() {
		turn++
		currentPlayer = (currentPlayer + 1) % 2
	}

	haveWinner := func() bool {
		return turn == maxTurns
	}

	winningPlayer := func() int {
		return currentPlayer
	}

	// Run the game
	PlayGame(start, takeTurn, haveWinner, winningPlayer)

	if turn != maxTurns {
		t.Errorf("Expected turn to be %d, got %d", maxTurns, turn)
	}
}

func TestFunctionalGameLogic(t *testing.T) {
	turn := 1
	maxTurns := 5
	currentPlayer := 0

	start := func() {
		// Game start
	}

	takeTurn := func() {
		turn++
		currentPlayer = (currentPlayer + 1) % 2
	}

	haveWinner := func() bool {
		return turn == maxTurns
	}

	winningPlayer := func() int {
		return currentPlayer
	}

	// Simulate game
	start()
	for !haveWinner() {
		takeTurn()
	}

	if turn != maxTurns {
		t.Errorf("Expected final turn %d, got %d", maxTurns, turn)
	}

	if winningPlayer() != currentPlayer {
		t.Error("Expected correct winning player")
	}
}

func TestPlayerAlternation(t *testing.T) {
	currentPlayer := 0

	takeTurn := func() {
		currentPlayer = (currentPlayer + 1) % 2
	}

	// Player 0
	if currentPlayer != 0 {
		t.Error("Expected to start with player 0")
	}

	takeTurn()
	// Player 1
	if currentPlayer != 1 {
		t.Error("Expected player 1 after first turn")
	}

	takeTurn()
	// Player 0
	if currentPlayer != 0 {
		t.Error("Expected player 0 after second turn")
	}
}

func TestGameCompletion(t *testing.T) {
	turn := 1
	completed := false

	haveWinner := func() bool {
		return turn == 3
	}

	start := func() {}
	takeTurn := func() { turn++ }

	start()
	for !haveWinner() {
		takeTurn()
	}
	completed = true

	if !completed {
		t.Error("Expected game to complete")
	}

	if turn != 3 {
		t.Errorf("Expected turn 3, got %d", turn)
	}
}

func TestClosureState(t *testing.T) {
	turn := 1
	maxTurns := 4
	currentPlayer := 0

	takeTurn := func() {
		turn++
		currentPlayer = (currentPlayer + 1) % 2
	}

	haveWinner := func() bool {
		return turn == maxTurns
	}

	// Closures should maintain state
	takeTurn()
	if turn != 2 {
		t.Error("Expected closure to update turn")
	}

	if haveWinner() {
		t.Error("Expected game not won yet")
	}

	for !haveWinner() {
		takeTurn()
	}

	if turn != maxTurns {
		t.Errorf("Expected final turn %d, got %d", maxTurns, turn)
	}
}

func TestFunctionalTemplateGameFlow(t *testing.T) {
	gameOutput := ""
	turn := 1
	maxTurns := 2
	currentPlayer := 0

	start := func() {
		gameOutput += "Start;"
	}

	takeTurn := func() {
		turn++
		currentPlayer = (currentPlayer + 1) % 2
		gameOutput += "Turn;"
	}

	haveWinner := func() bool {
		return turn == maxTurns
	}

	winningPlayer := func() int {
		return currentPlayer
	}

	PlayGame(start, takeTurn, haveWinner, winningPlayer)

	if gameOutput == "" {
		t.Error("Expected game to produce output")
	}
}
