package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isDigit(char byte) bool {
	if char >= '0' && char <= '9' {
		return true
	}

	return false
}

func scrapeNumber(lines [][]string, row, col int, cache []string) (int, []string) {
	ch := lines[row][col]

	if ch == "." || !isDigit(byte(ch[0])) {
		return 0, cache
	}

	// check cache if this index is already used
	for _, num := range cache {
		if num == strconv.Itoa(row)+strconv.Itoa(col) {
			return 0, cache
		}
	}

	// scrape: start
	start := col
	for j := col; j >= 0; j-- {
		if isDigit(byte(lines[row][j][0])) {
			if j == 0 {
				start = j
			}

			continue
		}

		start = j + 1
		break
	}

	// scrape: end
	end := start
	for i := start; i <= len(lines[row])-1; i++ {
		if isDigit(byte(lines[row][i][0])) {
			// add to cache
			cache = append(cache, strconv.Itoa(row)+strconv.Itoa(i))

			if i == len(lines[row])-1 {
				end = i+1
			}
			continue
		}

		end = i
		break
	}

	num, err := strconv.Atoi(strings.Join(lines[row][start:end], ""))

	if err != nil {
		panic(err)
	}

	return num, cache
}

func getSumOfAdjacentData(lines [][]string, row, col int) int {
	cache := []string{}

	tl, cache := scrapeNumber(lines, row-1, col-1, cache)
	tm, cache := scrapeNumber(lines, row-1, col, cache)
	tr, cache := scrapeNumber(lines, row-1, col+1, cache)
	l, cache := scrapeNumber(lines, row, col-1, cache)
	r, cache := scrapeNumber(lines, row, col+1, cache)
	bl, cache := scrapeNumber(lines, row+1, col-1, cache)
	bm, cache := scrapeNumber(lines, row+1, col, cache)
	br, _ := scrapeNumber(lines, row+1, col+1, cache)

	return tl + tm + tr + l + r + bl + bm + br

}

func main() {
	data, err := os.ReadFile("/home/bipin/Documents/Github/advent_of_code_2023_go/cmd/day3/input")

	if err != nil {
		panic(err)
	}

	text := string(data)

	// split and tidy up data
	lines := [][]string{}
	for _, line := range strings.Split(text, "\n") {
		cleaned := strings.TrimSpace(line)
		lines = append(lines, strings.Split(cleaned, ""))
	}

	total := 0

	// iterate through each characters
	for i, line := range lines {
		for j, char := range line {
			if isDigit(byte(char[0])) || char == "." {
				continue
			}

			total += getSumOfAdjacentData(lines, i, j)
		}
	}

	fmt.Println("Total: ", total)
}
