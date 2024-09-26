package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// WriteWITToFile writes the generated WIT to the specified file path
func WriteWITToFile(filePath, wit string) error {
	// Ensure the directory structure exists
	err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create directories for WIT file: %w", err)
	}

	// Write the WIT content to the file
	err = os.WriteFile(filepath.Clean(filePath), []byte(wit), 0644)
	if err != nil {
		return fmt.Errorf("failed to write WIT to file: %w", err)
	}

	return nil
}
