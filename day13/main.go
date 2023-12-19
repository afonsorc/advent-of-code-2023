package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, _ := os.Open("input.txt")
	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	solution1 := solve(input, 0)
	solution2 := solve(input, 1)

	fmt.Print("Part 1: ")
	fmt.Println(solution1)

	fmt.Print("Part 2: ")
	fmt.Println(solution2)
}

func solve(input []string, maxSmudges int) any {

	solution := 0
	lava := make([]string, 0)

	for i, line := range input {
		if line == "" || i == len(input)-1 {
			if i == len(input)-1 {
				lava = append(lava, line)
			}
			row := getMirrorRow(lava, maxSmudges)
			col := getMirrorCol(lava, maxSmudges)
			solution += row*100 + col
			lava = nil
			continue
		}
		lava = append(lava, line)
	}
	return solution
}

func getMirrorRow(lava []string, maxSmudges int) int {
	for row := 0; row < len(lava)-1; row++ {
		smudges := 0
		for rowUp, rowDown := row, row+1; rowUp >= 0 && rowDown < len(lava); rowUp, rowDown = rowUp-1, rowDown+1 {
			for col := 0; col < len(lava[0]); col++ {
				if lava[rowUp][col] != lava[rowDown][col] {
					smudges++
				}
			}

		}
		if smudges == maxSmudges {
			return row + 1
		}
	}
	return 0
}

func getMirrorCol(lava []string, maxSmudges int) int {
	for col := 0; col < len(lava[0])-1; col++ {
		smudges := 0
		for colLeft, colRight := col, col+1; colLeft >= 0 && colRight < len(lava[0]); colLeft, colRight = colLeft-1, colRight+1 {
			for row := 0; row < len(lava); row++ {
				if lava[row][colLeft] != lava[row][colRight] {
					smudges++
				}
			}
		}
		if smudges == maxSmudges {
			return col + 1
		}
	}
	return 0
}
