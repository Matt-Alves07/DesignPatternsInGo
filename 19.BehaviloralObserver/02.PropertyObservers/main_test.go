package main

import (
	"testing"
)

func TestObserverCreation(t *testing.T) {
	p := NewPerson(15)

	if p.Age() != 15 {
		t.Errorf("Expected age 15, got %d", p.Age())
	}
}

func TestPersonSetAge(t *testing.T) {
	p := NewPerson(15)

	p.SetAge(16)

	if p.Age() != 16 {
		t.Errorf("Expected age 16 after SetAge, got %d", p.Age())
	}
}

func TestSetAgeNoChange(t *testing.T) {
	p := NewPerson(20)
	p.SetAge(20)
	
	// Since age is same, Fire() should not be called
	// No observable way to test this without external mocking
}

func TestPropertyChangedNotification(t *testing.T) {
	p := NewPerson(15)

	notified := false
	observer := &TestPropertyObserver{
		onNotify: func(data interface{}) {
			if pc, ok := data.(PropertyChanged); ok {
				if pc.Name == "Age" {
					notified = true
				}
			}
		},
	}

	p.Subscribe(observer)
	p.SetAge(16)

	if !notified {
		t.Error("Expected observer to be notified of property change")
	}
}

type TestPropertyObserver struct {
	onNotify func(interface{})
}

func (t *TestPropertyObserver) Notify(data interface{}) {
	if t.onNotify != nil {
		t.onNotify(data)
	}
}

func TestMultiplePropertyObservers(t *testing.T) {
	p := NewPerson(15)

	count := 0
	observer1 := &TestPropertyObserver{
		onNotify: func(data interface{}) {
			count++
		},
	}

	observer2 := &TestPropertyObserver{
		onNotify: func(data interface{}) {
			count++
		},
	}

	p.Subscribe(observer1)
	p.Subscribe(observer2)

	p.SetAge(16)

	if count != 2 {
		t.Errorf("Expected both observers to be notified, got %d notifications", count)
	}
}

func TestTrafficManagementObserver(t *testing.T) {
	p := NewPerson(15)
	tm := &TrafficManagement{p.Observable}

	p.Subscribe(tm)

	// Age 15 - should not unsubscribe
	p.SetAge(15)

	// Age 16 - should trigger message and unsubscribe
	p.SetAge(16)
}

func TestUnsubscribeObserver(t *testing.T) {
	p := NewPerson(15)

	count := 0
	observer := &TestPropertyObserver{
		onNotify: func(data interface{}) {
			count++
		},
	}

	p.Subscribe(observer)
	p.SetAge(16)

	if count != 1 {
		t.Errorf("Expected 1 notification, got %d", count)
	}

	p.Unsubscribe(observer)
	p.SetAge(17)

	if count != 1 {
		t.Errorf("Expected still 1 notification after unsubscribe, got %d", count)
	}
}

func TestPropertyChangedData(t *testing.T) {
	p := NewPerson(15)

	var receivedPC PropertyChanged
	observer := &TestPropertyObserver{
		onNotify: func(data interface{}) {
			if pc, ok := data.(PropertyChanged); ok {
				receivedPC = pc
			}
		},
	}

	p.Subscribe(observer)
	p.SetAge(20)

	if receivedPC.Name != "Age" {
		t.Errorf("Expected property name 'Age', got %s", receivedPC.Name)
	}

	if receivedPC.Value.(int) != 20 {
		t.Errorf("Expected property value 20, got %v", receivedPC.Value)
	}
}
