package cmd

import "testing"

func TestCountWords(t *testing.T) {
	filename := "../test/test.txt"
	expectedWordsCount := 58164

	wordsCount, err := countWords(filename)
	if err != nil {
		t.Fatalf("Error counting words: %v", err)
	}

	if wordsCount != expectedWordsCount {
		t.Errorf("expected %d lines, but got %d lines", expectedWordsCount, wordsCount)
	}
}
