package main

import (
	_ "embed"
	"fmt"
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
	problemInput.PrettyPrint()
	problemInput.ApplyMoveSequence()
	println()
	println()
	problemInput.PrettyPrint()
	return strconv.Itoa(problemInput.SumBoxGPSValues())
}

func (s *solver) SolvePart2(input string) string {
	// TODO: implement
	return ""
}

type Point struct {
	Row, Col int
}

type ProblemInput struct {
	Grid          [][]rune
	RobotPosition Point

	MoveSequence []rune
}

func (pi *ProblemInput) SumBoxGPSValues() int {
	sum := 0
	for r, row := range pi.Grid {
		for c, val := range row {
			if val == 'O' {
				sum += 100*r + c
			}
		}
	}
	return sum
}

func (pi *ProblemInput) ApplyMoveSequence() {
	for _, move := range pi.MoveSequence {
		pi.ApplyMove(move)
	}
}

func (pi *ProblemInput) ApplyMove(moveType rune) {
	// < V ^ >
	moveVectors := map[rune]Point{
		'<': {0, -1},
		'v': {1, 0},
		'^': {-1, 0},
		'>': {0, 1},
	}
	if _, ok := moveVectors[moveType]; !ok {
		panic(fmt.Sprintf("Invalid move: %v", moveType))
	}
	move := moveVectors[moveType]
	if pi.CanMakeMove(pi.RobotPosition, move) {
		pi.MakeMove(pi.RobotPosition, move)
		pi.RobotPosition = Point{pi.RobotPosition.Row + move.Row, pi.RobotPosition.Col + move.Col}
	}
}

func (pi *ProblemInput) MakeMove(start Point, move Point) {
	// Move the robot pushing boxes if necessary
	if pi.Grid[start.Row+move.Row][start.Col+move.Col] == 'O' {
		nextStart := Point{start.Row + move.Row, start.Col + move.Col}
		pi.MakeMove(nextStart, move)
	}

	// Move point
	// start -> .
	// nextStart -> state value
	startValue := pi.Grid[start.Row][start.Col]
	pi.Grid[start.Row][start.Col] = '.'
	pi.Grid[start.Row+move.Row][start.Col+move.Col] = startValue
}

func (pi *ProblemInput) CanMakeMove(start Point, move Point) bool {
	// if we run into a wall '#" we can't move
	if pi.Grid[start.Row+move.Row][start.Col+move.Col] == '#' {
		return false
	}

	// if we run into a box 'O' we try to push the box
	// if the box runs into a wall we can't move
	// if the box runs into another box ....
	if pi.Grid[start.Row+move.Row][start.Col+move.Col] == 'O' {
		nextStart := Point{start.Row + move.Row, start.Col + move.Col}
		return pi.CanMakeMove(nextStart, move)
	}

	return true
}

func (pi *ProblemInput) PrettyPrint() {
	fmt.Printf("Robot at: %v\n", pi.RobotPosition)
	for _, row := range pi.Grid {
		fmt.Println(string(row))
	}
}

func parseInput(input string) *ProblemInput {
	pi := &ProblemInput{
		Grid: make([][]rune, 0),

		MoveSequence: []rune{},
	}

	gridMode := true
	for r, line := range common.ReadAsLines(input) {
		// empty line
		if len(strings.TrimSpace(line)) == 0 {
			gridMode = false
			continue
		}

		if gridMode {
			pi.Grid = append(pi.Grid, []rune(line))
			for c, el := range line {
				if el == '@' {
					pi.RobotPosition = Point{r, c}
				}
			}
		} else {
			for _, r := range line {
				pi.MoveSequence = append(pi.MoveSequence, r)
			}
		}
	}

	return pi
}
