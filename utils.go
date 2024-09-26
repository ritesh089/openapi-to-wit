package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// WriteWITToFile writes the generated WIT to the specified file path
func WriteWITToFile(filePath, wit string) error {
	err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create directories for WIT file: %w", err)
	}

	err = ioutil.WriteFile(filepath.Clean(filePath), []byte(wit), 0644)
	if err != nil {
		return fmt.Errorf("failed to write WIT to file: %w", err)
	}

	return nil
}
