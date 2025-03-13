package utils

import (
	"fmt"
	"os"
)

// WriteText writes a single string to a file
func WriteText(filename string, content string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(content + "\n")
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	return nil
}
