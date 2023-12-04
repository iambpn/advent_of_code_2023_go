package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isDigit(char byte) bool {
	if char >= 48 && char <= 57 {
		return true
	}

	return false
}

func is3CharWord(word string) (string, bool) {
	switch strings.ToLower(word) {
	case "one":
		return "1", true
	case "two":
		return "2", true
	case "six":
		return "6", true
	default:
		return "0", false
	}
}

func is5CharWord(word string) (string, bool) {
	switch strings.ToLower(word) {
	case "seven":
		return "7", true
	case "eight":
		return "8", true
	case "three":
		return "3", true
	default:
		return "0", false
	}
}

func is4CharWord(word string) (string, bool) {
	switch strings.ToLower(word) {
	case "nine":
		return "9", true
	case "five":
		return "5", true
	case "four":
		return "4", true
	case "zero":
		return "0", true
	default:
		return "0", false
	}
}

func findFirstDigit(line string) string {
	chars := []byte(line)
	for idx, char := range chars {
		if isDigit(char) {
			return string(char)
		}

		if idx+5 < len(line) {
			if w, ok := is5CharWord(line[idx : idx+5]); ok {
				return w
			}
		}

		if idx+4 < len(line) {
			if w, ok := is4CharWord(line[idx : idx+4]); ok {
				return w
			}
		}

		if idx+3 < len(line) {
			if w, ok := is3CharWord(line[idx : idx+3]); ok {
				return w
			}
		}
	}

	return "0"
}

func findLastDigit(line string) string {
	lineChars := []byte(line)

	for i := len(lineChars) - 1; i >= 0; i-- {
		if isDigit(lineChars[i]) {
			return string(lineChars[i])
		}

		if i-5 >= 0 {
			if w, ok := is5CharWord(line[i-4 : i+1]); ok {
				return w
			}
		}

		if i-4 >= 0 {
			if w, ok := is4CharWord(line[i-3 : i+1]); ok {
				return w
			}
		}

		if i-3 >= 0 {
			if w, ok := is3CharWord(line[i-2 : i+1]); ok {
				return w
			}
		}
	}

	return "0"
}

func main() {
	data, err := os.ReadFile("/home/bipin/Documents/Github/advent_of_code_2023_go/cmd/day1/input.txt")

	if err != nil {
		panic(err)
	}

	inp := string(data)

	lines := strings.Split(inp, "\n")

	var total = 0

	for _, line := range lines {
		line := strings.TrimSpace(line)
		n1 := findFirstDigit(line)
		n2 := findLastDigit(line)

		val, err := strconv.Atoi(n1 + n2)

		if err != nil {
			panic(err)
		}

		total += val
	}

	fmt.Printf("Total calibration value: %v \n", total)
}
