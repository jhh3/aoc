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
	// fmt.Printf("CurRow: %d, CurCol: %d, Direction: %s\n", problemInput.CurRow, problemInput.CurCol, problemInput.DirectionString())
	visitedCount := problemInput.VisitGrid()
	return strconv.Itoa(visitedCount)
}

func (s *solver) SolvePart2(input string) string {
	problemInput := parseInput(input)
	waysToCreateALoop := problemInput.CountWaysToCreaeALoop()
	return strconv.Itoa(waysToCreateALoop)
}

type ProblemInput struct {
	Grid [][]rune

	VisitedCount int
	Direction    int // 0 = up, 1 = right, 2 = down, 3 = left
	CurRow       int
	CurCol       int
}

func (pi *ProblemInput) VisitGrid() int {
	// Get the next location the guard would try to move to
	nextRow, nextCol := pi.NextRowCol()

	// Mark current location as visited
	currentElement := pi.Grid[pi.CurRow][pi.CurCol]
	if currentElement != 'X' {
		pi.Grid[pi.CurRow][pi.CurCol] = 'X'
		pi.VisitedCount++
	}

	// Inbounds check
	if nextRow < 0 || nextRow >= len(pi.Grid) || nextCol < 0 || nextCol >= len(pi.Grid[0]) {
		// If moving out of bounds, return the number of visited locations
		return pi.VisitedCount
	}

	// Obstacle check
	if nextElement := pi.Grid[nextRow][nextCol]; nextElement == '#' {
		// if obstacle, turn right
		pi.Direction = (pi.Direction + 1) % 4
		nextRow, nextCol = pi.NextRowCol()
	}

	// Move forward
	pi.CurRow = nextRow
	pi.CurCol = nextCol

	// Recurse
	return pi.VisitGrid()
}

func (pi *ProblemInput) CountWaysToCreaeALoop() int {
	count := 0

	for r, row := range pi.Grid {
		for c, element := range row {
			// check if we can place an obstacle here
			if element == '.' {
				copyProblemInput := ProblemInput{
					Grid:         make([][]rune, len(pi.Grid)),
					VisitedCount: 0,
					Direction:    pi.Direction,
					CurRow:       pi.CurRow,
					CurCol:       pi.CurCol,
				}
				// copy the grid
				for i := range pi.Grid {
					copyProblemInput.Grid[i] = make([]rune, len(pi.Grid[i]))
					copy(copyProblemInput.Grid[i], pi.Grid[i])
				}

				copyProblemInput.Grid[r][c] = '#'

				threshold := len(pi.Grid) * len(pi.Grid[0]) * 2
				if copyProblemInput.LoopExists(pi.CurRow, pi.CurCol, pi.Direction, 0, threshold) {
					count++
				}
			}
		}
	}

	return count
}

func (pi *ProblemInput) LoopExists(startRow, startCol, direction, num, threshold int) bool {
	// Check if move forward leads to loop
	if num > threshold {
		return true
	}

	// Get the next location the guard would try to move to
	nextRow, nextCol := pi.NextRowCol()

	// Inbounds check
	if nextRow < 0 || nextRow >= len(pi.Grid) || nextCol < 0 || nextCol >= len(pi.Grid[0]) {
		// If moving out of bounds, no loop
		return false
	}

	// Obstacle check
	if nextElement := pi.Grid[nextRow][nextCol]; nextElement == '#' {
		// if obstacle, turn right
		pi.Direction = (pi.Direction + 1) % 4
	} else {
		// Move forward
		pi.CurRow = nextRow
		pi.CurCol = nextCol
	}

	// inLocationWeStarted := pi.CurRow == startRow && pi.CurCol == startCol
	// if inLocationWeStarted {
	// 	fmt.Printf("\t >> Loop exists if we return to start location\n")
	// }

	// Recurse
	return pi.LoopExists(startRow, startCol, direction, num+1, threshold)
}

func (pi *ProblemInput) NextRowCol() (int, int) {
	switch pi.Direction {
	case 0:
		return pi.CurRow - 1, pi.CurCol
	case 1:
		return pi.CurRow, pi.CurCol + 1
	case 2:
		return pi.CurRow + 1, pi.CurCol
	case 3:
		return pi.CurRow, pi.CurCol - 1
	}
	return -1, -1
}

func (pi *ProblemInput) DirectionString() string {
	switch pi.Direction {
	case 0:
		return "up"
	case 1:
		return "right"
	case 2:
		return "down"
	case 3:
		return "left"
	}
	return "unknown"
}

func parseInput(input string) ProblemInput {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	result := ProblemInput{
		Grid:         make([][]rune, len(lines)),
		VisitedCount: 0,
	}
	for i, row := range lines {
		result.Grid[i] = make([]rune, len(row))
		for j, element := range row {
			result.Grid[i][j] = element
			// check for start
			// if element in [^, >, v, <]
			if element == '^' || element == '>' || element == 'v' || element == '<' {
				result.CurRow = i
				result.CurCol = j
				switch element {
				case '^':
					result.Direction = 0
				case '>':
					result.Direction = 1
				case 'v':
					result.Direction = 2
				case '<':
					result.Direction = 3

				}
			}

		}
	}

	return result
}
