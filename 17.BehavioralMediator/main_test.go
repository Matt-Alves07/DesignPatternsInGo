package main

import (
	"testing"
)

func TestPersonCreationMediator(t *testing.T) {
	person := NewPerson("John")

	if person.Name != "John" {
		t.Errorf("Expected name 'John', got %s", person.Name)
	}

	if len(person.chatLog) != 0 {
		t.Errorf("Expected empty chat log initially, got %d", len(person.chatLog))
	}
}

func TestPersonReceive(t *testing.T) {
	person := NewPerson("Alice")

	person.Receive("Bob", "Hello")

	if len(person.chatLog) != 1 {
		t.Errorf("Expected 1 message in chat log, got %d", len(person.chatLog))
	}
}

func TestChatRoomCreation(t *testing.T) {
	room := &ChatRoom{}

	if room == nil {
		t.Fatal("Expected ChatRoom to be created")
	}

	if len(room.people) != 0 {
		t.Errorf("Expected 0 people initially, got %d", len(room.people))
	}
}

func TestChatRoomJoin(t *testing.T) {
	room := &ChatRoom{}
	alice := NewPerson("Alice")

	room.Join(alice)

	if len(room.people) != 1 {
		t.Errorf("Expected 1 person in room, got %d", len(room.people))
	}

	if alice.Room != room {
		t.Error("Expected person's room to be set")
	}
}

func TestChatRoomBroadcast(t *testing.T) {
	room := &ChatRoom{}
	alice := NewPerson("Alice")
	bob := NewPerson("Bob")

	room.Join(alice)
	room.Join(bob)
	// After joins: alice has 1 (Bob joins), bob has 0

	room.Broadcast("Alice", "Hello everyone")
	// After broadcast: bob receives 1 message

	if len(bob.chatLog) != 1 {
		t.Errorf("Expected 1 message for Bob, got %d", len(bob.chatLog))
	}
}

func TestChatRoomPrivateMessage(t *testing.T) {
	room := &ChatRoom{}
	alice := NewPerson("Alice")
	bob := NewPerson("Bob")
	charlie := NewPerson("Charlie")

	room.Join(alice)
	room.Join(bob)
	room.Join(charlie)
	// After joins: alice has 2, bob has 1, charlie has 0

	beforeCount := len(bob.chatLog)
	alice.PrivateMessage("Bob", "Secret message")
	afterCount := len(bob.chatLog)

	if afterCount != beforeCount+1 {
		t.Errorf("Expected Bob to receive private message, before=%d after=%d", beforeCount, afterCount)
	}

	// Charlie should not receive private message
	charlieInitial := len(charlie.chatLog)
	alice.PrivateMessage("Bob", "Another secret")
	charlieAfter := len(charlie.chatLog)

	if charlieAfter > charlieInitial {
		t.Error("Expected Charlie not to receive private message")
	}
}

func TestPersonSay(t *testing.T) {
	room := &ChatRoom{}
	alice := NewPerson("Alice")
	bob := NewPerson("Bob")

	room.Join(alice)
	room.Join(bob)
	// After joins: alice has 1 (Bob joins), bob has 0

	alice.Say("Hi Bob")
	// After say: bob receives 1 message

	if len(bob.chatLog) != 1 {
		t.Errorf("Expected Bob to receive 1 message from Alice, got %d", len(bob.chatLog))
	}
}

func TestChatRoomMultipleJoins(t *testing.T) {
	room := &ChatRoom{}
	alice := NewPerson("Alice")
	bob := NewPerson("Bob")
	charlie := NewPerson("Charlie")

	room.Join(alice)
	room.Join(bob)
	room.Join(charlie)

	if len(room.people) != 3 {
		t.Errorf("Expected 3 people in room, got %d", len(room.people))
	}

	// All should receive join announcements
	if len(alice.chatLog) < 2 { // Bob join + Charlie join
		t.Error("Expected Alice to receive other join announcements")
	}
}

func TestMediatorPatternDecoupling(t *testing.T) {
	room := &ChatRoom{}
	person1 := NewPerson("Person1")
	person2 := NewPerson("Person2")

	room.Join(person1)
	room.Join(person2)

	// People communicate only through ChatRoom
	person1.Say("Message")

	// Person2 receives through mediator
	if len(person2.chatLog) == 0 {
		t.Error("Expected message through mediator")
	}

	// Persons don't have direct reference to each other
	if person1.Room != room || person2.Room != room {
		t.Error("Expected both to reference room")
	}
}
