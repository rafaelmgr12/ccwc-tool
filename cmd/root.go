package cmd

import (
	"fmt"
	"log"
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
	RunE: func(cmd *cobra.Command, args []string) error {
		// Use RunE to return errors instead of log.Fatalf
		if err := execCommand(args); err != nil {
			log.Printf("Command execution failed: %v", err)
			return err // Return the error instead of calling log.Fatalf
		}
		return nil
	},
}

// The `Execute` function sets up flags for byte, line, and word counting and then executes the root
// command.
func Execute() {
	// Define the flags in a separated function
	defineFlags()

	// Execute the root command and handle the error gracefully
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Command execution failed:", err)
		os.Exit(1)
	}
}
