package main

import (
	"testing"
)

func TestStateCreation(t *testing.T) {
	var state State = OffHook

	if state != OffHook {
		t.Errorf("Expected OffHook state, got %v", state)
	}
}

func TestStateName(t *testing.T) {
	tests := []struct {
		state State
		name  string
	}{
		{OffHook, "OffHook"},
		{Connecting, "Connecting"},
		{Connected, "Connected"},
		{OnHold, "OnHold"},
		{OnHook, "OnHook"},
	}

	for _, tt := range tests {
		if tt.state.String() != tt.name {
			t.Errorf("Expected '%s', got '%s'", tt.name, tt.state.String())
		}
	}
}

func TestTriggerName(t *testing.T) {
	tests := []struct {
		trigger Trigger
		name    string
	}{
		{CallDialed, "CallDialed"},
		{HungUp, "HungUp"},
		{CallConnected, "CallConnected"},
		{PlacedOnHold, "PlacedOnHold"},
		{TakenOffHold, "TakenOffHold"},
		{LeftMessage, "LeftMessage"},
	}

	for _, tt := range tests {
		if tt.trigger.String() != tt.name {
			t.Errorf("Expected '%s', got '%s'", tt.name, tt.trigger.String())
		}
	}
}

func TestStateTransitions(t *testing.T) {
	// OffHook -> CallDialed -> Connecting
	transitions := rules[OffHook]

	if len(transitions) == 0 {
		t.Error("Expected transitions from OffHook")
	}

	if transitions[0].Trigger != CallDialed {
		t.Errorf("Expected CallDialed trigger, got %v", transitions[0].Trigger)
	}

	if transitions[0].State != Connecting {
		t.Errorf("Expected Connecting state, got %v", transitions[0].State)
	}
}

func TestConnectingStateTransitions(t *testing.T) {
	transitions := rules[Connecting]

	expects := map[Trigger]State{
		HungUp:         OffHook,
		CallConnected:  Connected,
	}

	for trigger, expectedState := range expects {
		found := false
		for _, tr := range transitions {
			if tr.Trigger == trigger && tr.State == expectedState {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected transition Connecting --%v--> %v", trigger, expectedState)
		}
	}
}

func TestConnectedStateTransitions(t *testing.T) {
	transitions := rules[Connected]

	if len(transitions) < 3 {
		t.Errorf("Expected at least 3 transitions from Connected, got %d", len(transitions))
	}
}

func TestOnHoldStateTransitions(t *testing.T) {
	transitions := rules[OnHold]

	if len(transitions) != 2 {
		t.Errorf("Expected 2 transitions from OnHold, got %d", len(transitions))
	}
}

func TestPhoneStateMachine(t *testing.T) {
	state := OffHook

	// Dial call
	state = rules[state][0].State // CallDialed -> Connecting
	if state != Connecting {
		t.Errorf("Expected Connecting, got %v", state)
	}

	// Call connects
	state = rules[state][1].State // CallConnected -> Connected
	if state != Connected {
		t.Errorf("Expected Connected, got %v", state)
	}

	// Put on hold
	state = rules[state][2].State // PlacedOnHold -> OnHold
	if state != OnHold {
		t.Errorf("Expected OnHold, got %v", state)
	}

	// Take off hold
	state = rules[state][0].State // TakenOffHold -> Connected
	if state != Connected {
		t.Errorf("Expected Connected, got %v", state)
	}
}

func TestInvalidTransitionHandling(t *testing.T) {
	// OnHook state has no transitions
	transitions := rules[OnHook]

	if len(transitions) > 0 {
		t.Errorf("Expected no transitions from OnHook, got %d", len(transitions))
	}
}

func TestTriggerResultStructure(t *testing.T) {
	tr := TriggerResult{
		Trigger: CallDialed,
		State:   Connecting,
	}

	if tr.Trigger != CallDialed {
		t.Error("Expected trigger field to be set")
	}

	if tr.State != Connecting {
		t.Error("Expected state field to be set")
	}
}
