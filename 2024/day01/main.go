package main

import (
	_ "embed"
	"sort"
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

	// sort the lists
	sortedLeftList := make([]int, len(problemInput.LeftList))
	sortedRightList := make([]int, len(problemInput.RightList))
	copy(sortedLeftList, problemInput.LeftList)
	copy(sortedRightList, problemInput.RightList)
	sort.Ints(sortedLeftList)
	sort.Ints(sortedRightList)

	// calculate the difference
	diff := 0
	for i, left := range sortedLeftList {
		right := sortedRightList[i]
		diff += common.AbsInt(left - right)
	}

	return strconv.Itoa(diff)
}

func (s *solver) SolvePart2(input string) string {
	// panic("not implemented")
	occurenceCountMap := make(map[int]int)
	for _, num := range parseInput(input).RightList {
		if _, ok := occurenceCountMap[num]; !ok {
			occurenceCountMap[num] = 0
		}
		occurenceCountMap[num]++
	}

	similarityScore := 0
	for _, num := range parseInput(input).LeftList {
		// in go, this returns the zero value, 0, if it doesn't exist, which is what we want
		count := occurenceCountMap[num]
		similarityScore += count * num
	}

	return strconv.Itoa(similarityScore)
}

type ProblemInput struct {
	LeftList  []int
	RightList []int
}

func parseInput(input string) *ProblemInput {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	leftList := []int{}
	rightList := []int{}

	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) != 2 {
			panic("invalid input format")
		}
		left, err1 := strconv.Atoi(parts[0])
		right, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			panic("invalid integer in input")
		}
		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	return &ProblemInput{
		LeftList:  leftList,
		RightList: rightList,
	}
}
