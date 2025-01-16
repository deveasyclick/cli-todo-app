package file_service

import (
	"bufio"
	"fmt"
	"os"
)

// SaveToFile writes encrypted data to a file.
func SaveToFile(filePath, token string) error {
	// Open the file for writing, create it if it doesn't exist, and truncate it if it does
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	// Write the email and token, each on a new line
	_, err = fmt.Fprintf(file, "%s\n", token)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	return nil
}

// ReadFromFile reads encrypted data from a file.
func ReadFromFile(filePath string) (string, error) {
	// Open the file for reading
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	// Use a scanner to read lines
	scanner := bufio.NewScanner(file)

	// Read the email and token
	var token string

	if scanner.Scan() {
		token = scanner.Text()
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	return token, nil
}

// DeleteFile deletes the file storing the data.
func DeleteFile(filename string) error {
	return os.Remove(filename)
}
