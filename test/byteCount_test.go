package test

import (
	"os"
	"testing"

	"github.com/rafaelmgr12/cwcc-tool/internal/count"
)

func TestCountBytes(t *testing.T) {
	filename := "test.txt"
	expectedBytesCount := 342190

	file, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Error opening file %v", err)
	}
	defer file.Close()

	byteCount, err := count.CountBytes(file)
	if err != nil {
		t.Fatalf("Error counting bytes %v", err)
	}

	if byteCount != expectedBytesCount {
		t.Errorf("expeted %d bytes, but go %d bytes", expectedBytesCount, byteCount)
	}
}
