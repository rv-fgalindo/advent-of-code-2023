package main

import (
	"aoc/utilities"
	"fmt"
	"strconv"
)

func main() {
	Part1(false)
	Part2(false)
}

func Part1(debug bool) {
	fmt.Println("Starting Part 1")

	path := "input.txt"

	lines, err := utilities.LoadFile(path)
	if err != nil {
		fmt.Println("Error loading file:", err)
		return
	}
	sumParts := 0
	for lineIdx, line := range lines {
		for charIdx := 0; charIdx < len(line); charIdx++ {
			char := line[charIdx : charIdx+1]
			if char != "." {
				// moved it one step below so that we are over calling the conversion function when we know it's a period
				// created separate function to check if it's a number
				// will be used in the future
				if !isNumeric(char) {
					// found a symbol
					sumParts += checkAdjacent(lines, lineIdx, charIdx)
					if debug {
						fmt.Printf("char: %v \n", char)
					}
				}
			}
		}
	}
	fmt.Printf("sumParts: %v \n", sumParts)
}

// will check if the char is a number
func isNumeric(char string) bool {
	_, err := strconv.Atoi(char)
	if err != nil {
		// found a symbol or period
		return false
	}
	return true
}

// will convert a string to an int
// if error returns 0
func convertNum(chars string) int {
	num, err := strconv.Atoi(chars)
	if err != nil {
		fmt.Printf("Error converting string to int:\nString Passed in%v\nError:%v", chars, err)
		return 0
	}
	return num
}

// checks all of the surrounding chars and will return the sum of those nums around a symbol
// *** lines[lineIdx-1]
// *i* lines[lineIdx][charIdx]
// *** lines[lineIdx+1]
func checkAdjacent(lines []string, lineIdx int, charIdx int) int {
	uniqueNums := make(map[int]bool)
	// check top left
	topLeft := findNumberAdjacent(lines, lineIdx-1, charIdx-1)
	uniqueNums[topLeft] = true
	// check top
	top := findNumberAdjacent(lines, lineIdx-1, charIdx)
	uniqueNums[top] = true
	// check top right
	topRight := findNumberAdjacent(lines, lineIdx-1, charIdx+1)
	uniqueNums[topRight] = true
	// check left
	left := findNumberAdjacent(lines, lineIdx, charIdx-1)
	uniqueNums[left] = true
	// check right
	right := findNumberAdjacent(lines, lineIdx, charIdx+1)
	uniqueNums[right] = true
	// check bottom left
	bottomLeft := findNumberAdjacent(lines, lineIdx+1, charIdx-1)
	uniqueNums[bottomLeft] = true
	// check bottom
	bottom := findNumberAdjacent(lines, lineIdx+1, charIdx)
	uniqueNums[bottom] = true
	// check bottom right
	bottomRight := findNumberAdjacent(lines, lineIdx+1, charIdx+1)
	uniqueNums[bottomRight] = true
	sum := 0
	for num := range uniqueNums {
		sum += num

	}
	return sum
}

// given a char will check if it's a number and find the complete number
func findNumberAdjacent(lines []string, lineIdx int, charIdx int) int {
	firstChar := lines[lineIdx][charIdx : charIdx+1]
	if firstChar == "." {
		return 0
	} else if isNumeric(firstChar) {
		// we will want to check left and right for the full number until we hit a non-number on both sides
		// will create a loop while leftStop and rightStop are false
		// when we hit a nonnumber for left (will add one back to index) or right (will subtract one back ) will set the break to true and loop will eventually break
		leftStop := false
		rightStop := false
		leftIdx := charIdx - 1
		rightIdx := charIdx + 1
		////////
		////////
		// will want to add a condition to break if left goes negative or right goes past the length of the line
		////////
		////////
		for !leftStop || !rightStop {
			// check left
			if !leftStop {
				if leftIdx < 0 {
					leftStop = true
					leftIdx++
					continue
				}
				leftChar := lines[lineIdx][leftIdx : leftIdx+1]
				if isNumeric(leftChar) {
					leftIdx--
				} else {
					leftStop = true
					leftIdx++
				}
			}

			// check right
			if !rightStop {
				if rightIdx >= len(lines[lineIdx]) {
					rightStop = true
					rightIdx--
					continue
				}
				rightChar := lines[lineIdx][rightIdx : rightIdx+1]
				if isNumeric(rightChar) {
					rightIdx++
				} else {
					rightStop = true
					rightIdx--
				}
			}
		}
		// now we have the left and right index of the full number
		num := convertNum(lines[lineIdx][leftIdx : rightIdx+1])
		return num

	}
	return 0
}

func checkAdjacent2(lines []string, lineIdx int, charIdx int) int {
	uniqueNums := make(map[int]bool)
	// check top left
	topLeft := findNumberAdjacent2(lines, lineIdx-1, charIdx-1)
	uniqueNums[topLeft] = true
	// check top
	top := findNumberAdjacent2(lines, lineIdx-1, charIdx)
	uniqueNums[top] = true
	// check top right
	topRight := findNumberAdjacent2(lines, lineIdx-1, charIdx+1)
	uniqueNums[topRight] = true
	// check left
	left := findNumberAdjacent2(lines, lineIdx, charIdx-1)
	uniqueNums[left] = true
	// check right
	right := findNumberAdjacent2(lines, lineIdx, charIdx+1)
	uniqueNums[right] = true
	// check bottom left
	bottomLeft := findNumberAdjacent2(lines, lineIdx+1, charIdx-1)
	uniqueNums[bottomLeft] = true
	// check bottom
	bottom := findNumberAdjacent2(lines, lineIdx+1, charIdx)
	uniqueNums[bottom] = true
	// check bottom right
	bottomRight := findNumberAdjacent2(lines, lineIdx+1, charIdx+1)
	uniqueNums[bottomRight] = true
	sum := 0

	// by default we will return 1 if nothing
	if len(uniqueNums) > 2 {
		// because we only want gear ratio if another part is present we will multiply by
		// map {1,3,} want to return 0
		// map {1,3,5} want to return 15
		temp := 1
		for num := range uniqueNums {
			temp *= num
		}
		sum = temp
	}
	return sum
}

func findNumberAdjacent2(lines []string, lineIdx int, charIdx int) int {
	firstChar := lines[lineIdx][charIdx : charIdx+1]
	if firstChar == "." {
		return 1
	} else if isNumeric(firstChar) {
		// we will want to check left and right for the full number until we hit a non-number on both sides
		// will create a loop while leftStop and rightStop are false
		// when we hit a nonnumber for left (will add one back to index) or right (will subtract one back ) will set the break to true and loop will eventually break
		leftStop := false
		rightStop := false
		leftIdx := charIdx - 1
		rightIdx := charIdx + 1
		////////
		////////
		// will want to add a condition to break if left goes negative or right goes past the length of the line
		////////
		////////
		for !leftStop || !rightStop {
			// check left
			if !leftStop {
				if leftIdx < 0 {
					leftStop = true
					leftIdx++
					continue
				}
				leftChar := lines[lineIdx][leftIdx : leftIdx+1]
				if isNumeric(leftChar) {
					leftIdx--
				} else {
					leftStop = true
					leftIdx++
				}
			}

			// check right
			if !rightStop {
				if rightIdx >= len(lines[lineIdx]) {
					rightStop = true
					rightIdx--
					continue
				}
				rightChar := lines[lineIdx][rightIdx : rightIdx+1]
				if isNumeric(rightChar) {
					rightIdx++
				} else {
					rightStop = true
					rightIdx--
				}
			}
		}
		// now we have the left and right index of the full number
		num := convertNum(lines[lineIdx][leftIdx : rightIdx+1])
		return num

	}
	return 1
}

func Part2(debug bool) {
	fmt.Println("Starting Part 2")

	path := "input.txt"

	lines, err := utilities.LoadFile(path)
	if err != nil {
		fmt.Println("Error loading file:", err)
		return
	}
	sumParts := 0
	for lineIdx, line := range lines {
		for charIdx := 0; charIdx < len(line); charIdx++ {
			char := line[charIdx : charIdx+1]
			// we ar looking for gear so will only check for gear
			if char == "*" {
				// found a symbol
				sumParts += checkAdjacent2(lines, lineIdx, charIdx)
				if debug {
					fmt.Printf("char: %v \n", char)
				}
			}
		}
	}
	fmt.Printf("sumParts: %v \n", sumParts)
}
