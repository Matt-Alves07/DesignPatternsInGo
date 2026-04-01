package main

import (
	"testing"
)

func TestCreatureCreation(t *testing.T) {
	creature := NewCreature("Goblin", 1, 1)

	if creature.Name != "Goblin" {
		t.Errorf("Expected name 'Goblin', got %s", creature.Name)
	}

	if creature.Attack != 1 {
		t.Errorf("Expected attack 1, got %d", creature.Attack)
	}

	if creature.Defense != 1 {
		t.Errorf("Expected defense 1, got %d", creature.Defense)
	}
}

func TestCreatureString(t *testing.T) {
	creature := NewCreature("Orc", 5, 3)
	output := creature.String()

	if output != "Orc (5/3)" {
		t.Errorf("Expected 'Orc (5/3)', got '%s'", output)
	}
}

func TestDoubleAttackModifier(t *testing.T) {
	creature := NewCreature("Goblin", 2, 1)
	modifier := NewDoubleAttackModifier(creature)

	if modifier.creature != creature {
		t.Error("Expected modifier to reference creature")
	}

	modifier.Handle()

	if creature.Attack != 4 {
		t.Errorf("Expected attack to be doubled to 4, got %d", creature.Attack)
	}
}

func TestIncreasedDefenseModifier(t *testing.T) {
	creature := NewCreature("Goblin", 1, 1)
	modifier := NewIncreasedDefenseModifier(creature)

	modifier.Handle()

	if creature.Defense != 2 {
		t.Errorf("Expected defense to increase to 2, got %d", creature.Defense)
	}
}

func TestIncreasedDefenseModifierNoEffectOnStrongAttacker(t *testing.T) {
	creature := NewCreature("Goblin", 5, 1)
	modifier := NewIncreasedDefenseModifier(creature)

	originalDefense := creature.Defense
	modifier.Handle()

	if creature.Defense != originalDefense {
		t.Errorf("Expected defense to remain %d, got %d", originalDefense, creature.Defense)
	}
}

func TestNoBonusesModifier(t *testing.T) {
	creature := NewCreature("Goblin", 1, 1)
	modifier := NewNoBonusesModifier(creature)

	originalAttack := creature.Attack
	originalDefense := creature.Defense

	modifier.Handle()

	if creature.Attack != originalAttack {
		t.Error("Expected attack to be unchanged")
	}

	if creature.Defense != originalDefense {
		t.Error("Expected defense to be unchanged")
	}
}

func TestChainOfResponsibility(t *testing.T) {
	creature := NewCreature("Goblin", 1, 1)
	root := NewCreatureModifier(creature)

	root.Add(NewDoubleAttackModifier(creature))
	root.Add(NewIncreasedDefenseModifier(creature))

	root.Handle()

	if creature.Attack != 2 {
		t.Errorf("Expected attack to be doubled to 2, got %d", creature.Attack)
	}

	if creature.Defense != 2 {
		t.Errorf("Expected defense to increase to 2, got %d", creature.Defense)
	}
}

func TestMultipleModifiers(t *testing.T) {
	creature := NewCreature("Goblin", 1, 1)
	root := NewCreatureModifier(creature)

	root.Add(NewDoubleAttackModifier(creature))
	root.Add(NewDoubleAttackModifier(creature))

	root.Handle()

	if creature.Attack != 4 {
		t.Errorf("Expected attack to be 4, got %d", creature.Attack)
	}
}

func TestModifierAdd(t *testing.T) {
	creature := NewCreature("Goblin", 1, 1)
	root := NewCreatureModifier(creature)
	mod1 := NewDoubleAttackModifier(creature)
	mod2 := NewDoubleAttackModifier(creature)

	root.Add(mod1)
	if root.next != mod1 {
		t.Error("Expected first modifier to be added to root")
	}

	root.Add(mod2)
	if mod1.next != mod2 {
		t.Error("Expected second modifier to be chained after first")
	}
}
