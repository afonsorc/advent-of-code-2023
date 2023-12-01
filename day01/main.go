package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func readInputFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {

	input, err := readInputFile("input.txt")
	if err != nil {
		fmt.Println("Error: Failed to read input file")
		os.Exit(1)
	}

	solution1 := solve(input, false)
	fmt.Print("Part 1: ")
	fmt.Println(solution1)

	solution2 := solve(input, true)
	fmt.Print("Part 2: ")
	fmt.Println(solution2)
}

func solve(input []string, isPartTwo bool) any {

	digits := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	solution := 0
	first_digit := 0
	last_digit := 0

	for _, line := range input {

		first_index := len(line) - 1
		last_index := 0

		for i := 0; i < len(line); i++ {
			if unicode.IsDigit(rune(line[i])) {
				first_digit = int(line[i] - 48)
				first_index = i
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(line[i])) {
				last_digit = int(line[i] - 48)
				last_index = i
				break
			}
		}

		if isPartTwo {
			for digit, spelled_digit := range digits {
				if strings.Contains(line, spelled_digit) {

					// check if there is a spelled out digit at the start
					index_from_start := strings.Index(line, spelled_digit)
					if index_from_start < first_index {
						first_index = index_from_start
						first_digit = digit + 1
					}

					// check if there is a spelled out digit at the end
					index_from_end := strings.LastIndex(line, spelled_digit)
					if index_from_end > last_index {
						last_index = index_from_end
						last_digit = digit + 1
					}
				}
			}
		}

		solution += first_digit*10 + last_digit
	}
	return solution
}
