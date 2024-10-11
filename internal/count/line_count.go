package count

import "io"

func CountLines(r io.Reader) (int, error) {
	lineCount := 0

	// Use readFile with the processLine callback to count lines.
	err := readFile(r, func(line string) {
		lineCount++ // Increment for each line
	}, nil, nil) // The chunk and rune processors are set to nil

	if err != nil {
		return 0, err
	}

	return lineCount, nil
}
