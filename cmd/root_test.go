package cmd

import (
	"bytes"
	"io"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Helper function to capture both stdout and log output (to handle fmt.Println and log output)
func captureAllOutput(f func()) string {
	// Create pipes to capture both stdout and logs
	rOut, wOut, _ := os.Pipe()
	rLog, wLog, _ := os.Pipe()

	// Save original stdout and log output
	stdout := os.Stdout
	stderr := os.Stderr

	// Redirect stdout and log output
	os.Stdout = wOut
	log.SetOutput(wLog)

	// Buffers to capture output
	var bufOut bytes.Buffer
	var bufLog bytes.Buffer

	// Channels to wait for completion of goroutines
	doneOut := make(chan bool)
	doneLog := make(chan bool)

	// Goroutine to capture stdout
	go func() {
		io.Copy(&bufOut, rOut)
		doneOut <- true
	}()

	// Goroutine to capture log output
	go func() {
		io.Copy(&bufLog, rLog)
		doneLog <- true
	}()

	// Execute the function
	f()

	// Close the write ends of the pipes
	wOut.Close()
	wLog.Close()

	// Restore original stdout and log output
	os.Stdout = stdout
	os.Stderr = stderr

	// Wait for goroutines to finish
	<-doneOut
	<-doneLog

	// Combine both outputs
	return bufOut.String() + bufLog.String()
}

// Helper function to create a temporary file and write data into it
func createTempFile(t *testing.T, content string) *os.File {
	// Create a temporary file in the system's temp directory
	tmpFile, err := os.CreateTemp("", "testfile_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	// Write the provided content into the temp file
	if _, err := tmpFile.Write([]byte(content)); err != nil {
		t.Fatalf("Failed to write content to temp file: %v", err)
	}

	// Close the file and reopen it to reset the file pointer
	tmpFile.Close()

	// Reopen the file in read mode
	tmpFile, err = os.Open(tmpFile.Name())
	if err != nil {
		t.Fatalf("Failed to reopen temp file: %v", err)
	}

	// Return the temp file (which can be used in tests)
	return tmpFile
}

func TestRootCommand_ExecuteWithTempFile(t *testing.T) {
	// Create a temporary file with sample content
	tmpFile := createTempFile(t, "This is a test file\nwith multiple lines\nand some words.")
	defer os.Remove(tmpFile.Name()) // Ensure the file is deleted after the test

	// Mock flags for testing
	countBytesFlag = false
	countLineFlag = false
	countWordsFlag = false
	countCharsFlag = false

	// Define arguments for the test, using the temporary file path
	args := []string{tmpFile.Name()}
	rootCmd.SetArgs(args)

	// Capture the output of the command execution
	output := captureAllOutput(func() {
		err := rootCmd.Execute()
		assert.NoError(t, err)
	})

	assert.Contains(t, output, tmpFile.Name(), "Output should contain the temporary file name")
}

func TestRootCommand_BytesFlag(t *testing.T) {
	// Create a temporary file with sample content
	tmpFile := createTempFile(t, "Some content to count bytes.")
	defer os.Remove(tmpFile.Name()) // Ensure the file is deleted after the test

	// Set the flag for counting bytes
	countBytesFlag = true
	countLineFlag = false
	countWordsFlag = false
	countCharsFlag = false

	// Define arguments for the test, using the temporary file path
	args := []string{tmpFile.Name()}
	rootCmd.SetArgs(args)

	// Capture the output of the command execution
	output := captureAllOutput(func() {
		err := rootCmd.Execute()
		assert.NoError(t, err)
	})

	assert.Contains(t, output, tmpFile.Name(), "Output should contain the temporary file name")
	assert.Contains(t, output, "28", "Output should contain byte count (28 bytes for this file)")
}

func TestRootCommand_NoFileProvided(t *testing.T) {
	// Mock flags for testing, ensuring no flags are set
	countBytesFlag = false
	countLineFlag = false
	countWordsFlag = false
	countCharsFlag = false

	// Define an empty argument list (simulating no file provided)
	args := []string{}
	rootCmd.SetArgs(args)

	// Capture all output (both stdout and log output)
	output := captureAllOutput(func() {
		err := rootCmd.Execute()
		assert.Error(t, err, "An error should occur when no file is provided")
	})

	// Check that the output contains the error message
	assert.Contains(t, output, "need a file to read or input from stdin", "Output should indicate that a file is needed")
}

func TestRootCommand_LineFlag(t *testing.T) {
	// Create a temporary file with sample content
	tmpFile := createTempFile(t, "Line one\nLine two\nLine three.")
	defer os.Remove(tmpFile.Name()) // Ensure the file is deleted after the test

	// Set the flag for counting lines
	countBytesFlag = false
	countLineFlag = true
	countWordsFlag = false
	countCharsFlag = false

	// Define arguments for the test, using the temporary file path
	args := []string{tmpFile.Name()}
	rootCmd.SetArgs(args)

	// Capture the output of the command execution
	output := captureAllOutput(func() {
		err := rootCmd.Execute()
		assert.NoError(t, err)
	})

	assert.Contains(t, output, tmpFile.Name(), "Output should contain the temporary file name")
	assert.Contains(t, output, "3", "Output should contain line count (3 lines for this file)")
}

func TestRootCommand_WordsFlag(t *testing.T) {
	// Create a temporary file with sample content
	tmpFile := createTempFile(t, "This is a sample text file with several words.")
	defer os.Remove(tmpFile.Name()) // Ensure the file is deleted after the test

	// Set the flag for counting words
	countBytesFlag = false
	countLineFlag = false
	countWordsFlag = true
	countCharsFlag = false

	// Define arguments for the test, using the temporary file path
	args := []string{tmpFile.Name()}
	rootCmd.SetArgs(args)

	// Capture the output of the command execution
	output := captureAllOutput(func() {
		err := rootCmd.Execute()
		assert.NoError(t, err)
	})

	assert.Contains(t, output, tmpFile.Name(), "Output should contain the temporary file name")
	assert.Contains(t, output, "9", "Output should contain word count (9 words for this file)")
}
