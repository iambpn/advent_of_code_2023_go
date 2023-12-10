package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// loop over all the inputs from mid point with early exit
func findRange(times []int, distances []int) [][]int {
	winRanges := [][]int{}

	for idx, time := range times {
		midIdx := time / 2
		winRange := []int{}

		// find upper bound
	upperBound:
		for i := midIdx; i <= time; i++ {
			diff := time - i
			distance := i * diff

			if distance > distances[idx] {
				winRange = append(winRange, distance)
			} else {
				break upperBound
			}
		}

		// find lower bound
	lowerBound:
		for i := midIdx - 1; i >= 0; i-- {
			diff := time - i
			distance := i * diff

			if distance > distances[idx] {
				winRange = append(winRange, distance)
			} else {
				break lowerBound
			}
		}

		winRanges = append(winRanges, winRange)
	}

	return winRanges
}

func cleanDataPart1(text string) ([]int, []int) {
	time := []int{}
	distance := []int{}

	for _, val := range strings.Split(strings.TrimSpace(strings.Split(strings.Split(text, "\n")[0], ":")[1]), " ") {
		if val == "" {
			continue
		}

		num, err := strconv.Atoi(strings.TrimSpace(val))
		if err != nil {
			panic(err)
		}
		time = append(time, num)
	}

	for _, val := range strings.Split(strings.TrimSpace(strings.Split(strings.Split(text, "\n")[1], ":")[1]), " ") {
		if val == "" {
			continue
		}

		num, err := strconv.Atoi(strings.TrimSpace(val))
		if err != nil {
			panic(err)
		}
		distance = append(distance, num)
	}

	return time, distance
}

func part1(text string) int {
	time, distance := cleanDataPart1(text)
	winRanges := findRange(time, distance)

	total := 1
	for _, winRange := range winRanges {
		total *= len(winRange)
	}

	return total
}

func cleanDataPart2(text string) (int, int) {
	lines := strings.Split(text, "\n")

	chars := strings.Split(strings.Split(lines[0], ":")[1], " ")
	value := ""

	for _, ch := range chars {
		value += strings.TrimSpace(ch)
	}

	time, err := strconv.Atoi(value)

	if err != nil {
		panic(err)
	}

	chars = strings.Split(strings.Split(lines[1], ":")[1], " ")
	value = ""

	for _, ch := range chars {
		value += strings.TrimSpace(ch)
	}

	distance, err := strconv.Atoi(value)

	if err != nil {
		panic(err)
	}

	return time, distance
}

func findPossibleWinCount(time, distance int) int {
	minMultiplier := 1
	minMultiplicand := 1

loweBound:
	for i := time/2 - 1; i >= 0; i-- {
		diff := time - i
		calcDistance := i * diff

		if calcDistance > distance {
			minMultiplier = i
			minMultiplicand = diff
		} else {
			break loweBound
		}
	}

	return (minMultiplicand - minMultiplier) + 1
}

func part2(text string) int {
	time, distance := cleanDataPart2(text)

	return findPossibleWinCount(time, distance)
}

func main() {
	data, err := os.ReadFile("/home/bipin/Documents/Github/advent_of_code_2023_go/cmd/day6/input")

	if err != nil {
		panic(err)
	}

	text := string(data)

	// total := part1(text)

	total := part2(text)

	fmt.Println("Total: ", total)
}
