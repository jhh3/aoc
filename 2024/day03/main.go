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

	mulRe := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	// do() regex
	doRe := regexp.MustCompile(`do\(\)`)
	// don't() regex
	dontRe := regexp.MustCompile(`don't\(\)`)

	doMatches := doRe.FindAllStringIndex(input, -1)
	dontMatches := dontRe.FindAllStringIndex(input, -1)

	sum := 0
	do := true
	doIndex := 0
	dontIndex := 0
	start := 0

	for doIndex < len(doMatches) && dontIndex < len(dontMatches) {
		if do {
			// find the next dont (past start)
			dontInputIndex := dontMatches[dontIndex][1]
			if dontInputIndex < start {
				dontIndex++
				continue
			} else {
				inputSubstr := input[start:dontInputIndex]
				mulMatches := mulRe.FindAllStringSubmatch(inputSubstr, -1)
				for _, match := range mulMatches {
					// multiply
					x, _ := strconv.Atoi(match[1])
					y, _ := strconv.Atoi(match[2])
					sum += x * y
				}
				start = dontInputIndex
				do = false
			}
		} else {
			// find the next do (past start)
			doInputIndex := doMatches[doIndex][1]
			if doInputIndex < start {
				doIndex++
				continue
			} else {
				start = doInputIndex
				do = true
			}
		}
	}

	if do {
		// do to end
		inputSubstr := input[start:]
		mulMatches := mulRe.FindAllStringSubmatch(inputSubstr, -1)
		for _, match := range mulMatches {
			// multiply
			x, _ := strconv.Atoi(match[1])
			y, _ := strconv.Atoi(match[2])
			sum += x * y
		}
	}

	return strconv.Itoa(sum)
}
