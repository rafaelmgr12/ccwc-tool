package count

import "io"

// The CountBytes function reads from an io.Reader and counts the total number of bytes.
func CountBytes(r io.Reader) (int, error) {
	byteCount := 0

	// Use readFile with the processChunk callback to count bytes
	err := readFile(r, nil, func(chunk []byte) int {
		byteCount += len(chunk) // Increment the byte count by the length of the chunk
		return len(chunk)       // Return the length of the chunk, though this value is not used
	}, nil) // processRune is set to nil since it's not needed for counting bytes

	if err != nil {
		return 0, err
	}

	return byteCount, nil
}
