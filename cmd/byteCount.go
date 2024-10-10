package cmd

import (
	"io/ioutil"
	"os"
)

func countBytes(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return 0, err
	}

	return len(content), nil
}
