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

var cardRank = map[rune]int{
	'2': 1,
	'3': 2,
	'4': 3,
	'5': 4,
	'6': 5,
	'7': 6,
	'8': 7,
	'9': 8,
	't': 9,
	'j': 10,
	'q': 11,
	'k': 12,
	'a': 13,
}

func compareCard(a, b rune) int {
	if a == b {
		return 0
	}

	if cardRank[a] > cardRank[b] {
		return 1
	}

	return -1
}

func sortCards(cards map[string]int) {
	groupCards := map[string][]string{}

	for key, _ := range cards {
		card_type, _ := getTypeAndSimilarCount(key)
		groupCards[card_type] = append(groupCards[card_type], key)
	}

	// sort group cards
	
	fmt.Println(groupCards)
}

func part1(cards map[string]int) int {
	// sortedCards := map[string]int{}
	sortCards(cards)
	return 0
}

func main() {
	data, err := os.ReadFile("/home/bipin/Documents/Github/advent_of_code_2023_go/cmd/day7/input.example")
	if err != nil {
		panic(err)
	}

	text := string(data)

	parsedData := parseData(text)

	total := part1(parsedData)

	fmt.Print("Total: ", total)
}
