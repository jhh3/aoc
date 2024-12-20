package main

import (
	_ "embed"
	"image"
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
	return common.Itoa(problemInput.CountCheats(2, 100))
}

func (s *solver) SolvePart2(input string) string {
	problemInput := parseInput(input)
	return common.Itoa(problemInput.CountCheats(20, 100))
}

type ProblemInput struct {
	Grid  map[image.Point]rune
	Start image.Point
	End   image.Point
}

func parseInput(input string) *ProblemInput {
	grid, start, end := map[image.Point]rune{}, image.Point{}, image.Point{}
	for y, s := range strings.Fields(input) {
		for x, r := range s {
			if r == 'S' {
				start = image.Point{x, y}
			}
			if r == 'E' {
				end = image.Point{x, y}
			}

			grid[image.Point{x, y}] = r
		}
	}
	return &ProblemInput{
		Grid:  grid,
		Start: start,
		End:   end,
	}
}

func (pi *ProblemInput) CountCheats(cheatDistance, goalTimeSave int) int {
	// Get distance from start to all points
	queue, dist := []image.Point{pi.Start}, map[image.Point]int{pi.Start: 0}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		for _, d := range []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
			n := p.Add(d)
			if _, ok := dist[n]; !ok && pi.Grid[n] != '#' {
				queue, dist[n] = append(queue, n), dist[p]+1
			}
		}
	}

	// Find time saves with manhattan distance
	// if we were to cheat
	count := 0
	for p1 := range dist {
		for p2 := range dist {
			d := common.AbsInt(p2.X-p1.X) + common.AbsInt(p2.Y-p1.Y)
			if d <= cheatDistance && dist[p2] >= dist[p1]+d+goalTimeSave {
				count++
			}
		}
	}

	return count
}
