package cmd

import (
	"os"
	"testing"
)

func TestCountChars(t *testing.T) {
	// Sample test file and expected character count
	filename := "../test/test.txt"
	expectedCharCount := 339292 // Set this to the actual number of characters in the test.txt file
	file, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Error opening file %v", err)
	}
	defer file.Close()
	// Call the function to count characters
	charCount, err := countChars(file)
	if err != nil {
		t.Fatalf("Error counting characters: %v", err)
	}

	// Compare the result with the expected character count
	if charCount != expectedCharCount {
		t.Errorf("expected %d characters, but got %d characters", expectedCharCount, charCount)
	}
}
