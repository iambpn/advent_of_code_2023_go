package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CubeCount struct {
	red   int32
	blue  int32
	green int32
}

func getHighestCubesFromEachGame(line string) CubeCount {
	count := CubeCount{
		red:   0,
		blue:  0,
		green: 0,
	}

	game := strings.Split(line, ":")[1]
	for _, round := range strings.Split(game, ";") {
		cubes := strings.Split(round, ",")
		for _, cube := range cubes {
			values := strings.Split(strings.TrimSpace(cube), " ")
			color := strings.TrimSpace(values[1])
			num, err := strconv.Atoi(strings.TrimSpace(values[0]))

			if err != nil {
				panic(err)
			}

			switch color {
			case "blue":
				{
					if count.blue < int32(num) {
						count.blue = int32(num)
					}
					break
				}
			case "red":
				{
					if count.red < int32(num) {
						count.red = int32(num)
					}
					break
				}
			case "green":
				{
					if count.green < int32(num) {
						count.green = int32(num)
					}
					break
				}
			}
		}
	}

	return count
}

func part1(textData string) int {
	lines := strings.Split(textData, "\n")

	total := 0
	for idx, line := range lines {
		cubeCount := getHighestCubesFromEachGame(line)

		// verify if game is possible
		if cubeCount.red > totalCube.red || cubeCount.blue > totalCube.blue || cubeCount.green > totalCube.green {
			continue
		}

		total += idx + 1
	}

	return total
}

func part2(textData string) int {
	lines := strings.Split(textData, "\n")

	total := 0
	for _, line := range lines {
		cubeCount := getHighestCubesFromEachGame(line)

		power := cubeCount.blue * cubeCount.green * cubeCount.red
		total += int(power)
	}

	return total
}

var totalCube = CubeCount{
	red:   12,
	blue:  14,
	green: 13,
}

func main() {
	data, err := os.ReadFile("/home/bipin/Documents/Github/advent_of_code_2023_go/cmd/day2/input.txt")

	if err != nil {
		panic(err)
	}

	textData := string(data)

	// total := part1(textData)
	total := part2(textData)

	fmt.Println("Total Game Sum:", total)

}
