package cmd

import (
	"bufio"
	"io"
)

// countLines reads from an io.Reader and returns the total number of lines along with any errors encountered.
func countLines(r io.Reader) (int, error) {
	reader := bufio.NewReader(r)
	lineCount := 0

	for {
		_, err := reader.ReadString('\n') // Read until a newline character
		if err != nil {
			if err == io.EOF {
				break // Break the loop if end of file is reached
			}
			return 0, err
		}
		lineCount++ // Increment the line count for each newline character
	}

	return lineCount, nil
}
