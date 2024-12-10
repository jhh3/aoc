package main

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/jhh3/aoc/common"
)

//go:embed input.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	common.RunFromSolver(&solver{}, input)
}

//--------------------------------------------------------------------
// Solution
//--------------------------------------------------------------------

type solver struct{}

func (s *solver) SolvePart1(input string) string {
	problemInput := parseInput(input)
	safeCount := 0

	for _, report := range problemInput.ReactReports {
		if isReportSafe(report) {
			safeCount++
		}
	}

	return strconv.Itoa(safeCount)
}

func isReportSafe(report []int) bool {
	// all increasing or decreasing
	// adjacent numbers differ by 1-3 at least/most

	lastNumber := report[0]
	isIncreasing := report[1] > lastNumber
	for _, number := range report[1:] {
		diff := common.AbsInt(number - lastNumber)
		if diff < 1 || diff > 3 {
			return false
		}

		numberIncreased := number > lastNumber
		if numberIncreased != isIncreasing {
			return false
		}

		lastNumber = number
	}

	return true
}

func (s *solver) SolvePart2(input string) string {
	// TODO
	return ""
}

type ProblemInput struct {
	ReactReports [][]int
}

func parseInput(input string) ProblemInput {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	result := ProblemInput{
		ReactReports: make([][]int, len(lines)),
	}
	for i, row := range lines {
		parts := strings.Fields(row)
		result.ReactReports[i] = make([]int, len(parts))
		for j, part := range parts {
			intPart, err := strconv.Atoi(part)
			if err != nil {
				panic("Invalid integer input")
			}
			result.ReactReports[i][j] = intPart
		}
	}

	return result
}
