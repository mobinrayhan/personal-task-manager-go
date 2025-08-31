package io

import (
	"encoding/json"
	"fmt"
	"os"
)

func WriteToFile(filename string, data interface{}) {
	jsonData, _ := json.MarshalIndent(data, "", " ")
	err := os.WriteFile(filename, jsonData, 0664)

	if err != nil {
		panic(err)
	}
}

func ReadFromFile(filename string, target any) error {
	data, err := os.ReadFile(filename)

	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	if len(data) == 0 {
		return nil
	}

	if err := json.Unmarshal(data, target); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return nil
}
