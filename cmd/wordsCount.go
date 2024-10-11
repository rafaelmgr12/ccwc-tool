package cmd

import (
	"bufio"
	"os"
	"strings"
)

// countWords reads the file line by line and counts the words in the file
func countWords(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	wordCount := 0
	scanner := bufio.NewScanner(file)

	// Scan the file line by line
	for scanner.Scan() {
		line := scanner.Text()

		// Use strings.Fields to split the line into words
		// strings.Fields splits by any whitespace and returns a slice of words
		words := strings.Fields(line)
		wordCount += len(words)
	}

	// Check for errors while scanning
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return wordCount, nil
}

