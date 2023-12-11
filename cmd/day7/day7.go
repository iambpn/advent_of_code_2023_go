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

type SortedData struct {
	card_type         string
	max_similar_count int
	value             int
	key               string
}

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
		if max > val {
			max = val
		}
	}

	card_type := ""
	switch max {
	case 5:
		card_type = const_five
		break
	case 4:
		card_type = const_four
		break
	case 3:
		if len(counter) == 2 {
			card_type = const_fullHouse
		} else {
			card_type = const_three
		}
		break
	case 2:
		if len(counter) == 3 {
			card_type = const_twoPair
		} else {
			card_type = const_onePair
		}
		break
	default:
		card_type = const_highCard
		break
	}

	return card_type, max
}

func sortCards(cards map[string]int) {
	sortedCards := []SortedData{}

	for key, value := range cards {
		card_type, max := getTypeAndSimilarCount(key)
		for i := 0; i <= len(sortedCards); i++ {
			if len(sortedCards) <= 0 {
				sortedCards = append(sortedCards, SortedData{
					card_type:         card_type,
					max_similar_count: max,
					value:             value,
					key:               key,
				})
			}

			card := sortedCards[i]
			nextCard := sortedCards[i+1]
			if card.max_similar_count == max {
				if card.card_type == card_type {
					// compare word by word
				} else {

				}
			}
		}
	}
}

func part1(cards map[string]int) int {
	// sortedCards := map[string]int{}
	getTypeAndSimilarCount("bipin")

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
