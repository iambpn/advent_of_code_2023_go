package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	const_seed                  = "seeds"
	const_seedToSoil            = "seed-to-soil"
	const_soilToFertilizer      = "soil-to-fertilizer"
	const_fertilizerToWater     = "fertilizer-to-water"
	const_waterToLight          = "water-to-light"
	const_lightToTemperature    = "light-to-temperature"
	const_temperatureToHumidity = "temperature-to-humidity"
	const_humidityToLocation    = "humidity-to-location"
)

func cleanData(text string) map[string][][]int {
	lines := strings.Split(text, "\n")

	valueMap := map[string][][]int{}

	name := ""
	for idx, line := range lines {
		value := line
		if idx == 0 {
			value = strings.Split(line, ":")[1]
			name = strings.TrimSpace(strings.Split(line, ":")[0])
		}

		if len(line) == 0 {
			continue
		}

		if line[len(line)-1] == ':' {
			name = strings.TrimSpace(strings.Split(line, " ")[0])
			continue
		}

		values := strings.Split(strings.TrimSpace(value), " ")

		numValues := []int{}
		for _, val := range values {
			num, err := strconv.Atoi(val)

			if err != nil {
				panic(err)
			}

			numValues = append(numValues, num)
		}

		valueMap[name] = append(valueMap[name], numValues)
	}

	return valueMap
}

func getMapValue(key int, values [][]int) int {
	for _, row := range values {
		dest := row[0]
		src := row[1]
		length := row[2]

		if key >= src && key <= src+length {
			if key > src {
				return dest + (key - src)
			}

			return dest + (src - key)
		}
	}

	return key
}

func part1(text string) int {
	input := cleanData(text)

	minLocation := math.MaxInt
	for _, num := range input[const_seed][0] {
		location := getMapValue(
			getMapValue(
				getMapValue(
					getMapValue(
						getMapValue(
							getMapValue(
								getMapValue(num, input[const_seedToSoil]),
								input[const_soilToFertilizer]),
							input[const_fertilizerToWater]),
						input[const_waterToLight]),
					input[const_lightToTemperature]),
				input[const_temperatureToHumidity]),
			input[const_humidityToLocation])

		if location <= minLocation {
			minLocation = location
		}
	}

	return minLocation
}

func part2(text string) int {
	input := cleanData(text)

	minLocation := math.MaxInt
	nums := input[const_seed][0]

	for i := 0; i < len(nums); i += 2 {
		start := nums[i]
		length := nums[i+1]

		for j := start; j <= start+length; j++ {
			location := getMapValue(
				getMapValue(
					getMapValue(
						getMapValue(
							getMapValue(
								getMapValue(
									getMapValue(j, input[const_seedToSoil]),
									input[const_soilToFertilizer]),
								input[const_fertilizerToWater]),
							input[const_waterToLight]),
						input[const_lightToTemperature]),
					input[const_temperatureToHumidity]),
				input[const_humidityToLocation])

			if location <= minLocation {
				minLocation = location
			}
		}
	}

	return minLocation
}

func main() {
	data, err := os.ReadFile("/home/bipin/Documents/Github/advent_of_code_2023_go/cmd/day5/input")

	if err != nil {
		panic(err)
	}

	text := string(data)

	res := part2(text)

	fmt.Println("Min location: ", res)
}
