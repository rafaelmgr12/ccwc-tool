package cmd

import (
	"bufio"
	"io"
	"unicode/utf8"
)

// countChars reads from an io.Reader and returns the total number of characters along with any errors encountered.
func countChars(r io.Reader) (int, error) {
	reader := bufio.NewReader(r)
	charCount := 0

	for {
		runeValue, _, err := reader.ReadRune() // Read one rune (supports multibyte)
		if err != nil {
			if err == io.EOF {
				break
			}
			return 0, err
		}
		if runeValue != utf8.RuneError { // Only count valid characters
			charCount++
		}
	}

	return charCount, nil
}
