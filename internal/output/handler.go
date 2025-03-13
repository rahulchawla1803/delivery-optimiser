package output

import (
	"fmt"
	"os"
	"path/filepath"
)

// PrepareOutputDirectory ensures output folder exists & clears old files
func PrepareOutputDirectory() error {
	outputDir := "output"

	// Ensure directory exists
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create output directory: %v", err)
		}
		return nil // No need to clear files if directory was just created
	}

	// Clear old files inside the directory, but do not delete the directory itself
	err := filepath.Walk(outputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return os.Remove(path) // Delete only files
		}
		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to clear output directory: %v", err)
	}

	return nil
}
