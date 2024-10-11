package count

import (
	"bufio"
	"io"
	"unicode/utf8"
)

// readFile processes each line, chunk, or rune of an input using provided callback functions.
func readFile(r io.Reader, processLine func(string), processChunk func([]byte) int, processRune func(rune) int) error {
	// If processLine is not nil, use it to process each line
	if processLine != nil {
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			processLine(scanner.Text())
		}
		return scanner.Err()
	}

	// If processRune is provided, use ReadRune to process each rune
	if processRune != nil {
		reader := bufio.NewReader(r)
		for {
			runeValue, _, err := reader.ReadRune() // Read one rune (supports multibyte characters)
			if err != nil {
				if err == io.EOF {
					break
				}
				return err
			}
			if runeValue != utf8.RuneError { // Only count valid runes
				processRune(runeValue)
			}
		}
		return nil
	}

	// If processChunk is not nil, process file in chunks
	if processChunk != nil {
		reader := bufio.NewReader(r)
		buffer := make([]byte, 4096)
		for {
			n, err := reader.Read(buffer)
			if err != nil && err != io.EOF {
				return err
			}
			if n == 0 {
				break
			}
			buffer = buffer[:n]
			processChunk(buffer)
		}
	}
	return nil
}
