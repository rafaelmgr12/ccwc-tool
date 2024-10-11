package count

import (
	"io"
	"strings"
)

// CountWords reads from an io.Reader and counts the number of words.
func CountWords(r io.Reader) (int, error) {
	wordCount := 0

	// Use readFile with the processLine callback to count words
	err := readFile(r, func(line string) {
		words := strings.Fields(line) // Split line into words
		wordCount += len(words)       // Increment word count by the number of words in the line
	}, nil, nil) // The chunk and rune processors are set to nil as they are not used for counting words

	if err != nil {
		return 0, err
	}

	return wordCount, nil
}
