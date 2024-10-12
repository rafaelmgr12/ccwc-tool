package cmd

import (
	"fmt"
	"os"

	"github.com/rafaelmgr12/cwcc-tool/internal/count"
	"github.com/spf13/cobra"
)

// The `execCommand` function in Go reads input from a file or stdin, checks for flags, and executes
// count functions based on the provided flags.
func execCommand(cmd *cobra.Command, args []string) {
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
	var fileName string

	if len(args) > 0 {
		reader, err = os.Open(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer reader.Close()
		fileName = reader.Name() // Captura o nome do arquivo
	} else {
		reader = os.Stdin
		// Se for stdin, defina o nome como "stdin"
		fileName = ""
		// If stdin is a terminal and flags are present, do not block
		if isTerminal {
			fmt.Println("error: reading from stdin is not supported in terminal mode without file")
			os.Exit(1)
		}
	}

	// Passa o nome do arquivo para a função que executa as contagens
	executeCountFlags(reader, fileName)
}

// executeCountFlags processes the flags and executes the appropriate count functions
func executeCountFlags(reader *os.File, fileName string) {
	if countBytesFlag {
		byteCount, err := count.CountBytes(reader)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("%d %s\n", byteCount, fileName)
	}

	if countLineFlag {
		lineCount, err := count.CountLines(reader)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("%d %s\n", lineCount, fileName)
	}

	if countWordsFlag {
		wordCount, err := count.CountWords(reader)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("%d %s\n", wordCount, fileName)
	}

	if countCharsFlag {
		charCount, err := count.CountChars(reader)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("%d %s\n", charCount, fileName)
	}

	// Default behavior if no flags are provided
	if !countBytesFlag && !countLineFlag && !countWordsFlag && !countCharsFlag {
		reader.Seek(0, os.SEEK_SET)
		lineCount, err := count.CountLines(reader)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		reader.Seek(0, os.SEEK_SET)
		wordCount, err := count.CountWords(reader)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		reader.Seek(0, os.SEEK_SET)
		byteCount, err := count.CountBytes(reader)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("%d %d %d %s\n", lineCount, wordCount, byteCount, fileName)
	}
}
