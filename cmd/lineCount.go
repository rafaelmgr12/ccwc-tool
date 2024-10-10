package cmd

import (
	"bufio"
	"os"
)

// countLines reads the file in chunks and counts the number of newline characters
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
