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
)

// Root command setup
var rootCmd = &cobra.Command{
	Use:   "ccwc",
	Short: "ccwc is a word count tool",
	Long: `A Go implementation of the classic Unix 'wc' command-line tool to count bytes, chars, words, and lines.
	
Usage:
  ccwc -c [file]   count the number of bytes in the file
  ccwc -l [file]   count the lines in the file
  ccwc -w [file]   count the words in the file`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Error: You must provide a filename.")
			cmd.Help() // Show help if no arguments are provided
			return
		}

		// Handle byte counting
		if countBytesFlag {
			for _, filename := range args {
				byteCount, err := countBytes(filename)
				if err != nil {
					fmt.Printf("Error reading file %s: %v\n", filename, err)
					continue
				}
				fmt.Printf("%d %s\n", byteCount, filename)
			}
		}

		// Handle line counting
		if countLineFlag {
			for _, filename := range args {
				lineCount, err := countLines(filename)
				if err != nil {
					fmt.Printf("Error reading file %s: %v\n", filename, err)
					continue
				}
				fmt.Printf("%d %s\n", lineCount, filename)
			}
		}

		// Handle word couting
		if countWordsFlag {
			for _, filename := range args {
				wordCount, err := countWords(filename)
				if err != nil {
					fmt.Printf("Error reading file %s: %v\n", filename, err)
					continue
				}
				fmt.Printf("%d %s\n", wordCount, filename)
			}
		}

		// If no flags are set, show help
		if !countBytesFlag && !countLineFlag && !countWordsFlag {
			cmd.Help()
		}
	},
}

// The `Execute` function sets up flags for byte, line, and word counting and then executes the root
// command.
func Execute() {
	// Add the flags for byte and line counting
	rootCmd.Flags().BoolVarP(&countBytesFlag, "bytes", "c", false, "Count bytes in the file")
	rootCmd.Flags().BoolVarP(&countLineFlag, "lines", "l", false, "Count lines in the file")
	rootCmd.Flags().BoolVarP(&countWordsFlag, "words", "w", false, "Count words in the file")

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
