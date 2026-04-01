package main

import (
	"sync"
	"testing"
)

func TestGameCreation(t *testing.T) {
	game := &Game{sync.Map{}}

	if game == nil {
		t.Fatal("Expected game to be created")
	}
}

func TestCreatureInGameCreation(t *testing.T) {
	game := &Game{sync.Map{}}
	creature := NewCreature(game, "Goblin", 2, 2)

	if creature.Name != "Goblin" {
		t.Errorf("Expected name 'Goblin', got %s", creature.Name)
	}

	if creature.Attack() != 2 {
		t.Errorf("Expected attack 2, got %d", creature.Attack())
	}

	if creature.Defense() != 2 {
		t.Errorf("Expected defense 2, got %d", creature.Defense())
	}
}

func TestDoubleAttackModifier(t *testing.T) {
	game := &Game{sync.Map{}}
	goblin := NewCreature(game, "Goblin", 2, 2)

	initialAttack := goblin.Attack()

	modifier := NewDoubleAttackModifier(game, goblin)

	modifiedAttack := goblin.Attack()

	if modifiedAttack != initialAttack*2 {
		t.Errorf("Expected attack to be doubled, got %d (was %d)", modifiedAttack, initialAttack)
	}

	modifier.Close()
}

func TestModifierRemoval(t *testing.T) {
	game := &Game{sync.Map{}}
	goblin := NewCreature(game, "Goblin", 2, 2)

	modifier := NewDoubleAttackModifier(game, goblin)
	attackWithModifier := goblin.Attack()

	modifier.Close()

	attackWithoutModifier := goblin.Attack()

	if attackWithModifier != attackWithoutModifier*2 {
		t.Errorf("Expected modifier to be removed, attack should be halved")
	}
}

func TestCreatureString(t *testing.T) {
	game := &Game{sync.Map{}}
	creature := NewCreature(game, "Orc", 5, 3)
	output := creature.String()

	if output != "Orc (5/3)" {
		t.Errorf("Expected 'Orc (5/3)', got '%s'", output)
	}
}

func TestGameSubscribeAndFire(t *testing.T) {
	game := &Game{sync.Map{}}

	called := false
	var observer Observer = &testObserver{
		handleFunc: func(q *Query) {
			called = true
		},
	}

	game.Subscribe(observer)
	game.Fire(&Query{})

	if !called {
		t.Error("Expected observer to be called")
	}
}

type testObserver struct {
	handleFunc func(*Query)
}

func (t *testObserver) Handle(q *Query) {
	if t.handleFunc != nil {
		t.handleFunc(q)
	}
}

func TestGameUnsubscribe(t *testing.T) {
	game := &Game{sync.Map{}}

	called := false
	observer := &testObserver{
		handleFunc: func(q *Query) {
			called = true
		},
	}

	game.Subscribe(observer)
	game.Unsubscribe(observer)
	game.Fire(&Query{})

	if called {
		t.Error("Expected observer to not be called after unsubscribe")
	}
}

func TestQueryAttack(t *testing.T) {
	q := Query{CreatureName: "Goblin", WhatToQuery: Attack, Value: 5}

	if q.WhatToQuery != Attack {
		t.Error("Expected query to be for Attack")
	}
}

func TestQueryDefense(t *testing.T) {
	q := Query{CreatureName: "Goblin", WhatToQuery: Defense, Value: 3}

	if q.WhatToQuery != Defense {
		t.Error("Expected query to be for Defense")
	}
}

func TestBrokerChainPattern(t *testing.T) {
	game := &Game{sync.Map{}}
	goblin := NewCreature(game, "Goblin", 1, 1)

	initialAttack := goblin.Attack()

	modifier := NewDoubleAttackModifier(game, goblin)
	modifiedAttack := goblin.Attack()

	if modifiedAttack == initialAttack {
		t.Error("Expected attack to be modified by broker chain")
	}

	modifier.Close()
}
