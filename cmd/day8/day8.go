package main

import (
	"fmt"
	"os"
	"strings"
)

func parseData(text string) ([]string, map[string][2]string) {
	data := strings.Split(text, "\n")
	textDirections := data[0]
	destinationMaps := data[2:]

	directions := strings.Split(textDirections, "")
	parsedDestinationMaps := map[string][2]string{}

	for _, direction := range destinationMaps {
		keyVal := strings.Split(direction, " = ")
		values := strings.Split(keyVal[1], ", ")
		parsedDestinationMaps[keyVal[0]] = [2]string{
			values[0][1:],
			values[1][:len(values[1])-1],
		}
	}

	return directions, parsedDestinationMaps
}

func findDirection(maps map[string][2]string, nextDestination string, directions []string, nextDirectionIdx int) int {
	if nextDestination == "ZZZ" {
		return 0
	}

	if nextDirectionIdx == len(directions) {
		nextDirectionIdx = 0
	}

	idx := 0
	if directions[nextDirectionIdx] == "R" {
		idx = 1
	}

	return 1 + findDirection(maps, maps[nextDestination][idx], directions, nextDirectionIdx+1)
}

func part1(text string) int {
	directions, destinationMaps := parseData(text)

	steps := findDirection(destinationMaps, "AAA", directions, 0)

	return steps
}

func part2(text string) int {
	directions, destinationMaps := parseData(text)

	nextDestinations := []string{}

	for key, _ := range destinationMaps {
		if key[len(key)-1] == 'A' {
			nextDestinations = append(nextDestinations, key)
		}
	}

	i := 0
	steps := 0
	for {
		destZCount := 0
		for _, key := range nextDestinations {
			if key[len(key)-1] == 'Z' {
				destZCount += 1
			}
		}

		if destZCount == len(nextDestinations) {
			break
		}

		if i == len(directions) {
			i = 0
		}

		idx := 0
		if directions[i] == "R" {
			idx = 1
		}

		nextDest := []string{}
		for _, key := range nextDestinations {
			nextDest = append(nextDest, destinationMaps[key][idx])
		}

		nextDestinations = nextDest

		steps += 1
		i += 1
	}

	return steps
}

func main() {
	data, err := os.ReadFile("/home/bipin/Documents/Github/advent_of_code_2023_go/cmd/day8/input.example3")

	if err != nil {
		panic(err)
	}

	text := string(data)
	// total := part1(text)
	total := part2(text)

	fmt.Println("total steps: ", total)
}
