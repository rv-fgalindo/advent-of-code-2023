package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	msg := "Hello World"
	fmt.Println(msg)
	// // fmt.Println(string(msg[1]))
	// return

	// Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
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
		return
	}

	// Print the lines or use them as needed
	fmt.Println("Lines in the file:")
	// fmt.Printf("index 0\n%v", lines[0])
	for i, line := range lines {
		first := "temp"
		last := "temp"
		for _, unicodeChar := range line {
			// converting unicode value to string
			char := string(unicodeChar)
			if intValue, err := strconv.Atoi(char); err == nil {
				last = char
				if first == "temp" {
					first = char
				}
				fmt.Printf("Successfully converted '%v' to integer: %d\n", char, intValue)
			}
		}
		fmt.Printf("Iteration %v\nLine %v\nFirst %v\nLast %v\n", i, line, first, last)
		lines[i] = first + last
		// line = first + last
	}
	// tracking sum
	sum := 0
	for i, line := range lines {
		if intValue, err := strconv.Atoi(line); err == nil {
			sum += intValue
		} else {
			fmt.Println("Error on line ", i, err)
		}
	}
	fmt.Printf("\n\nFinal Value: %v\n", sum)
}
