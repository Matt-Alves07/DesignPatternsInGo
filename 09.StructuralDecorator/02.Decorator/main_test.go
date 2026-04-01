package main

import (
	"testing"
)

func TestBirdAge(t *testing.T) {
	bird := Bird{}
	bird.SetAge(5)

	if bird.Age() != 5 {
		t.Errorf("Expected age 5, got %d", bird.Age())
	}
}

func TestBirdFly(t *testing.T) {
	bird := Bird{}
	bird.SetAge(10)
	bird.Fly()
	// Should not panic, output is to stdout
}

func TestBirdFlyTooYoung(t *testing.T) {
	bird := Bird{}
	bird.SetAge(5)
	bird.Fly()
	// Should not fly (age < 10)
}

func TestLizardAge(t *testing.T) {
	lizard := Lizard{}
	lizard.SetAge(8)

	if lizard.Age() != 8 {
		t.Errorf("Expected age 8, got %d", lizard.Age())
	}
}

func TestLizardCrawl(t *testing.T) {
	lizard := Lizard{}
	lizard.SetAge(5)
	lizard.Crawl()
	// Should crawl (age < 10)
}

func TestDragonCreation(t *testing.T) {
	dragon := NewDragon()

	if dragon == nil {
		t.Fatal("Expected dragon to be created")
	}

	if dragon.Age() != 0 {
		t.Errorf("Expected initial age 0, got %d", dragon.Age())
	}
}

func TestDragonSetAge(t *testing.T) {
	dragon := NewDragon()
	dragon.SetAge(10)

	if dragon.Age() != 10 {
		t.Errorf("Expected age 10, got %d", dragon.Age())
	}
}

func TestDragonCanFly(t *testing.T) {
	dragon := NewDragon()
	dragon.SetAge(10)
	dragon.Fly()
	// Should be able to fly at age 10 (bird age >= 10)
}

func TestDragonCanCrawl(t *testing.T) {
	dragon := NewDragon()
	dragon.SetAge(5)
	dragon.Crawl()
	// Should be able to crawl at age 5 (lizard age < 10)
}

func TestDragonFlyWhenTooYoung(t *testing.T) {
	dragon := NewDragon()
	dragon.SetAge(5)
	dragon.Fly()
	// Should not fly (age < 10)
}

func TestDragonCrawlWhenTooOld(t *testing.T) {
	dragon := NewDragon()
	dragon.SetAge(15)
	dragon.Crawl()
	// Should not crawl (age >= 10)
}

func TestDragonMultipleBehaviors(t *testing.T) {
	dragon := NewDragon()
	dragon.SetAge(10)

	// At age 10, should be able to fly (bird) but not crawl (lizard)
	dragon.Fly()

	dragon.SetAge(5)
	// At age 5, should not be able to fly (bird) but should crawl (lizard)
	dragon.Crawl()

	if dragon.Age() != 5 {
		t.Errorf("Expected age 5, got %d", dragon.Age())
	}
}

func TestDragonSynchronizedAge(t *testing.T) {
	dragon := NewDragon()
	dragon.SetAge(20)

	// Both bird and lizard should have the same age
	if dragon.Age() != 20 {
		t.Errorf("Expected synchronized age 20, got %d", dragon.Age())
	}

	// Setting age should update both
	dragon.SetAge(8)
	if dragon.Age() != 8 {
		t.Errorf("Expected synchronized age 8, got %d", dragon.Age())
	}
}
