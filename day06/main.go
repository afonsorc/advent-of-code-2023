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

	type Race struct {
		time int
		dist int
	}

	margin := 1
	races := make([]Race, 0)

	times := strings.Fields(input[0])
	dists := strings.Fields(input[1])

	for i := 1; i < len(times); i++ {
		if isPartTwo {
			times[i] = strings.Join(times[1:], "")
			dists[i] = strings.Join(dists[1:], "")
		}

		time, _ := strconv.Atoi(times[i])
		dist, _ := strconv.Atoi(dists[i])
		races = append(races, Race{time, dist})
		if isPartTwo {
			break
		}
	}

	for _, race := range races {
		wins := 0
		for time := 0; time < race.time; time++ {
			if (race.time-time)*time > race.dist {
				wins++
			}
		}
		margin *= wins
	}
	return margin
}
