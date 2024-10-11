package count

import (
	"io"
)

// The `CountChars` function reads from an `io.Reader` and counts the number of characters in the
// input.
func CountChars(r io.Reader) (int, error) {
	charCount := 0

	err := readFile(r, nil, nil, func(r rune) int {
		charCount++
		return charCount
	})

	if err != nil {
		return 0, err
	}

	return charCount, nil
}
