package cmd

import (
	"fmt"
	"os"

	"github.com/rafaelmgr12/cwcc-tool/internal/count"
)

// The `execCommand` function in Go reads input from a file or stdin, checks for flags, and executes
// count functions based on the provided flags.
func execCommand(args []string) error {
	info, err := os.Stdin.Stat()
	if err != nil {
		return fmt.Errorf("failed to stat stdin: %w", err)
	}

	isTerminal := (info.Mode() & os.ModeCharDevice) != 0

	// No args, no flags, and stdin is a terminal, show usage and exit
	if shouldShowMessageUsage(args, isTerminal) {
		return fmt.Errorf("error: need a file to read or input from stdin")
	}

	reader, fileName, err := getReaderAndFileName(args, isTerminal)
	if err != nil {
		return fmt.Errorf("failed to get reader and file name: %w", err)
	}
	defer reader.Close()

	// Pass the file name to the function that performs the counts
	if err := executeCountFlags(reader, fileName); err != nil {
		return fmt.Errorf("failed to execute count flags: %w", err)
	}
	return nil
}

func getReaderAndFileName(args []string, isTerminal bool) (*os.File, string, error) {
	if len(args) > 0 {
		file, err := os.Open(args[0])
		if err != nil {
			return nil, "", err
		}
		return file, file.Name(), nil
	}

	if isTerminal {
		return nil, "", fmt.Errorf("error: reading from stdin is not supported in terminal mode without file")
	}

	return os.Stdin, "", nil

}

func shouldShowMessageUsage(args []string, isTerminal bool) bool {
	return len(args) <= 0 && !countBytesFlag && !countLineFlag && !countWordsFlag && !countCharsFlag && isTerminal
}

// executeCountFlags processes the flags and executes the appropriate count functions
func executeCountFlags(reader *os.File, fileName string) error {
	if countBytesFlag {
		byteCount, err := count.CountBytes(reader)
		if err != nil {
			return fmt.Errorf("failed to count bytes: %w", err)
		}
		fmt.Printf("%d %s\n", byteCount, fileName)
	}

	if countLineFlag {
		lineCount, err := count.CountLines(reader)
		if err != nil {
			return fmt.Errorf("failed to count lines: %w", err)
		}
		fmt.Printf("%d %s\n", lineCount, fileName)
	}

	if countWordsFlag {
		wordCount, err := count.CountWords(reader)
		if err != nil {
			return fmt.Errorf("failed to count words: %w", err)
		}
		fmt.Printf("%d %s\n", wordCount, fileName)
	}

	if countCharsFlag {
		charCount, err := count.CountChars(reader)
		if err != nil {
			return fmt.Errorf("failed to count chars: %w", err)
		}
		fmt.Printf("%d %s\n", charCount, fileName)
	}

	// Default behavior if no flags are provided
	if !countBytesFlag && !countLineFlag && !countWordsFlag && !countCharsFlag {
		reader.Seek(0, os.SEEK_SET)
		lineCount, err := count.CountLines(reader)
		if err != nil {
			return fmt.Errorf("failed to count lines: %w", err)
		}

		reader.Seek(0, os.SEEK_SET)
		wordCount, err := count.CountWords(reader)
		if err != nil {
			return fmt.Errorf("failed to count words: %w", err)
		}
		reader.Seek(0, os.SEEK_SET)
		byteCount, err := count.CountBytes(reader)
		if err != nil {
			return fmt.Errorf("failed to count bytes: %w", err)
		}
		fmt.Printf("%d %d %d %s\n", lineCount, wordCount, byteCount, fileName)
	}

	return nil
}
