package cmd

import (
	"fmt"
	"os"

	"github.com/rafaelmgr12/cwcc-tool/internal/count"
	"github.com/spf13/cobra"
)

func execCommand(cmd *cobra.Command, args []string) {
	// Check if stdin is a terminal
	info, err := os.Stdin.Stat()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	isTerminal := (info.Mode() & os.ModeCharDevice) != 0

	// No args, no flags, and stdin is a terminal, show usage and exit
	if len(args) <= 0 && !countBytesFlag && !countLineFlag && !countWordsFlag && !countCharsFlag && isTerminal {
		fmt.Println("error: need a file to read or input from stdin")
		cmd.Usage()
		os.Exit(1)
	}

	var reader *os.File

	if len(args) > 0 {
		reader, err = os.Open(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer reader.Close()
	} else {
		reader = os.Stdin
		// If stdin is a terminal and flags are present, do not block
		if isTerminal {
			fmt.Println("error: reading from stdin is not supported in terminal mode without file")
			os.Exit(1)
		}
	}

	if countBytesFlag {
		byteCount, err := count.CountBytes(reader)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("%d\n", byteCount)
	}

	if countLineFlag {
		lineCount, err := count.CountLines(reader)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("%d\n", lineCount)
	}

	if countWordsFlag {
		wordCount, err := count.CountWords(reader)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("%d\n", wordCount)
	}

	if countCharsFlag {
		charCount, err := count.CountChars(reader)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("%d\n", charCount)
	}

	// Default behavior if no flags are provided
	if !countBytesFlag && !countLineFlag && !countWordsFlag && !countCharsFlag {
		lineCount, err := count.CountLines(reader)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		wordCount, err := count.CountWords(reader)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Need to reset reader to count bytes after reading lines and words
		reader.Seek(0, os.SEEK_SET)
		byteCount, err := count.CountBytes(reader)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("%d %d %d\n", lineCount, wordCount, byteCount)
	}
}
