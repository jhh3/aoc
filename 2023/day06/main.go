package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/jhh3/aoc/common"
)

const (
	DAY  = 6
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
	result := 1
	races := ps.parseInput(input)
	for _, race := range races {
		result *= race.CalculateNumberOfWaysToWin()
	}

	return strconv.Itoa(result)
}

func (ps *solver) SolvePart2(input string) string {
	return ""
}

//--------------------------------------------------------------------
// Helpers
//--------------------------------------------------------------------

type RaceTiming struct {
	time     int
	distance int
}

func (rt RaceTiming) CalculateNumberOfWaysToWin() int {
	result := 0
	for speed := 1; speed <= rt.time; speed++ {
		timeToTravel := rt.time - speed
		distance := speed * timeToTravel
		if distance > rt.distance {
			result++
		}
	}
	return result
}

func (rt RaceTiming) String() string {
	return fmt.Sprintf("t: %d d: %d", rt.time, rt.distance)
}

//--------------------------------------------------------------------
// Parsing code
//--------------------------------------------------------------------

var space = regexp.MustCompile(`\s+`)

func (ps *solver) parseInput(input string) []RaceTiming {
	lines := strings.Split(string(input), "\n")

	times := strings.Split(space.ReplaceAllString((lines[0]), " "), " ")
	durations := strings.Split(space.ReplaceAllString((lines[1]), " "), " ")

	var result = []RaceTiming{}
	for i, time := range times {
		if i == 0 {
			continue
		}

		result = append(result, RaceTiming{
			time:     common.MustAtoi(time),
			distance: common.MustAtoi(durations[i]),
		})
	}

	return result
}
