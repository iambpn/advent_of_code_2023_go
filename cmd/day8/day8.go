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

func findDirectionPart2(maps map[string][2]string, nextDestination string, directions []string, nextDirectionIdx int) int {
	if nextDestination[len(nextDestination)-1] == 'Z' {
		return 0
	}

	if nextDirectionIdx == len(directions) {
		nextDirectionIdx = 0
	}

	idx := 0
	if directions[nextDirectionIdx] == "R" {
		idx = 1
	}

	return 1 + findDirectionPart2(maps, maps[nextDestination][idx], directions, nextDirectionIdx+1)
}

func part1(text string) int {
	directions, destinationMaps := parseData(text)

	steps := findDirection(destinationMaps, "AAA", directions, 0)

	return steps
}

func getGreatestCommonDivisor(x, y int) int {
	for y != 0 {
		tmp := x
		x = y
		y = tmp % y
	}

	return x
}

func getLcm(steps []int) int {
	lcm := 1
	for _, num := range steps {
		gcd := getGreatestCommonDivisor(lcm, num)
		lcm = (lcm * num) / gcd
	}
	return lcm
}

func part2(text string) int {
	directions, destinationMaps := parseData(text)

	nextDestinations := []string{}

	for key, _ := range destinationMaps {
		if key[len(key)-1] == 'A' {
			nextDestinations = append(nextDestinations, key)
		}
	}

	firstSteps := []int{}
	for _, dest := range nextDestinations {
		step := findDirectionPart2(destinationMaps, dest, directions, 0)
		firstSteps = append(firstSteps, step)
	}

	return getLcm(firstSteps)
}

func main() {
	data, err := os.ReadFile("/home/bipin/Documents/Github/advent_of_code_2023_go/cmd/day8/input")

	if err != nil {
		panic(err)
	}

	text := string(data)
	// total := part1(text)
	total := part2(text)

	fmt.Println("total steps: ", total)
}
