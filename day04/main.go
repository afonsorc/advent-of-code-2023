package main

import (
	"bufio"
	"fmt"
	"os"
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
	solution2 := solve(input, true)

	fmt.Print("Part 1: ")
	fmt.Println(solution1)

	fmt.Print("Part 2: ")
	fmt.Println(solution2)
}

func solve(input []string, isPartTwo bool) any {
	points := 0
	copies := [200]int{}

	// set the original card for each
	for copy := 0; copy < len(copies); copy++ {
		copies[copy] = 1
	}

	for i, line := range input {
		splitIndex := -1
		cardPoints := 0
		winningNumbersOnCard := 0
		winning := [100]bool{false}

		line = strings.Join(strings.Fields(line), " ")
		card := strings.Split(line, " ")
		for index, number := range card {
			if number == "|" {
				splitIndex = index
			}
		}

		for index := 2; index < splitIndex; index++ {
			number, _ := strconv.Atoi(card[index])
			winning[number] = true
		}

		// get winning numbers from the current card
		for index := splitIndex + 1; index < len(card); index++ {
			number, _ := strconv.Atoi(card[index])
			if winning[number] {
				winningNumbersOnCard++
				if cardPoints == 0 {
					cardPoints = 1
				} else {
					cardPoints *= 2
				}
			}
		}

		// add the copies of the following cards
		for copy := i; copy < i+winningNumbersOnCard; copy++ {
			copies[copy+1] += copies[i]
		}

		// sum points from the current card
		if isPartTwo {
			points += copies[i]
		} else {
			points += cardPoints
		}
	}
	return points
}
