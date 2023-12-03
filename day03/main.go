package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func solve(input []string, isPartTwo bool) int {

	type Gear struct {
		ratio int
		parts int
	}

	gears := make([]Gear, len(input)*len(input[0]))

	for i := 0; i < len(gears); i++ {
		gears[i].ratio = 1
		gears[i].parts = 0
	}

	numberIndex := -1
	schematic := 0
	for row := 0; row < len(input); row++ {
		for col := 0; col < len(input[row]); col++ {
			if unicode.IsDigit(rune(input[row][col])) {
				if numberIndex == -1 {
					numberIndex = col
				}
			}

			// either digit and last column or non digit and expecting one
			if (unicode.IsDigit(rune(input[row][col])) && col+1 == len(input[row])) || (!unicode.IsDigit(rune(input[row][col])) && numberIndex != -1) {

				// correct when number is not in last column
				if !unicode.IsDigit(rune(input[row][col])) {
					col--
				}

				// search for one symbol in area
				gearIndex := 0
				isPartNumber := false
				number, _ := strconv.Atoi(input[row][numberIndex:min(col+1, len(input[row]))])
				for height := max(0, row-1); height < min(len(input), row+2); height++ {
					for width := max(0, numberIndex-1); width < min(len(input[row]), col+2); width++ {

						// skip if looking at the actual number
						if height == row && width >= numberIndex && width <= col {
							continue
						}

						// search for any non-dot symbols
						if !isPartTwo && (input[height][width] < 48 || input[height][width] > 57) && input[height][width] != 46 {
							isPartNumber = true
							break
						}
						if isPartTwo && input[height][width] == 42 {
							gearIndex = height*len(input[row]) + width
							isPartNumber = true
						}
					}
				}

				if isPartNumber {
					if isPartTwo {
						gears[gearIndex].ratio *= number
						gears[gearIndex].parts++
					} else {
						schematic += number
					}
				}
				numberIndex = -1
			}
		}
	}
	if isPartTwo {
		for _, gear := range gears {
			if gear.parts > 1 {
				schematic += gear.ratio
			}
		}
	}
	return schematic
}
