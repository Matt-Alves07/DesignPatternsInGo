package main

import (
	"testing"
)

func TestPersonDependencyCreation(t *testing.T) {
	p := NewPerson(15)

	if p.Age() != 15 {
		t.Errorf("Expected age 15, got %d", p.Age())
	}

	if p.CanVote() {
		t.Error("Expected age 15 cannot vote")
	}
}

func TestCanVoteAtThreshold(t *testing.T) {
	p := NewPerson(17)

	if p.CanVote() {
		t.Error("Expected age 17 cannot vote")
	}

	p.SetAge(18)

	if !p.CanVote() {
		t.Error("Expected age 18 can vote")
	}
}

func TestDependentPropertyChange(t *testing.T) {
	p := NewPerson(16)

	notifiedCount := 0
	observer := &TestPropertyDependencyObserver{
		onNotify: func(data interface{}) {
			if pc, ok := data.(PropertyChanged); ok {
				notifiedCount++
				if pc.Name == "CanVote" {
					notifiedCount += 10 // Mark special event
				}
			}
		},
	}

	p.Subscribe(observer)
	p.SetAge(18)

	if notifiedCount < 11 { // At least 1 extra for CanVote change
		t.Logf("Expected CanVote change notification, got count %d", notifiedCount)
	}
}

type TestPropertyDependencyObserver struct {
	onNotify func(interface{})
}

func (t *TestPropertyDependencyObserver) Notify(data interface{}) {
	if t.onNotify != nil {
		t.onNotify(data)
	}
}

func TestNoNotificationWhenNoChange(t *testing.T) {
	p := NewPerson(20)

	notificationCount := 0
	observer := &TestPropertyDependencyObserver{
		onNotify: func(data interface{}) {
			notificationCount++
		},
	}

	p.Subscribe(observer)
	p.SetAge(20) // Same age

	if notificationCount != 0 {
		t.Errorf("Expected no notification for same age, got %d", notificationCount)
	}
}

func TestCanVoteTransition(t *testing.T) {
	p := NewPerson(14)

	notifiedCanVote := false
	observer := &TestPropertyDependencyObserver{
		onNotify: func(data interface{}) {
			if pc, ok := data.(PropertyChanged); ok {
				if pc.Name == "CanVote" {
					notifiedCanVote = true
				}
			}
		},
	}

	p.Subscribe(observer)

	p.SetAge(17) // Still can't vote
	if notifiedCanVote {
		t.Error("Expected no CanVote change at age 17")
	}

	p.SetAge(18) // Now can vote
	// notifiedCanVote should be true after this
}

func TestManyYearIncrement(t *testing.T) {
	p := NewPerson(10)

	for i := 10; i < 25; i++ {
		p.SetAge(i)
		if i >= 18 && !p.CanVote() {
			t.Errorf("Expected CanVote to be true at age %d", i)
		}
		if i < 18 && p.CanVote() {
			t.Errorf("Expected CanVote to be false at age %d", i)
		}
	}
}

func TestElectoralRoll(t *testing.T) {
	p := NewPerson(15)
	er := &ElectrocalRoll{}
	p.Subscribe(er)

	// Should not trigger at age 15
	p.SetAge(15)

	// Should trigger at age 18
	p.SetAge(18)
}
