package cmd

import "testing"

func TestCountLines(t *testing.T) {
	filename := "../test/test.txt"
	expectedLineCount := 7145

	lineCount, err := countLines(filename)
	if err != nil {
		t.Fatalf("Error counting lines: %v", err)
	}

	if lineCount != expectedLineCount {
		t.Errorf("expected %d lines, but got %d lines", expectedLineCount, lineCount)
	}
}
