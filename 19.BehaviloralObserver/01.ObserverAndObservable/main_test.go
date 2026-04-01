package main

import (
	"container/list"
	"testing"
)

func TestPersonObserverCreation(t *testing.T) {
	person := NewPerson("Boris")

	if person.Name != "Boris" {
		t.Errorf("Expected name 'Boris', got %s", person.Name)
	}
}

func TestPersonObservableInit(t *testing.T) {
	person := NewPerson("Alice")

	if person.subs == nil {
		t.Error("Expected person to have initialized subs list")
	}
}

func TestDoctorServiceNotification(t *testing.T) {
	person := NewPerson("Boris")
	doctor := &DoctorService{}

	person.Subscribe(doctor)
	person.CatchACold()

	// Doctor should be notified (via Fire method calling Notify)
	// We can verify this indirectly by ensuring no panic
}

func TestObservableSubscribe(t *testing.T) {
	observable := &Observable{list.New()}

	called := false
	observer := &TestObserver{
		notifyFunc: func(data interface{}) {
			called = true
		},
	}

	observable.Subscribe(observer)
	observable.Fire("test")

	if !called {
		t.Error("Expected observer to be notified")
	}
}

type TestObserver struct {
	notifyFunc func(interface{})
}

func (t *TestObserver) Notify(data interface{}) {
	if t.notifyFunc != nil {
		t.notifyFunc(data)
	}
}

func TestObservableUnsubscribe(t *testing.T) {
	observable := &Observable{list.New()}

	called := false
	observer := &TestObserver{
		notifyFunc: func(data interface{}) {
			called = true
		},
	}

	observable.Subscribe(observer)
	observable.Unsubscribe(observer)
	observable.Fire("test")

	if called {
		t.Error("Expected observer not to be notified after unsubscribe")
	}
}

func TestMultipleObservers(t *testing.T) {
	observable := &Observable{list.New()}

	callCount := 0
	observer1 := &TestObserver{
		notifyFunc: func(data interface{}) {
			callCount++
		},
	}

	observer2 := &TestObserver{
		notifyFunc: func(data interface{}) {
			callCount++
		},
	}

	observable.Subscribe(observer1)
	observable.Subscribe(observer2)
	observable.Fire("test")

	if callCount != 2 {
		t.Errorf("Expected 2 observer calls, got %d", callCount)
	}
}

func TestPersonCatchCold(t *testing.T) {
	person := NewPerson("Boris")

	count := 0
	observer := &TestObserver{
		notifyFunc: func(data interface{}) {
			if name, ok := data.(string); ok && name == "Boris" {
				count++
			}
		},
	}

	person.Subscribe(observer)
	person.CatchACold()

	if count != 1 {
		t.Error("Expected person to fire cold event")
	}
}

func TestObserverPatternDecoupling(t *testing.T) {
	person := NewPerson("Patient")
	doctor := &DoctorService{}

	// Doctor doesn't know about person initially
	person.Subscribe(doctor)

	// When person catches cold, doctor is notified
	person.CatchACold()

	// Direct reference not necessary - communication through observer pattern
}
