package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func cleanData(text string) ([][]int, [][]int) {
	lines := strings.Split(text, "\n")

	winNumbers := [][]int{}
	yourNumber := [][]int{}

	for idx, line := range lines {
		card := strings.Split(line, ":")
		nums := strings.Split(card[1], "|")

		winNumbers = append(winNumbers, []int{})
	winScope:
		for _, el := range strings.Split(nums[0], " ") {
			if el == "" {
				continue winScope
			}
			num, err := strconv.Atoi(strings.TrimSpace(el))

			if err != nil {
				panic(err)
			}

			winNumbers[idx] = append(winNumbers[idx], num)
		}

		yourNumber = append(yourNumber, []int{})
	yourScope:
		for _, el := range strings.Split(nums[1], " ") {
			if el == "" {
				continue yourScope
			}

			num, err := strconv.Atoi(strings.TrimSpace(el))

			if err != nil {
				panic(err)
			}

			yourNumber[idx] = append(yourNumber[idx], num)
		}
	}

	return winNumbers, yourNumber
}

func getWinningNumbers(win, your [][]int) [][]int {
	wonNumbers := [][]int{}

	for i, nums := range your {
		wonNumber := []int{}
		for _, num := range nums {
		winLoop:
			for _, winNum := range win[i] {
				if num == winNum {
					wonNumber = append(wonNumber, winNum)
					break winLoop
				}
			}
		}
		wonNumbers = append(wonNumbers, wonNumber)
	}

	return wonNumbers
}

func part1(text string) int {
	win, your := cleanData(text)

	wonNumbers := getWinningNumbers(win, your)

	total := 0

	for _, wonNumber := range wonNumbers {
		value := 0

		//  initially assign 1
		if len(wonNumber) > 0 {
			value = 1
		}

		// from next index multiply by 2
		for i := 1; i < len(wonNumber); i++ {
			value = value * 2
		}

		total += value
	}

	return total
}

func initializeValue(val int, ok bool) int {
	if ok {
		return val + 1
	}

	return 1
}

func recursiveCount(wonNumbers [][]int, cardCounts map[int]int, row int) map[int]int {
	for i := 1; i <= len(wonNumbers[row]); i++ {
		val, ok := cardCounts[row+i]
		cardCounts[row+i] = initializeValue(val, ok)
		cardCounts = recursiveCount(wonNumbers, cardCounts, row+i)
	}

	return cardCounts
}

func part2(text string) int {
	win, your := cleanData(text)

	wonNumbers := getWinningNumbers(win, your)

	cardCounts := map[int]int{}

	for idx, _ := range wonNumbers {
		val, ok := cardCounts[idx]
		cardCounts[idx] = initializeValue(val, ok)
		cardCounts = recursiveCount(wonNumbers, cardCounts, idx)
	}

	total := 0
	for _, val := range cardCounts {
		total += val
	}

	return total
}

func main() {
	data, err := os.ReadFile("/home/bipin/Documents/Github/advent_of_code_2023_go/cmd/day4/input")

	if err != nil {
		panic(err)
	}

	text := string(data)

	// total := part1(text)
	total := part2(text)

	fmt.Println("total: ", total)
}
