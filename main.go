package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

// countBytes counts the number of bytes in the file
func countBytes(filename string) (int, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	// Read the entire file into memory
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return 0, err
	}

	// Return the length (number of bytes)
	return len(content), nil
}

func main() {
	// Define the -c flag to count bytes
	byteFlag := flag.Bool("c", false, "Count bytes in the file")

	// Parse the flags from the command line
	flag.Parse()

	// Check if a file is provided
	if len(flag.Args()) < 1 {
		fmt.Println("Usage: ccwc -c <filename>")
		return
	}

	// Get the filename from the arguments
	filename := flag.Arg(0)

	// If -c flag is provided, count the number of bytes
	if *byteFlag {
		byteCount, err := countBytes(filename)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		fmt.Printf("%d %s\n", byteCount, filename)
	}
}
