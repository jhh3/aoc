package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/jhh3/aoc/common"
)

const (
	DAY  = 8
	YEAR = 2023
)

func main() {
	common.RunFromSolver(&solver{}, YEAR, DAY)
}

//--------------------------------------------------------------------
// Solution
//--------------------------------------------------------------------

type solver struct{}

func (ps *solver) SolvePart1(input string) string {
	data := parseInput(input)

	current := "AAA"
	numSteps := 0

	for {
		for _, instruction := range data.instructions {
			numSteps++

			next := data.nodeToNeighbors[current].Go(string(instruction))
			current = next

			if current == "ZZZ" {
				goto Escape
			}
		}
	}
Escape:

	return strconv.Itoa(numSteps)
}

func (ps *solver) SolvePart2(input string) string {
	data := parseInput(input)

	answers := make([]int, 0)
	for node := range data.nodeToNeighbors {
		if node[2] == 'A' {
			numSteps := solve(node, data)
			answers = append(answers, numSteps)
		}
	}

	answer := common.LCM(answers[0], answers[1], answers[2:]...)
	return strconv.Itoa(answer)
}

//--------------------------------------------------------------------
// Helpers
//--------------------------------------------------------------------

func solve(start string, data Input) int {
	current := start
	numSteps := 0

	for {
		for _, instruction := range data.instructions {
			numSteps++

			next := data.nodeToNeighbors[current].Go(string(instruction))
			current = next

			if current[2] == 'Z' {
				return numSteps
			}
		}
	}
}

func AllHaveSuffix(nodes []string, suffix string) bool {
	for _, node := range nodes {
		if node[2] != 'Z' {
			return false
		}
	}
	return true
}

//--------------------------------------------------------------------
// Parsing code
//--------------------------------------------------------------------

type Node struct {
	value string

	left  string // *Node
	right string // *Node
}

func (n Node) Go(instruction string) string {
	if instruction == "L" {
		return n.left
	}
	return n.right
}

type Input struct {
	instructions string

	nodeToNeighbors map[string]Node
}

func parseInput(input string) Input {
	result := Input{
		nodeToNeighbors: make(map[string]Node),
	}

	lines := strings.Split(string(input), "\n")
	for idx, line := range lines {
		cleanLine := strings.TrimSpace(line)
		if cleanLine == "" {
			continue
		}

		if idx == 0 {
			result.instructions = cleanLine
			continue
		}

		node := parseNode(cleanLine)
		result.nodeToNeighbors[node.value] = node
	}

	return result
}

var nodeRegex = regexp.MustCompile(`([A-Z]+) = \(([A-Z]+), ([A-Z]+)\)`)

func parseNode(input string) Node {
	// Regex to parse node, e.g.
	// BKM = (CDC, PSH)
	matches := nodeRegex.FindStringSubmatch(input)
	return Node{
		value: matches[1],
		left:  matches[2],
		right: matches[3],
	}
}
