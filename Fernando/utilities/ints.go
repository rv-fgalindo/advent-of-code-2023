package utilities

import (
	"fmt"
	"strconv"
)

// will check if the char is a number
func IsNumeric(char string) bool {
	_, err := strconv.Atoi(char)
	if err != nil {
		// found a symbol or period
		return false
	}
	return true
}

// will convert a string to an int
// if error returns 0
func ConvertNum(chars string) int {
	num, err := strconv.Atoi(chars)
	if err != nil {
		fmt.Printf("Error converting string to int:\nString Passed in%v\nError:%v", chars, err)
		return 0
	}
	return num
}