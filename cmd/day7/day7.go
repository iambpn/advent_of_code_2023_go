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

// Part 1
// func getTypeAndSimilarCount(word string) (string, int) {
// 	counter := map[rune]int{}
// 	for _, ch := range word {
// 		counter[ch] += 1
// 	}

// 	max := 0
// 	for _, val := range counter {
// 		if max < val {
// 			max = val
// 		}
// 	}

// 	card_type := ""
// 	switch max {
// 	case 5:
// 		card_type = const_five
// 	case 4:
// 		card_type = const_four
// 	case 3:
// 		if len(counter) == 2 {
// 			card_type = const_fullHouse
// 		} else {
// 			card_type = const_three
// 		}
// 	case 2:
// 		if len(counter) == 3 {
// 			card_type = const_twoPair
// 		} else {
// 			card_type = const_onePair
// 		}
// 	default:
// 		card_type = const_highCard
// 	}

// 	return card_type, max
// }

// Part 2
func getTypeAndSimilarCount(word string) (string, int) {
	counter := map[rune]int{}
	for _, ch := range word {
		counter[ch] += 1
	}

	max := 0
	maxRune := '0'
	for ch, val := range counter {
		if max < val {
			max = val
			maxRune = ch
		}
	}

	if maxRune == 'J' {
		// get second max number
		innerMax := 0
		counter['J'] = 0
		for _, val := range counter {
			if innerMax < val {
				innerMax = val
			}
		}

		max += innerMax
	} else {
		max += counter['J']
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

// Part 1
// var cardRank = map[rune]int{
// 	'2': 1,
// 	'3': 2,
// 	'4': 3,
// 	'5': 4,
// 	'6': 5,
// 	'7': 6,
// 	'8': 7,
// 	'9': 8,
// 	't': 9,
// 	'j': 10,
// 	'q': 11,
// 	'k': 12,
// 	'a': 13,
// }

// Part 2
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
	'j': 0,
	'q': 11,
	'k': 12,
	'a': 13,
}

func compareCard(a, b string) int {
	for i, _ := range a {
		if a[i] == b[i] {
			continue
		}

		if cardRank[rune(strings.ToLower(a)[i])] > cardRank[rune(strings.ToLower(b)[i])] {
			return 1
		} else {
			return -1
		}
	}

	return -1
}

func groupCards(cards map[string]int) map[string][]string {
	groupCards := map[string][]string{}

	for key, _ := range cards {
		card_type, _ := getTypeAndSimilarCount(key)
		groupCards[card_type] = append(groupCards[card_type], key)
	}
	return groupCards
}

func sortCards(groupedCards map[string][]string) map[string][]string {
	for group, cards := range groupedCards {
		sortedCards := make([]string, len(cards))
		copy(sortedCards, cards)

		i, j := 0, 0
		for i = 1; i < len(sortedCards); i++ {
			key := sortedCards[i]
			j = i - 1

			for j >= 0 && compareCard(sortedCards[j], key) > 0 {
				// swap
				sortedCards[j+1] = sortedCards[j]
				j -= 1
			}

			sortedCards[j+1] = key
		}
		groupedCards[group] = sortedCards
	}

	return groupedCards
}

func concatGroupedSortedCards(cards map[string][]string) []string {
	sortedCards := []string{}

	if val, ok := cards[const_highCard]; ok {
		sortedCards = append(sortedCards, val...)
	}

	if val, ok := cards[const_onePair]; ok {
		sortedCards = append(sortedCards, val...)
	}

	if val, ok := cards[const_twoPair]; ok {
		sortedCards = append(sortedCards, val...)
	}

	if val, ok := cards[const_three]; ok {
		sortedCards = append(sortedCards, val...)
	}

	if val, ok := cards[const_fullHouse]; ok {
		sortedCards = append(sortedCards, val...)
	}

	if val, ok := cards[const_four]; ok {
		sortedCards = append(sortedCards, val...)
	}

	if val, ok := cards[const_five]; ok {
		sortedCards = append(sortedCards, val...)
	}

	return sortedCards
}

func part1(cards map[string]int) int {
	groupedCards := groupCards(cards)
	groupedSortedCards := sortCards(groupedCards)
	sortedCards := concatGroupedSortedCards(groupedSortedCards)

	total := 0
	for i, key := range sortedCards {
		mult := cards[key] * (i + 1)
		total += mult
	}

	return total
}

func part2(cards map[string]int) int {
	groupedCards := groupCards(cards)
	groupedSortedCards := sortCards(groupedCards)
	sortedCards := concatGroupedSortedCards(groupedSortedCards)

	total := 0
	for i, key := range sortedCards {
		mult := cards[key] * (i + 1)
		total += mult
	}

	return total
}

func main() {
	data, err := os.ReadFile("/Users/imac/Documents/Github/advent_of_code_2023_go/cmd/day7/input")
	if err != nil {
		panic(err)
	}

	text := string(data)

	parsedData := parseData(text)

	// total := part1(parsedData)
	total := part2(parsedData)

	fmt.Println("Total: ", total)
}
