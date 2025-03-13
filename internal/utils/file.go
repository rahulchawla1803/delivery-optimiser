package utils

import (
	"fmt"
	"os"
)

// WriteText writes text instructions to a file
func WriteText(filename string, lines []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	for _, line := range lines {
		if _, err := file.WriteString(line + "\n"); err != nil {
			return fmt.Errorf("error writing to file: %v", err)
		}
	}
	return nil
}
