package main

import "testing"

func TestSquare(t *testing.T) {
	// Test case: Square of 4 should be 16
	if result := Square(4); result != 16 {
		t.Errorf("Expected 16, but got %d", result)
	}

	// Test case: Square of 0 should be 0
	if result := Square(0); result != 0 {
		t.Errorf("Expected 0, but got %d", result)
	}
}
