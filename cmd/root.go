package cmd

import (
	"fmt"
	"os"

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
		if len(args) == 0 {
			fmt.Println("Error: You must provide a filename.")
			cmd.Help() // Show help if no arguments are provided
			return
		}

		for _, filename := range args {
			// If no flags are set, default to count bytes, lines, and words
			if !countBytesFlag && !countLineFlag && !countWordsFlag && !countCharsFlag {
				countBytesFlag = true
				countLineFlag = true
				countWordsFlag = true
			}

			// Handle byte counting
			if countBytesFlag {
				byteCount, err := countBytes(filename)
				if err != nil {
					fmt.Printf("Error reading file %s: %v\n", filename, err)
					continue
				}
				fmt.Printf("%8d ", byteCount)
			}

			// Handle line counting
			if countLineFlag {
				lineCount, err := countLines(filename)
				if err != nil {
					fmt.Printf("Error reading file %s: %v\n", filename, err)
					continue
				}
				fmt.Printf("%8d ", lineCount)
			}

			// Handle word counting
			if countWordsFlag {
				wordCount, err := countWords(filename)
				if err != nil {
					fmt.Printf("Error reading file %s: %v\n", filename, err)
					continue
				}
				fmt.Printf("%8d ", wordCount)
			}

			// Handle character counting
			if countCharsFlag {
				charCount, err := countChars(filename)
				if err != nil {
					fmt.Printf("Error reading file %s: %v\n", filename, err)
					continue
				}
				fmt.Printf("%8d ", charCount)
			}

			// Print the filename at the end of the output
			fmt.Printf("%s\n", filename)
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
