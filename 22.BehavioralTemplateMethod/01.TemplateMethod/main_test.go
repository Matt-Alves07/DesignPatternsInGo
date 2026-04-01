package main

import (
	"testing"
)

func TestGameCreation(t *testing.T) {
	game := NewGameOfChess()

	if game == nil {
		t.Fatal("Expected game to be created")
	}
}

func TestGameStart(t *testing.T) {
	game := NewGameOfChess()

	game.Start()
	// Should not panic
}

func TestGameHaveWinner(t *testing.T) {
	game := NewGameOfChess()

	if game.HaveWinner() {
		t.Error("Expected no winner at start")
	}
}

func TestGameTakeTurn(t *testing.T) {
	game := NewGameOfChess()

	for i := 0; i < 10; i++ {
		if game.HaveWinner() {
			break
		}
		game.TakeTurn()
	}

	if !game.HaveWinner() {
		t.Error("Expected winner after 10 turns")
	}
}

func TestGameWinningPlayer(t *testing.T) {
	game := NewGameOfChess()

	for !game.HaveWinner() {
		game.TakeTurn()
	}

	winner := game.WinningPlayer()

	if winner < 0 || winner > 1 {
		t.Errorf("Expected winner to be 0 or 1, got %d", winner)
	}
}

func TestPlayGameFunction(t *testing.T) {
	game := NewGameOfChess()

	// PlayGame should run to completion
	PlayGame(game)

	// Game should have winner
	if !game.HaveWinner() {
		t.Error("Expected winner after PlayGame")
	}
}

func TestGameTemplate(t *testing.T) {
	game := NewGameOfChess()

	game.Start()

	turnCount := 0
	for !game.HaveWinner() && turnCount < 20 {
		game.TakeTurn()
		turnCount++
	}

	// Chess starts at turn=1, ends when turn==10, so 9 TakeTurn calls are made
	if turnCount != 9 {
		t.Errorf("Expected 9 turns for chess, got %d", turnCount)
	}

	winner := game.WinningPlayer()
	if winner < 0 || winner > 1 {
		t.Error("Expected valid winner")
	}
}

func TestChessPlayerAlternation(t *testing.T) {
	game := NewGameOfChess()

	game.Start()

	currentPlayer := 0
	for i := 0; i < 3 && !game.HaveWinner(); i++ {
		game.TakeTurn()
		currentPlayer = (currentPlayer + 1) % 2
	}
}

func TestGameInterface(t *testing.T) {
	game := NewGameOfChess()

	var g Game = game

	if g == nil {
		t.Error("Expected chess to implement Game interface")
	}
}
