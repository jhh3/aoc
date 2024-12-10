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
	safeCount := problemInput.countSafe(false)
	return strconv.Itoa(safeCount)
}

func (s *solver) SolvePart2(input string) string {
	problemInput := parseInput(input)
	safeCount := problemInput.countSafe(true)
	return strconv.Itoa(safeCount)
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

func (pi *ProblemInput) countSafe(problemDamperOn bool) int {
	safeCount := 0

	for _, report := range pi.ReactReports {
		if isReportSafe(report, problemDamperOn) {
			safeCount++
		}
	}

	return safeCount
}

func isReportSafe(report []int, problemDamperOn bool) bool {
	// all increasing or decreasing
	// adjacent numbers differ by 1-3 at least/most
	// if problemDamperOn, then we can skip 1 level

	lastLevel := report[0]
	isIncreasing := report[1] > lastLevel
	for _, level := range report[1:] {
		safeTransition := true
		diff := common.AbsInt(level - lastLevel)
		if diff < 1 || diff > 3 {
			safeTransition = false
		}

		levelIncreased := level > lastLevel
		if levelIncreased != isIncreasing {
			safeTransition = false
		}

		if !safeTransition {
			if problemDamperOn {
				for i := range report {
					reportCopy := make([]int, len(report))
					copy(reportCopy, report[:])
					newReport := common.RemoveIndex(reportCopy, i)

					if isReportSafe(newReport, false) {
						return true
					}
				}
				return false
			} else {
				return false
			}
		} else {
			lastLevel = level
		}
	}

	return true
}
