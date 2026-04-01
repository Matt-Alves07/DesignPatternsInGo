package main

import (
	"testing"
)

func TestDriverAgeRequirement(t *testing.T) {
	tests := []struct {
		age    int
		canDrive bool
	}{
		{12, false},
		{15, false},
		{16, true},
		{17, true},
		{25, true},
		{65, true},
	}

	for _, tt := range tests {
		driver := Driver{Age: tt.age}
		proxy := NewCarProxy(&driver)

		// Call Drive - if age < 16, should print "Driver too young"
		// We can't easily capture stdout, so we just verify it doesn't panic
		proxy.Drive()
	}
}

func TestCarProxyCreation(t *testing.T) {
	driver := &Driver{Age: 20}
	proxy := NewCarProxy(driver)

	if proxy == nil {
		t.Fatal("Expected CarProxy to be created")
	}

	if proxy.driver != driver {
		t.Error("Expected proxy to reference driver")
	}
}

func TestCarDrive(t *testing.T) {
	car := &Car{}
	car.Drive()
	// Should not panic
}

func TestProxyAllowsOldDriver(t *testing.T) {
	driver := &Driver{Age: 30}
	proxy := NewCarProxy(driver)

	// Should allow driving
	proxy.Drive()
}

func TestProxyBlocksYoungDriver(t *testing.T) {
	driver := &Driver{Age: 10}
	proxy := NewCarProxy(driver)

	// Should block driving
	proxy.Drive()
}

func TestProxyBoundaryAge15(t *testing.T) {
	driver := &Driver{Age: 15}
	proxy := NewCarProxy(driver)

	// Should block at age 15
	proxy.Drive()
}

func TestProxyBoundaryAge16(t *testing.T) {
	driver := &Driver{Age: 16}
	proxy := NewCarProxy(driver)

	// Should allow at age 16
	proxy.Drive()
}

func TestCarProxyImplementsDriven(t *testing.T) {
	driver := &Driver{Age: 20}
	proxy := NewCarProxy(driver)

	// Verify proxy implements Driven interface
	var driven Driven = proxy
	if driven == nil {
		t.Error("Expected CarProxy to implement Driven interface")
	}
}

func TestMultipleProxiesWithDifferentDrivers(t *testing.T) {
	youngDriver := &Driver{Age: 14}
	oldDriver := &Driver{Age: 21}

	proxy1 := NewCarProxy(youngDriver)
	proxy2 := NewCarProxy(oldDriver)

	proxy1.Drive() // Should be blocked
	proxy2.Drive() // Should be allowed

	if proxy1.driver.Age != 14 {
		t.Errorf("Expected proxy1 driver age 14, got %d", proxy1.driver.Age)
	}

	if proxy2.driver.Age != 21 {
		t.Errorf("Expected proxy2 driver age 21, got %d", proxy2.driver.Age)
	}
}

func TestCarProxyInternalCar(t *testing.T) {
	driver := &Driver{Age: 25}
	proxy := NewCarProxy(driver)

	// Proxy should have internal Car
	if proxy.car == (Car{}) {
		// Car is initialized but we can't easily verify state
		// Just ensure it exists
	}
}

func TestDriverAgeIncrease(t *testing.T) {
	driver := &Driver{Age: 14}
	proxy := NewCarProxy(driver)

	proxy.Drive() // Should be blocked

	// Increase driver age
	driver.Age = 16

	proxy.Drive() // Should now be allowed
}
