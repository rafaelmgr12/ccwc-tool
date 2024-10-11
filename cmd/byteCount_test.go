package cmd

import (
	"os"
	"testing"
)

func TestCountBytes(t *testing.T) {
	filename := "../test/test.txt"
	expectedBytesCount := 342190

	file, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Error opening file %v", err)
	}
	defer file.Close()

	byteCount, err := countBytes(file)
	if err != nil {
		t.Fatalf("Error counting bytes %v", err)
	}

	if byteCount != expectedBytesCount {
		t.Errorf("expeted %d bytes, but go %d bytes", expectedBytesCount, byteCount)
	}
}
