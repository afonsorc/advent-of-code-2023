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
	reading := make([]int, 0)
	for i := 0; i < len(input); i++ {
		line := strings.Split(input[i], " ")
		for i := 0; i < len(line); i++ {
			test, _ := strconv.Atoi(line[i])
			reading = append(reading, test)
		}
		if isPartTwo {
			reading = extrapolateBackwards(reading, 0, len(reading))
			solution += reading[0]
		} else {
			reading = extrapolateForwards(reading, len(reading)-1)
			solution += reading[len(reading)-1]
		}
		reading = nil
	}
	return solution
}

func extrapolateForwards(line []int, length int) []int {
	isFinal := true
	for i := 0; i < length; i++ {
		line[i] = line[i+1] - line[i]
		if line[i] != 0 {
			isFinal = false
		}
	}
	if isFinal {
		line[length] = line[length] + line[length-1]
		return line
	} else {
		line = extrapolateForwards(line, length-1)
		line[length] = line[length] + line[length-1]
	}
	return line
}

func extrapolateBackwards(line []int, length int, oglen int) []int {
	isFinal := true
	for i := oglen - 1; i > length; i-- {
		line[i] = line[i] - line[i-1]
		if line[i] != 0 {
			isFinal = false
		}
	}
	if isFinal {
		line[length] = line[length] - line[length+1]
		return line
	} else {
		line = extrapolateBackwards(line, length+1, oglen)
		line[length] = line[length] - line[length+1]
	}
	return line
}
