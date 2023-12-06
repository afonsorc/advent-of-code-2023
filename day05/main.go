package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
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

	type Seed struct {
		start    int
		end      int
		isMapped bool
	}

	// get seeds
	input[0] = strings.Replace(input[0], "seeds: ", "", -1)
	line := strings.Split(input[0], " ")
	seeds := make([]Seed, 0)

	seedSkip := 1
	if isPartTwo {
		seedSkip = 2
	}

	for i := 0; i < len(line); i += seedSkip {
		seedStart, _ := strconv.Atoi(line[i])
		seedEnd := seedStart
		if isPartTwo {
			seedRange, _ := strconv.Atoi(line[i+1])
			seedEnd += seedRange - 1
		}
		seeds = append(seeds, Seed{seedStart, seedEnd, false})
	}

	for i := 1; i < len(input); i++ {
		// skip non-input lines
		if input[i] == "" || !unicode.IsDigit(rune(input[i][0])) {
			for j := 0; j < len(seeds); j++ {
				seeds[j].isMapped = false
			}
			continue
		}

		almanac := strings.Split(input[i], " ")
		mappedSeed, _ := strconv.Atoi(almanac[0])
		filterStart, _ := strconv.Atoi(almanac[1])
		filterRange, _ := strconv.Atoi(almanac[2])
		filterEnd := filterStart + filterRange - 1
		mapping := mappedSeed - filterStart

		for i := 0; i < len(seeds); i++ {
			if seeds[i].isMapped {
				continue
			}

			// get intersection range between seed and filter
			start := max(seeds[i].start, filterStart)
			end := min(seeds[i].end, filterEnd)
			if seeds[i].end < filterStart || filterEnd < seeds[i].start {
				continue
			}

			// new range at the start without mapping
			if start > seeds[i].start {
				seeds = append(seeds, Seed{seeds[i].start, start - 1, false})
				seeds[i].start = start
			}

			// new range at the end without mapping
			if end < seeds[i].end {
				seeds = append(seeds, Seed{end + 1, seeds[i].end, false})
				seeds[i].end = end
			}

			// map the intersected range
			seeds[i].start += mapping
			seeds[i].end += mapping
			seeds[i].isMapped = true
		}
	}

	closestLocation := seeds[0].start
	for _, seed := range seeds {
		closestLocation = min(closestLocation, seed.start)
	}
	return closestLocation
}
