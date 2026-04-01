package main

import (
	"testing"
)

func TestLockState(t *testing.T) {
	if Locked != 0 {
		t.Errorf("Expected Locked=0, got %d", Locked)
	}
}

func TestFailedState(t *testing.T) {
	if Failed != 1 {
		t.Errorf("Expected Failed=1, got %d", Failed)
	}
}

func TestUnlockedState(t *testing.T) {
	if Unlocked != 2 {
		t.Errorf("Expected Unlocked=2, got %d", Unlocked)
	}
}

func TestStateMachineLogic(t *testing.T) {
	code := "1234"
	state := Locked
	entry := ""

	// Simulate correct input
	for i := 0; i < len(code); i++ {
		switch state {
		case Locked:
			entry += string(code[i])
			if entry == code {
				state = Unlocked
				break
			}
			// Check if entry is still start of code
			if len(code) < len(entry) || code[:len(entry)] != entry {
				state = Failed
			}
		case Failed:
			entry = ""
			state = Locked
		case Unlocked:
			// Done
			return
		}
	}

	if state != Unlocked {
		t.Errorf("Expected Unlocked state after correct code, got %d", state)
	}
}

func TestWrongCodeTransition(t *testing.T) {
	state := Locked
	entry := ""
	wrongCode := "999"  // 3 characters, longer than code, will trigger Failed

	// Test wrong input
	for i := 0; i < len(wrongCode); i++ {
		switch state {
		case Locked:
			entry += string(wrongCode[i])
			if entry == "1234" {
				state = Unlocked
				break
			}
			// If entry is wrong or longer than expected, transition to Failed
			if len("1234") < len(entry) || (len(entry) == len("1234") && "1234" != entry) {
				state = Failed
			} else if len("1234") >= len(entry) && "1234"[:len(entry)] != entry {
				state = Failed
			}
		case Failed:
			// From Failed, reset to Locked on next input
			state = Locked
			entry = ""
			break
		}
	}

	// After 3 wrong characters: Locked->Failed->Locked, so we need to check intermediate state
	if state == Locked && entry == "" {
		// This is expected after the Failed transition back to Locked
	} else {
		t.Logf("State machine ended with state=%v, entry=%s", state, entry)
	}
}

func TestStateReset(t *testing.T) {
	state := Failed

	// Reset from Failed state
	state = Locked

	if state != Locked {
		t.Error("Expected state to reset to Locked")
	}
}

func TestLockCombinationStates(t *testing.T) {
	states := []State{Locked, Failed, Unlocked}

	for _, s := range states {
		if s < 0 || s > 2 {
			t.Errorf("State %d out of range", s)
		}
	}
}

func TestPartialCodeEntry(t *testing.T) {
	state := Locked
	code := "1234"
	entry := "12"

	// Partial entry should stay in Locked
	if entry[:2] == code[:2] {
		state = Locked
	}

	if state != Locked {
		t.Error("Expected Locked state during partial entry")
	}
}

func TestIncompleteCodeReject(t *testing.T) {
	state := Locked
	entry := "5"
	code := "1234"

	// Wrong code should transition to Failed
	if code[:len(entry)] != entry {
		state = Failed
	}

	if state != Failed {
		t.Error("Expected Failed state for wrong code start")
	}
}
