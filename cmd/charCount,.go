package cmd

import (
	"bufio"
	"os"
	"unicode/utf8"
)

// The `countChars` function reads a file and returns the total number of characters in the file along
// with any errors encountered.
func countChars(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	charCount := 0

	// Read the file and count characters
	for {
		runeValue, _, err := reader.ReadRune() // Read one rune (supports multibyte)
		if err != nil {
			break // End of file or error
		}
		if runeValue != utf8.RuneError { // Only count valid characters
			charCount++
		}
	}

	return charCount, nil
}
