package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	const_five      = "five"
	const_four      = "four"
	const_three     = "three"
	const_fullHouse = "full-house"
	const_twoPair   = "two-pair"
	const_onePair   = "one-pair"
	const_highCard  = "high-card"
)

func parseData(text string) map[string]int {
	lines := strings.Split(text, "\n")

	data := map[string]int{}

	for _, line := range lines {
		values := strings.Split(line, " ")
		num, err := strconv.Atoi(strings.TrimSpace(values[1]))

		if err != nil {
			panic(err)
		}

		data[strings.TrimSpace(values[0])] = num
	}

	return data
}

func getTypeAndSimilarCount(word string) (string, int) {
	counter := map[rune]int{}
	for _, ch := range word {
		counter[ch] += 1
	}

	max := 0
	for _, val := range counter {
		if max < val {
			max = val
		}
	}

	fmt.Println(counter, max)

	card_type := ""
	switch max {
	case 5:
		card_type = const_five
	case 4:
		card_type = const_four
	case 3:
		if len(counter) == 2 {
			card_type = const_fullHouse
		} else {
			card_type = const_three
		}
	case 2:
		if len(counter) == 3 {
			card_type = const_twoPair
		} else {
			card_type = const_onePair
		}
	default:
		card_type = const_highCard
	}

	return card_type, max
}

func sortCards(cards map[string]int) {
	groupCards := map[string][]string{}

	for key, _ := range cards {
		card_type, _ := getTypeAndSimilarCount(key)
		groupCards[card_type] = append(groupCards[card_type], key)
	}

	// sort group cards
	for group, cardKeys := range groupCards {
		sorted := []string{}
		for i, key := range cardKeys {
			nextKey := ""
			if len(cardKeys) > i {
				nextKey = cardKeys[i+1]
			}
			
		}
	}
}

func part1(cards map[string]int) int {
	// sortedCards := map[string]int{}
	sortCards(cards)
	return 0
}

func main() {
	data, err := os.ReadFile("/Users/imac/Documents/Github/advent_of_code_2023_go/cmd/day7/input.example")

	if err != nil {
		panic(err)
	}

	text := string(data)

	parsedData := parseData(text)

	total := part1(parsedData)

	fmt.Print("Total: ", total)
}
