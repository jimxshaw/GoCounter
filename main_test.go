package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	testInput := bytes.NewBufferString("one two three four five six\n")

	expected := 6
	got := count(testInput)

	if got != expected {
		t.Errorf("Expected %d, got %d instead.\n", expected, got)
	}
}
