package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	file, _ := os.Open("input.txt")
	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	solution1 := solve(input, false)
	fmt.Print("Part 1: ")
	fmt.Println(solution1)

	solution2 := solve(input, true)
	fmt.Print("Part 2: ")
	fmt.Println(solution2)
}

func solve(input []string, isPartTwo bool) any {

	type Hand struct {
		scores [5]int
		bid    int
	}

	winnings := 0
	hands := make([]Hand, 0)

	// high card 0, pair 100, two pair 200, trio 300, fullhouse 400, poker 500, five of a kind 600
	scores := map[int]int{1: 0, 2: 100, 3: 300, 4: 500, 5: 600}
	cardValue := map[rune]int{'A': 13, 'K': 12, 'Q': 11, 'J': 10, 'T': 9, '9': 8, '8': 7, '7': 6, '6': 5, '5': 4, '4': 3, '3': 2, '2': 1}

	if isPartTwo {
		cardValue['J'] = 0
	}

	for _, line := range input {
		hand := strings.Split(line, " ")
		handMap := map[rune]int{}
		bid, _ := strconv.Atoi(hand[1])
		newHand := Hand{[5]int{0}, bid}

		// process type of hand and add tiebreaker score
		for i, card := range hand[0] {
			handMap[card]++
			newHand.scores[i] = cardValue[card]
		}

		// replace jokers by the most common card
		if isPartTwo {
			mostCommonCard := '0'
			mostCommonCardAmount := 0
			for card, cardAmount := range handMap {
				if cardAmount > mostCommonCardAmount && card != 'J' {
					mostCommonCardAmount = cardAmount
					mostCommonCard = card
				}
			}
			handMap[mostCommonCard] += handMap['J']
			handMap['J'] = 0
		}

		for _, card := range handMap {
			newHand.scores[0] += scores[card]
		}
		hands = append(hands, newHand)
	}

	// sort hands by best hand and then by high card
	sort.Slice(hands, func(a int, b int) bool {
		i := 0
		for i = 0; hands[a].scores[i] == hands[b].scores[i]; i++ {
		}
		return hands[a].scores[i] < hands[b].scores[i]
	})

	for i, hand := range hands {
		winnings += hand.bid * (i + 1)
	}
	return winnings
}
