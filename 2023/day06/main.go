package main

import (
	"fmt"
	"regexp"
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
	result := int64(1)
	races := ps.parseInput(input, false)
	for _, race := range races {
		result *= race.CalculateNumberOfWaysToWin()
	}

	return fmt.Sprintf("%d", result)
}

func (ps *solver) SolvePart2(input string) string {
	race := ps.parseInput(input, true)[0]
	result := race.CalculateNumberOfWaysToWin()

	return fmt.Sprintf("%d", result)
}

//--------------------------------------------------------------------
// Helpers
//--------------------------------------------------------------------

type RaceTiming struct {
	time     int64
	distance int64
}

func (rt RaceTiming) CalculateNumberOfWaysToWin() int64 {
	result := int64(0)
	for speed := int64(1); speed <= rt.time; speed++ {
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

func (ps *solver) parseInput(input string, noSpaces bool) []RaceTiming {
	lines := strings.Split(string(input), "\n")

	rep := " "
	if noSpaces {
		rep = ""
	}

	timeLine := strings.TrimSpace(strings.Split(lines[0], ":")[1])
	durationsLine := strings.TrimSpace(strings.Split(lines[1], ":")[1])

	times := strings.Split(space.ReplaceAllString(timeLine, rep), " ")
	durations := strings.Split(space.ReplaceAllString(durationsLine, rep), " ")

	var result = []RaceTiming{}
	for i, time := range times {
		result = append(result, RaceTiming{
			time:     int64(common.MustAtoi(time)),
			distance: int64(common.MustAtoi(durations[i])),
		})
	}

	return result
}
