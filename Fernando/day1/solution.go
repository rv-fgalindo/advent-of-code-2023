package main

import (
	"aoc/utilities"
	"fmt"
	"strconv"
)

func main() {

	Part1()
	Part2()

}

func Part1(){
	msg := "Running Part 1"
	fmt.Println(msg)
	// how to get a specific char of a string (but is unicode value)
	// // fmt.Println(string(msg[1]))
	// return
	
	path := "input.txt"
	lines, err := utilities.LoadFile(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	// Print the lines or use them as needed
	// fmt.Println("Lines in the file:")
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
				// below line is to trick the compiler into thinking we are using the variable but only using in commented out print statement below
				_ = intValue
				// fmt.Printf("Successfully converted '%v' to integer: %d\n", char, intValue)
			}
		}
		// uncomment line below to see each iteration
		// fmt.Printf("Iteration %v\nLine %v\nFirst %v\nLast %v\n", i, line, first, last)
		
		// because go is pass by value, we need to update the original slice and not line (which is a copy of the element)
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
	fmt.Printf("Part 1 Final Value: %v\n\n", sum)
}

func Part2(){
	msg := "Running Part 2"
	fmt.Println(msg)
	// how to get a specific char of a string (but is unicode value)
	// // fmt.Println(string(msg[1]))
	// return

	// Open the file
	path := "input.txt"
	lines, err := utilities.LoadFile(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	
	// fmt.Printf("index 0\n%v", lines[0])
	for i, line := range lines {
		// limit iterations to debug
		// if i>=2 {break}
		first := "temp"
		last := "temp"
		length := len(line)
		for index, unicodeChar := range line {
			// converting unicode value to string
			char := string(unicodeChar)
			if intValue, err := strconv.Atoi(char); err == nil {
				last = char
				if first == "temp" {
					first = char
				}
				// below line is to trick the compiler into thinking we are using the variable but only using in commented out print statement below
				_ = intValue
				// fmt.Printf("Successfully converted '%v' to integer: %d\n", char, intValue)
			} else if length - index >= 3 {
				// create a substring for the next 3 chars 
				possibleDigit := line[index:index+3]
				// fmt.Printf("Possible Digit 3 letter: %v\n", possibleDigit)
				// create a switch case and then compare to 3 letter digits
				switch possibleDigit {
					// if match, update first and last accordingly
					case "one":
						// fmt.Printf("Match Found!\nIteration %v\nLine %v\nPossibleDigit %v\n", i, line, possibleDigit)
						last = "1"
						if first == "temp" {first = "1"}
					case "two":
						// fmt.Printf("Match Found!\nIteration %v\nLine %v\nPossibleDigit %v\n", i, line, possibleDigit)
						last = "2"
						if first == "temp" {first = "2"}
					case "six":
						// fmt.Printf("Match Found!\nIteration %v\nLine %v\nPossibleDigit %v\n", i, line, possibleDigit)
						last = "6"
						if first == "temp" {first = "6"}
				}		
				// if no match, check if four letters left if NOT continue to next loop iteration
				// if four letters left, create a substring for the next 4 chars
				if length - index >= 4 {
					// create a switch case and then compare to 4 letter digits
					possibleDigit = line[index:index+4]
					// fmt.Printf("Possible Digit 4 letter: %v\n", possibleDigit)
					switch possibleDigit {
						// if match, update first and last accordingly
						case "four":
							// fmt.Printf("Match Found!\nIteration %v\nLine %v\nPossibleDigit %v\n", i, line, possibleDigit)
							last = "4"
							if first == "temp" {first = "4"}
						case "five":
							// fmt.Printf("Match Found!\nIteration %v\nLine %v\nPossibleDigit %v\n", i, line, possibleDigit)
							last = "5"
							if first == "temp" {first = "5"}
						case "nine":
							// fmt.Printf("Match Found!\nIteration %v\nLine %v\nPossibleDigit %v\n", i, line, possibleDigit)
							last = "9"
							if first == "temp" {first = "9"}
					}
					if length - index >= 5 {
						// create a switch case and then compare to 5 letter digits
						possibleDigit = line[index:index+5]
						// fmt.Printf("Possible Digit 5 letter: %v\n", possibleDigit)
						switch possibleDigit {
							// if match, update first and last accordingly
							case "three":
								// fmt.Printf("Match Found!\nIteration %v\nLine %v\nPossibleDigit %v\n", i, line, possibleDigit)
								last = "3"
								if first == "temp" {first = "3"}
							case "seven":
								// fmt.Printf("Match Found!\nIteration %v\nLine %v\nPossibleDigit %v\n", i, line, possibleDigit)
								last = "7"
								if first == "temp" {first = "7"}
							case "eight":
								// fmt.Printf("Match Found!\nIteration %v\nLine %v\nPossibleDigit %v\n", i, line, possibleDigit)
								last = "8"
								if first == "temp" {first = "8"}
						}
					}
				}
			}
		}
		// uncomment line below to see each iteration
		// fmt.Printf("Iteration %v\nLine %v\nFirst %v\nLast %v\n", i, line, first, last)
		
		// because go is pass by value, we need to update the original slice and not line (which is a copy of the element)
		lines[i] = first + last
		// fmt.Printf("\n\nUpdating Lines[i]!\nLine %v\nFirst %v\nLast %v\n\n", line, first, last)
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
	fmt.Printf("Part 2 Final Value: %v\n", sum)
}
