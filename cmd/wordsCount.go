package cmd

import (
	"bufio"
	"io"
	"strings"
)

// countWords reads from an io.Reader and returns the total word count along with any errors encountered.
func countWords(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	wordCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		wordCount += len(words)
	}

	// Check for errors while scanning
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return wordCount, nil
}
