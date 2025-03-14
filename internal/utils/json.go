package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

// LoadJSON loads a JSON file and unmarshals it into a struct
func LoadJSON[T any](filePath string) (T, error) {
	var data T
	file, err := os.ReadFile(filePath)
	if err != nil {
		return data, fmt.Errorf("error reading file: %v", err)
	}

	err = json.Unmarshal(file, &data)
	if err != nil {
		return data, fmt.Errorf("error unmarshaling JSON: %v", err)
	}
	return data, nil
}

// PrintJSON pretty prints JSON data
func PrintJSON(title string, data interface{}) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err == nil {
		fmt.Printf("\n%s:\n%s\n", title, string(jsonData))
	}
}

// WriteJSON writes data to a JSON file
func WriteJSON(filename string, data interface{}) error {
	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %v", err)
	}

	if err := os.WriteFile(filename, file, 0644); err != nil {
		return fmt.Errorf("error writing JSON file: %v", err)
	}
	return nil
}
