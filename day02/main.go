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

	solution := 0
	cubes := map[string]int{"red": 12, "green": 13, "blue": 14}

	for i := 0; i < len(input); i++ {
		input[i] = strings.Replace(input[i], ":", "", -1)
		input[i] = strings.Replace(input[i], ";", "", -1)
		input[i] = strings.Replace(input[i], ",", "", -1)
	}

	for _, line := range input {

		game := strings.Split(line, " ")
		validGame := true
		if isPartTwo {
			clear(cubes)
		}

		for i := 2; i < len(game); i += 2 {
			if cubeAmount, _ := strconv.Atoi(game[i]); cubeAmount > cubes[game[i+1]] {
				validGame = false
				if isPartTwo {
					cubes[game[i+1]] = cubeAmount
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
