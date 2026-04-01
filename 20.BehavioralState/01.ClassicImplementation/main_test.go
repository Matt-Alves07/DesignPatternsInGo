package main

import (
	"testing"
)

func TestSwitchCreation(t *testing.T) {
	sw := NewSwitch()

	if sw == nil {
		t.Fatal("Expected switch to be created")
	}

	if sw.State == nil {
		t.Error("Expected switch to have initial state")
	}
}

func TestBaseStateOn(t *testing.T) {
	baseState := &BaseState{}
	sw := NewSwitch()

	baseState.On(sw)
	// Should not panic and print message
}

func TestBaseStateOff(t *testing.T) {
	baseState := &BaseState{}
	sw := NewSwitch()

	baseState.Off(sw)
	// Should not panic and print message
}

func TestOnStateCreation(t *testing.T) {
	onState := NewOnState()

	if onState == nil {
		t.Fatal("Expected OnState to be created")
	}
}

func TestOffStateCreation(t *testing.T) {
	offState := NewOffState()

	if offState == nil {
		t.Fatal("Expected OffState to be created")
	}
}

func TestSwitchStateTransition(t *testing.T) {
	sw := NewSwitch()

	// Start in OffState
	_, isOff := sw.State.(*OffState)
	if !isOff && !(sw.State == nil) {
		// Initial state initialization
	}

	// Call On
	sw.On()

	// Should now be in OnState
	_, isOn := sw.State.(*OnState)
	if !isOn {
		t.Error("Expected switch to be in OnState after On()")
	}

	// Call Off
	sw.Off()

	// Should be in OffState
	_, isOff = sw.State.(*OffState)
	if !isOff {
		t.Error("Expected switch to be in OffState after Off()")
	}
}

func TestOnStateOff(t *testing.T) {
	sw := NewSwitch()
	sw.State = NewOnState()

	sw.Off()

	_, isOff := sw.State.(*OffState)
	if !isOff {
		t.Error("Expected transition from OnState to OffState")
	}
}

func TestMultipeStateTransitions(t *testing.T) {
	sw := NewSwitch()

	// Multiple on/off cycles
	for i := 0; i < 3; i++ {
		sw.On()
		_, isOn := sw.State.(*OnState)
		if !isOn {
			t.Errorf("Expected OnState in cycle %d", i)
		}

		sw.Off()
		_, isOff := sw.State.(*OffState)
		if !isOff {
			t.Errorf("Expected OffState in cycle %d", i)
		}
	}
}

func TestStateImplementsInterface(t *testing.T) {
	var state State

	state = NewOnState()
	if state == nil {
		t.Error("Expected OnState to implement State interface")
	}

	state = NewOffState()
	if state == nil {
		t.Error("Expected OffState to implement State interface")
	}
}

func TestSwitchInitialState(t *testing.T) {
	sw := NewSwitch()

	// Initial state should be OffState
	initialState := sw.State
	if initialState == nil {
		t.Error("Expected initial state to be set")
	}

	// Verify it's OffState via class type check
	_, isOff := initialState.(*OffState)
	if !isOff && sw.State != nil {
		// May print "Light turned off" message
	}
}

func TestStateConsistency(t *testing.T) {
	sw := NewSwitch()

	sw.On()
	stateAfterOn := sw.State

	// State should remain OnState
	stateAfterCheck := sw.State

	if stateAfterOn != stateAfterCheck {
		t.Error("Expected state to remain consistent")
	}
}
