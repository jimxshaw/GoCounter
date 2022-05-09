package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	testInput := bytes.NewBufferString("one two three four five six\n")

	expected := 6
	got := count(testInput, false)

	if got != expected {
		t.Errorf("Expected %d, got %d instead.\n", expected, got)
	}
}

func TestCountLines(t *testing.T) {
	testInput := bytes.NewBufferString("This is Line 1.\nThis is Line 2.\nThis is Line 3.")

	expected := 3
	got := count(testInput, true)

	if got != expected {
		t.Errorf("Expected %d, got %d instead.\n", expected, got)
	}
}
