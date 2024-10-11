package cmd

import (
	"fmt"
	"os"

	"github.com/rafaelmgr12/cwcc-tool/internal/count"
	"github.com/spf13/cobra"
)

var (
	countBytesFlag bool
	countLineFlag  bool
	countWordsFlag bool
	countCharsFlag bool
)

// Root command setup
var rootCmd = &cobra.Command{
	Use:   "ccwc",
	Short: "ccwc is a word count tool",
	Long: `A Go implementation of the classic Unix 'wc' command-line tool to count bytes, chars, words, and lines.
	
Usage:
  ccwc -c [file]   count the number of bytes in the file
  ccwc -l [file]   count the lines in the file
  ccwc -w [file]   count the words in the file
  ccwc -m [file]   count the characters in the file
  ccwc [file]      count bytes, lines, and words (default behavior)`,
	Run: func(cmd *cobra.Command, args []string) {
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
	},
}

// The `Execute` function sets up flags for byte, line, and word counting and then executes the root
// command.
func Execute() {
	// Add the flags for byte, line, and word counting
	rootCmd.Flags().BoolVarP(&countBytesFlag, "bytes", "c", false, "Count bytes in the file")
	rootCmd.Flags().BoolVarP(&countLineFlag, "lines", "l", false, "Count lines in the file")
	rootCmd.Flags().BoolVarP(&countWordsFlag, "words", "w", false, "Count words in the file")
	rootCmd.Flags().BoolVarP(&countCharsFlag, "chars", "m", false, "Count characters in the file")

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
