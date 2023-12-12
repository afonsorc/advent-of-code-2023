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

	solution1 := solve(input, 2)
	solution2 := solve(input, 1000000)

	fmt.Print("Part 1: ")
	fmt.Println(solution1)

	fmt.Print("Part 2: ")
	fmt.Println(solution2)
}

func solve(input []string, distance int) any {

	type Pos struct {
		x int
		y int
	}

	solution := 0
	galaxies := make([]Pos, 0)
	space := make([][]string, len(input))
	for i := range space {
		space[i] = make([]string, len(input[i]))
	}

	emptyRows := make([]bool, len(space))
	emptyCols := make([]bool, len(space[0]))
	for i := range space {
		emptyRows[i] = true
		emptyCols[i] = true
	}

	for i, line := range input {
		for j, star := range line {
			space[i][j] = string(star)
			if string(star) == "#" {
				galaxies = append(galaxies, Pos{j, i})
				emptyRows[i] = false
				emptyCols[j] = false
			}
		}
	}

	for i, src := range galaxies {
		for _, dest := range galaxies[i+1:] {

			x := abs(src.x, dest.x)
			y := abs(src.y, dest.y)

			for row := range emptyRows {
				if emptyRows[row] && (row < src.y && row > dest.y || row < dest.y && row > src.y) {
					y += distance - 1
				}
			}
			for col := range emptyCols {
				if emptyCols[col] && (col < src.x && col > dest.x || col < dest.x && col > src.x) {
					x += distance - 1
				}
			}
			solution += x + y
		}
	}

	return solution
}

func abs(a, b int) int {
	if a >= b {
		return a - b
	} else {
		return b - a
	}
}
