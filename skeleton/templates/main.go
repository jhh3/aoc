package main

import (
	_ "embed"
	"strings"

	"github.com/jhh3/aoc/common"
)

const (
	DAY  = 3
	YEAR = 2023
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
	common.RunFromSolver(&solver{}, YEAR, DAY)
}

//--------------------------------------------------------------------
// Solution
//--------------------------------------------------------------------

type solver struct{}

func (s *solver) SolvePart1(input string) string {
	panic("not implemented")
}

func (s *solver) SolvePart2(input string) string {
	panic("not implemented")
}
