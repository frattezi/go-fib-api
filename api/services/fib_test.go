package services

import (
	"testing"
)

// TestZeroAsEntry calls Fib(0) to check base implementation
func TestZeroAsEntry(t *testing.T) {
	expected := 0
	result := Fib(0)
	if result != expected {
		t.Fatalf(`Expected result %v but got %v instead`, expected, result)
	}
}

// TestOneAsEntry calls Fib(1) to check base implementation
func TestOneAsEntry(t *testing.T) {
	expected := 1
	result := Fib(1)
	if result != expected {
		t.Fatalf(`Expected result %v but got %v instead`, expected, result)
	}
}

// TestFibResult validates Fibonacci sequence calculus for a intermedidate case
func TestFibResult(t *testing.T) {
	expected := 75025
	result := Fib(25)
	if result != expected {
		t.Fatalf(`Expected result %v but got %v instead`, expected, result)
	}
}
