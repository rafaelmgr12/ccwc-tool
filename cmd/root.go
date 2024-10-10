package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var countBytesFlag bool

var rootCmd = &cobra.Command{
	Use:   "ccwc",
	Short: "ccwc is a word count tool",
	Long: `A Go implementation of the classic Unix 'wc' command-line tool to count bytes, chars, words, and lines.
	
Usage:
  ccwc -c [file]   count the number of bytes in the file`,
	Run: func(cmd *cobra.Command, args []string) {
		if countBytesFlag && len(args) > 0 {
			filename := args[0]
			byteCount, err := countBytes(filename)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}
			fmt.Printf("%d %s\n", byteCount, filename)
		} else {
			cmd.Help()
		}
	},
}

func Execute() {
	rootCmd.Flags().BoolVarP(&countBytesFlag, "bytes", "c", false, "Count bytes in the file")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
