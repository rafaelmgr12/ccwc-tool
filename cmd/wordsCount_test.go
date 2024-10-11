package cmd

import (
	"os"
	"testing"
)

func TestCountWords(t *testing.T) {
	filename := "../test/test.txt"
	expectedWordsCount := 58164

	file, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Error opening file %v", err)
	}
	defer file.Close()

	wordsCount, err := countWords(file)
	if err != nil {
		t.Fatalf("Error counting words: %v", err)
	}

	if wordsCount != expectedWordsCount {
		t.Errorf("expected %d lines, but got %d lines", expectedWordsCount, wordsCount)
	}
}
