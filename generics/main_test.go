package main

import "testing"

func TestSumInt(t *testing.T) {
	output := Sum(1, 2, 3, 4)
	expectedOutput := 10
	if output != 10 {
		t.Errorf("Expected: %d, Actual: %d", expectedOutput, output)
	}
}

func TestSumFloat(t *testing.T) {
	output := Sum(1.1, 2.0, 3.0, 4.0)
	expectedOutput := 10.1
	if output != 10.1 {
		t.Errorf("Expected: %f, Actual: %f", expectedOutput, output)
	}
}
