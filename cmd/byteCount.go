package cmd

import (
	"bufio"
	"io"
)

// The countBytes function reads a file and returns the total number of bytes in the file along with
// any errors encountered.
func countBytes(r io.Reader) (int, error) {

	reader := bufio.NewReader(r)
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
