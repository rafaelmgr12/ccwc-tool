package cmd

import (
	"bufio"
	"os"
)

// The `countLines` function reads a file line by line and returns the total number of lines in the
// file along with any encountered errors.
func countLines(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	// Create a buffered reader to read the file in chunks
	reader := bufio.NewReader(file)
	lineCount := 0

	// Read the file line by line
	for {
		_, err := reader.ReadString('\n') // Read until a newline character
		if err != nil {
			if err.Error() == "EOF" {
				break // Break the loop if end of file is reached
			}
			return 0, err
		}
		lineCount++ // Increment the line count for each newline character
	}

	return lineCount, nil
}
