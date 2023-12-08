package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	left  string
	right string
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

	steps := 0
	nodes := map[string]Node{}
	cycles := map[string]int{}
	currentNodes := make([]string, 0)
	formatter := strings.NewReplacer("= ", "", ",", "", "(", "", ")", "")

	for i := 2; i < len(input); i++ {
		input[i] = formatter.Replace(input[i])
		instruction := strings.Split(input[i], " ")
		nodes[instruction[0]] = Node{instruction[1], instruction[2]}
		if instruction[0][2] == 'A' && isPartTwo {
			currentNodes = append(currentNodes, instruction[0])
		}
	}

	if isPartTwo {
		for _, node := range currentNodes {
			for steps = 0; ; steps++ {
				if node[2] == 'Z' && input[0][steps%len(input[0])] == input[0][0] && steps != 0 {
					break
				}
				node = moveToChild(input[0][steps%len(input[0])], nodes, node)
			}
			cycles[node] = steps
		}
	} else {
		node := "AAA"
		for steps = 0; ; steps++ {
			if node == "ZZZ" {
				return steps
			}
			node = moveToChild(input[0][steps%len(input[0])], nodes, node)
		}
	}
	steps = 1
	for _, cycle := range cycles {
		steps = leastCommonMultiple(steps, cycle)
	}
	return steps
}

func moveToChild(dir byte, nodes map[string]Node, current string) string {
	if dir == 'L' {
		return nodes[current].left
	} else {
		return nodes[current].right
	}
}

func greatestCommonDivisor(a, b int) int {
	for b != 0 {
		c := b
		b = a % b
		a = c
	}
	return a
}

func leastCommonMultiple(a int, b int) int {
	return a * b / greatestCommonDivisor(a, b)
}
