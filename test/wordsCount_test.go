package test

import (
	"os"
	"testing"

	"github.com/rafaelmgr12/cwcc-tool/internal/count"
)

func TestCountWords(t *testing.T) {
	filename := "test.txt"
	expectedWordsCount := 58164

	file, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Error opening file %v", err)
	}
	defer file.Close()

	wordsCount, err := count.CountWords(file)
	if err != nil {
		t.Fatalf("Error counting words: %v", err)
	}

	if wordsCount != expectedWordsCount {
		t.Errorf("expected %d lines, but got %d lines", expectedWordsCount, wordsCount)
	}
}
