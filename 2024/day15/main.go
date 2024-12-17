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
	problemInput := parseInput(input)
	pi2 := FromProblemInput(problemInput)
	pi2.PrettyPrint()
	pi2.ApplyMoveSequence()
	println()
	println()
	pi2.PrettyPrint()
	return strconv.Itoa(pi2.SumBoxGPSValues())
}

type Point struct {
	Row, Col int
}

type ProblemInput struct {
	Grid          [][]rune
	RobotPosition Point

	MoveSequence []rune
}

type ProblemInputPart2 struct {
	Grid          [][]rune
	RobotPosition Point

	MoveSequence []rune
}

func (pi *ProblemInputPart2) SumBoxGPSValues() int {
	sum := 0
	for r, row := range pi.Grid {
		for c, val := range row {
			if val == '[' {
				sum += 100*r + c
			}
		}
	}
	return sum
}

func (pi *ProblemInputPart2) ApplyMoveSequence() {
	for _, move := range pi.MoveSequence {
		pi.ApplyMove(move)
	}
}

func (pi *ProblemInputPart2) ApplyMove(moveType rune) {
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
	if pi.CanMakeMove(pi.RobotPosition, move, moveType) {
		pi.MakeMove(pi.RobotPosition, move, moveType)
		pi.RobotPosition = Point{pi.RobotPosition.Row + move.Row, pi.RobotPosition.Col + move.Col}
	}
}

func (pi *ProblemInputPart2) MakeMove(start Point, move Point, moveType rune) {
	isHorizontal := moveType == '<' || moveType == '>'

	nextIsLeftSide := pi.Grid[start.Row+move.Row][start.Col+move.Col] == '['
	nextIsRightSide := pi.Grid[start.Row+move.Row][start.Col+move.Col] == ']'

	if isHorizontal {
		// easy

		// Move the robot pushing boxes if necessary
		if nextIsLeftSide || nextIsRightSide {
			nextStart := Point{start.Row + move.Row, start.Col + move.Col}
			pi.MakeMove(nextStart, move, moveType)
		}
	} else {
		// need to move both sides
		if nextIsLeftSide {
			nextLeftStart := Point{start.Row + move.Row, start.Col + move.Col}
			nextRightStart := Point{start.Row + move.Row, start.Col + move.Col + 1}
			pi.MakeMove(nextLeftStart, move, moveType)
			pi.MakeMove(nextRightStart, move, moveType)
		} else if nextIsRightSide {
			nextRightStart := Point{start.Row + move.Row, start.Col + move.Col}
			nextLeftStart := Point{start.Row + move.Row, start.Col + move.Col - 1}
			pi.MakeMove(nextRightStart, move, moveType)
			pi.MakeMove(nextLeftStart, move, moveType)
		}
	}

	// Move point
	// start -> .
	// nextStart -> state value
	startValue := pi.Grid[start.Row][start.Col]
	pi.Grid[start.Row][start.Col] = '.'
	pi.Grid[start.Row+move.Row][start.Col+move.Col] = startValue
}

func (pi *ProblemInputPart2) CanMakeMove(start Point, move Point, moveType rune) bool {
	// if we run into a wall '#" we can't move
	if pi.Grid[start.Row+move.Row][start.Col+move.Col] == '#' {
		return false
	}

	nextIsLeftSide := pi.Grid[start.Row+move.Row][start.Col+move.Col] == '['
	nextIsRightSide := pi.Grid[start.Row+move.Row][start.Col+move.Col] == ']'

	isHorizontal := moveType == '<' || moveType == '>'
	isVertical := moveType == '^' || moveType == 'v'

	if nextIsLeftSide || nextIsRightSide {
		if isHorizontal {
			nextStart := Point{start.Row + move.Row, start.Col + move.Col}
			return pi.CanMakeMove(nextStart, move, moveType)
		}
		if isVertical {
			// need to try moving both sides
			if nextIsLeftSide {
				nextLeftStart := Point{start.Row + move.Row, start.Col + move.Col}
				nextRightStart := Point{start.Row + move.Row, start.Col + move.Col + 1}

				return pi.CanMakeMove(nextLeftStart, move, moveType) && pi.CanMakeMove(nextRightStart, move, moveType)
			}
			if nextIsRightSide {
				nextRightStart := Point{start.Row + move.Row, start.Col + move.Col}
				nextLeftStart := Point{start.Row + move.Row, start.Col + move.Col - 1}
				return pi.CanMakeMove(nextRightStart, move, moveType) && pi.CanMakeMove(nextLeftStart, move, moveType)
			}
		}

	}

	return true
}

func (pi *ProblemInputPart2) PrettyPrint() {
	fmt.Printf("Robot at: %v\n", pi.RobotPosition)
	for _, row := range pi.Grid {
		fmt.Println(string(row))
	}
}

func FromProblemInput(pi *ProblemInput) *ProblemInputPart2 {
	// The grid is twice as wide
	pi2 := &ProblemInputPart2{
		Grid: make([][]rune, 0),

		MoveSequence: pi.MoveSequence,
	}

	for r, row := range pi.Grid {
		pi2.Grid = append(pi2.Grid, make([]rune, 0))
		for _, val := range row {
			if val == '#' {
				pi2.Grid[r] = append(pi2.Grid[r], '#', '#')
			} else if val == 'O' {
				pi2.Grid[r] = append(pi2.Grid[r], '[', ']')
			} else if val == '.' {
				pi2.Grid[r] = append(pi2.Grid[r], '.', '.')
			} else if val == '@' {
				pi2.Grid[r] = append(pi2.Grid[r], val, '.')
			}
		}
	}

	pi2.RobotPosition = Point{pi.RobotPosition.Row, pi.RobotPosition.Col * 2}

	return pi2
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
