package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	solution2 := solve(input, true)

	fmt.Print("Part 1: ")
	fmt.Println(solution1)

	fmt.Print("Part 2: ")
	fmt.Println(solution2)
}

func solve(input []string, isPartTwo bool) any {

	solution := 0
	cubes := map[string]int{"red": 12, "green": 13, "blue": 14}

	for i := 0; i < len(input); i++ {
		input[i] = strings.Replace(input[i], ":", "", -1)
		input[i] = strings.Replace(input[i], ";", "", -1)
		input[i] = strings.Replace(input[i], ",", "", -1)
	}

	for _, line := range input {

		validGame := true
		if isPartTwo {
			clear(cubes)
		}
		game := strings.Split(line, " ")
		for i := 2; i < len(game); i += 2 {

			cubeAmount, _ := strconv.Atoi(game[i])
			if cubeAmount > cubes[game[i+1]] {
				if isPartTwo {
					cubes[game[i+1]] = cubeAmount
				} else {
					validGame = false
					break
				}
			}
		}

		if isPartTwo {
			solution += cubes["red"] * cubes["green"] * cubes["blue"]
		} else if validGame {
			id, _ := strconv.Atoi(game[1])
			solution += id
		}
	}
	return solution
}
