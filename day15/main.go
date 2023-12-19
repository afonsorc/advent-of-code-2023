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

	solution1 := solve2(input, false)
	solution2 := solve2(input, true)

	fmt.Print("Part 1: ")
	fmt.Println(solution1)

	fmt.Print("Part 2: ")
	fmt.Println(solution2)
}

func solve2(input []string, isPartTwo bool) any {
	solution := 0

	type Lens struct {
		box        int
		focusPower int
		order      int
	}

	lenses := map[string]Lens{}
	boxes := [256][]string{}
	for i := range boxes {
		boxes[i] = make([]string, 0)
	}

	sequence := strings.Split(input[0], ",")
	for i := range sequence {
		hash := 0
		labelSize := 0
		for labelSize = 0; labelSize < len(sequence[i]); labelSize++ {
			if isPartTwo && (sequence[i][labelSize] == '=' || sequence[i][labelSize] == '-') {
				break
			}
			hash += int(sequence[i][labelSize])
			hash *= 17
			hash %= 256
		}

		if isPartTwo {
			if sequence[i][labelSize] == '=' {
				focusPower, _ := strconv.Atoi(string(sequence[i][labelSize+1]))
				lens, ok := lenses[sequence[i][:labelSize]]
				if !ok {
					lenses[sequence[i][:labelSize]] = Lens{hash, focusPower, len(boxes[hash])}
					boxes[hash] = append(boxes[hash], sequence[i][:labelSize])
				} else {
					lenses[sequence[i][:labelSize]] = Lens{lens.box, focusPower, lens.order}
				}
			} else if sequence[i][labelSize] == '-' {
				_, ok := lenses[sequence[i][:labelSize]]
				if ok {
					lens := 0
					for lens = range boxes[hash] {
						if boxes[hash][lens] == sequence[i][:labelSize] {
							break
						}
					}
					if lens == len(boxes[hash]) {
						boxes[hash] = boxes[hash][:len(boxes[hash])-1]
					} else {
						boxes[hash] = append(boxes[hash][:lens], boxes[hash][lens+1:]...)
					}
					delete(lenses, sequence[i][:labelSize])
				}
			}
		} else {
			solution += hash
		}
	}

	if isPartTwo {
		for _, box := range boxes {
			for j, lens := range box {
				solution += (lenses[lens].box + 1) * (j + 1) * lenses[lens].focusPower
			}
		}
	}
	return solution
}
