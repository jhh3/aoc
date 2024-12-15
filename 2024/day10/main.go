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
	return strconv.Itoa(problemInput.ScoreTopology(true))
}

func (s *solver) SolvePart2(input string) string {
	problemInput := parseInput(input)
	return strconv.Itoa(problemInput.ScoreTopology(false))
}

type Point struct {
	Row, Col int
}

type ProblemInput struct {
	Topology [][]int

	Trailheads []Point
}

func (pi *ProblemInput) ScoreTopology(part1 bool) int {
	score := 0

	for _, th := range pi.Trailheads {
		if part1 {
			score += pi.ScoreTrailheadPart1(th)
		} else {
			score += pi.ScoreTrailheadPart2(th)
		}
	}

	return score
}

// ScoreTrailhead returns the number of 9-height positions reachable from the trailhead, gradually increasing by exactly 1 height at each step.
func (pi *ProblemInput) ScoreTrailheadPart2(th Point) int {
	score := 0

	height := pi.Topology[th.Row][th.Col]

	// base case
	if height == 9 {
		return 1
	}

	// recursive case
	for _, nextPoint := range pi.GetPossibleMoves(th, height) {
		score += pi.ScoreTrailheadPart2(nextPoint)
	}

	return score
}

// ScoreTrailhead returns the number of 9-height positions reachable from the trailhead, gradually increasing by exactly 1 height at each step.
func (pi *ProblemInput) ScoreTrailheadPart1(th Point) int {
	score := 0
	peaks := pi.GetReachablePeaks(th)
	seen := map[Point]bool{}
	for _, peak := range peaks {
		if ok := seen[peak]; !ok {
			seen[peak] = true
			score++
		}
	}

	return score
}

func (pi *ProblemInput) GetReachablePeaks(p Point) []Point {
	peaks := []Point{}

	height := pi.Topology[p.Row][p.Col]

	// base case
	if height == 9 {
		return []Point{p}
	}

	// recursive case
	for _, nextPoint := range pi.GetPossibleMoves(p, height) {
		peaks = append(peaks, pi.GetReachablePeaks(nextPoint)...)
	}

	return peaks
}

func (pi *ProblemInput) GetPossibleMoves(th Point, h int) []Point {
	possibleMoves := []Point{}
	height := pi.Topology[th.Row][th.Col]
	// attempt to move up, down, left, right
	for _, move := range []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		possibleNextPoint := Point{Row: th.Row + move.Row, Col: th.Col + move.Col}
		if !pi.IsInBounds(possibleNextPoint) {
			continue
		}

		nextHeight := pi.Topology[possibleNextPoint.Row][possibleNextPoint.Col]
		if nextHeight == height+1 {
			// we can move here
			possibleMoves = append(possibleMoves, possibleNextPoint)
		}
	}

	return possibleMoves
}

func (pi *ProblemInput) IsInBounds(p Point) bool {
	return p.Row >= 0 && p.Row < len(pi.Topology) && p.Col >= 0 && p.Col < len(pi.Topology[0])
}

func parseInput(input string) ProblemInput {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	pi := ProblemInput{
		Topology: make([][]int, len(lines)),
	}

	for i, row := range lines {
		pi.Topology[i] = make([]int, len(row))
		for j, element := range row {
			pi.Topology[i][j] = int(element - '0')
			if element == '0' {
				pi.Trailheads = append(pi.Trailheads, Point{Row: i, Col: j})
			}
		}
	}

	return pi
}
