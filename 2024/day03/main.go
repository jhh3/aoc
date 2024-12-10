package main

import (
	_ "embed"
	"regexp"
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
	sum := 0

	// each line contains many mul instructions
	// mul(X,Y) where X and Y are 1-3 digit numbers
	// extract the mul instructions and sum them
	// mul\(\d{1,3},\d{1,3}\)
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	// find all matches
	matches := re.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		// multiply
		x, err := strconv.Atoi(match[1])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(match[2])
		if err != nil {
			panic(err)
		}

		sum += x * y
	}

	return strconv.Itoa(sum)
}

func (s *solver) SolvePart2(input string) string {
	// todo
	return ""
}
