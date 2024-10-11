package test

import (
	"os"
	"testing"

	"github.com/rafaelmgr12/cwcc-tool/internal/count"
)

func TestCountLines(t *testing.T) {
	filename := "test.txt"
	expectedLineCount := 7145

	file, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Error opening file %v", err)
	}
	defer file.Close()

	lineCount, err := count.CountLines(file)
	if err != nil {
		t.Fatalf("Error counting lines: %v", err)
	}

	if lineCount != expectedLineCount {
		t.Errorf("expected %d lines, but got %d lines", expectedLineCount, lineCount)
	}
}
