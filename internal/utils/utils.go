package utils

import (
	"fmt"
	"os"
	"path"
)

func WriteToFile(filename string, content string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home dir", err)
		return
	}
	path := path.Join(homeDir, filename)
	fmt.Println("path", path)
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening or creating the file: %v\n", err)
		return
	}
	defer file.Close()

	// Write content to the file
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Printf("Error writing to the file: %v\n", err)
		return
	}
}
