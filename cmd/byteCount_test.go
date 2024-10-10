package cmd

import (
	"testing"
)

func TestCountBytes(t *testing.T) {
	filename := "../test/test.txt"
	expectedBytesCount := 342190

	byteCount, err := countBytes(filename)
	if err != nil {
		t.Fatalf("Error counting bytes %v", err)
	}

	if byteCount != expectedBytesCount {
		t.Errorf("expeted %d bytes, but go %d bytes", expectedBytesCount, byteCount)
	}
}
