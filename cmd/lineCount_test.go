package cmd

import (
	"os"
	"testing"
)

func TestCountLines(t *testing.T) {
	filename := "../test/test.txt"
	expectedLineCount := 7145

	file, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Error opening file %v", err)
	}
	defer file.Close()

	lineCount, err := countLines(file)
	if err != nil {
		t.Fatalf("Error counting lines: %v", err)
	}

	if lineCount != expectedLineCount {
		t.Errorf("expected %d lines, but got %d lines", expectedLineCount, lineCount)
	}
}
