package cmd

import (
	"bufio"
	"os"
)

// The countBytes function reads a file and returns the total number of bytes in the file along with
// any errors encountered.
func countBytes(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	byteCount := 0
	buffer := make([]byte, 4096)

	for {
		n, err := reader.Read(buffer)
		if err != nil && err.Error() != "EOF" {
			return 0, err
		}

		if n == 0 {
			break
		}

		byteCount += n
	}

	return byteCount, nil
}
