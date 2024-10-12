package cmd

// The `defineFlags` function sets up flags for byte, line, word, and character counting in a file.
func defineFlags() {
	// Add the flags for byte, line, and word counting
	rootCmd.Flags().BoolVarP(&countBytesFlag, "bytes", "c", false, "Count bytes in the file")
	rootCmd.Flags().BoolVarP(&countLineFlag, "lines", "l", false, "Count lines in the file")
	rootCmd.Flags().BoolVarP(&countWordsFlag, "words", "w", false, "Count words in the file")
	rootCmd.Flags().BoolVarP(&countCharsFlag, "chars", "m", false, "Count characters in the file")
}
