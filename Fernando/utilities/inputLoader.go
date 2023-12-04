package utilities

import (
	"bufio"
	"fmt"
	"os"
)

// LoadFile loads a file into a slice of strings given a path
// returns a slice of strings and an error (if any)
func LoadFile(path string) ([]string, error) {
	// Open the file
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	// Create a scanner to read the file
	scanner := bufio.NewScanner(file)

	// Create a slice to hold the lines
	var lines []string

	// Iterate through each line in the file
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	// Check for any errors that may have occurred during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}
	return lines, nil
}