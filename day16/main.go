package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Pos struct {
	x int
	y int
}

type Beam struct {
	pos Pos
	dir Pos
}

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
	tiles := make([]string, len(input))
	vector := [4]Pos{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	height := 1
	width := 1

	if isPartTwo {
		height = len(input)
		width = len(input[0])
	}

	// bfs
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {

			if (x != 0 && x != len(input[0])-1) && (y != 0 && y != len(input[0])-1) {
				continue
			}

			invalid := []bool{x != 0, y != 0, x != len(input[0])-1, y != len(input)-1}
			for d, startDir := range vector {

				if invalid[d] {
					continue
				}

				for i := range input {
					tiles[i] = strings.Clone(input[i])
				}

				energized := 0
				queue := make([]Beam, 0)
				queue = append(queue, Beam{Pos{x, y}, startDir})

				visited := map[int]Beam{}
				visited[0] = Beam{Pos{x, y}, startDir}

				for ; len(queue) > 0; energized++ {
					current := queue[0]
					queue = queue[1:]

					b := []byte(tiles[current.pos.y])
					if b[current.pos.x] == '#' {
						energized--
					}
					b[current.pos.x] = '#'
					tiles[current.pos.y] = string(b)

					for _, dir := range vector {
						next := Beam{Pos{current.pos.x + dir.x, current.pos.y + dir.y}, dir}
						if next.pos.x < 0 || next.pos.x >= len(input[0]) || next.pos.y < 0 || next.pos.y >= len(input) {
							continue
						}

						value, ok := visited[next.pos.y*len(input[0])+next.pos.x]
						if ok && next == value {
							continue
						}

						tile := input[current.pos.y][current.pos.x]
						if isValidPath(tile, current.dir, dir) || isValidSplitter(tile, current.dir, dir) || isValidMirror(tile, current.dir, dir) {
							queue = append(queue, next)
							visited[next.pos.y*len(input[0])+next.pos.x] = next
						}
					}
				}
				solution = max(solution, energized)
			}
		}
	}
	return solution
}

func isValidPath(tile byte, one Pos, two Pos) bool {
	return tile == '.' && one == two
}

func isValidSplitter(tile byte, one Pos, two Pos) bool {
	return (tile == '|' && one == two && one.x == 0) || (tile == '|' && (one.x == two.y || one.x == -two.y) && one.x != 0) || (tile == '-' && one == two && one.y == 0) || (tile == '-' && (one.x == two.y || one.x == -two.y) && one.y != 0)
}

func isValidMirror(tile byte, one Pos, two Pos) bool {
	return (tile == '/' && one.x == -two.y && one.y == -two.x) || (tile == '\\' && one.x == two.y && one.y == two.x)
}
