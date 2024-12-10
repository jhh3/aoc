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

	wordCount := 0
	for r, row := range problemInput.WordSearch {
		for c, letter := range row {
			if letter == 'X' {
				wordCount += problemInput.CountAtLocation(r, c)
			}
		}
	}

	return strconv.Itoa(wordCount)
}

func (s *solver) SolvePart2(input string) string {
	problemInput := parseInput(input)

	xCount := 0
	for r, row := range problemInput.WordSearch {
		for c, letter := range row {
			if letter == 'A' {
				xCount += problemInput.CountX(r, c)
			}
		}
	}

	return strconv.Itoa(xCount)
}

type ProblemInput struct {
	WordSearch [][]rune
}

func (pi *ProblemInput) CountX(row, col int) int {
	masOnDiagonalLeft := (pi.E(row-1, col-1, 'M') && pi.E(row+1, col+1, 'S')) || (pi.E(row-1, col-1, 'S') && pi.E(row+1, col+1, 'M'))

	masOnDiagonalRight := (pi.E(row-1, col+1, 'M') && pi.E(row+1, col-1, 'S')) || (pi.E(row-1, col+1, 'S') && pi.E(row+1, col-1, 'M'))

	if masOnDiagonalLeft && masOnDiagonalRight {
		return 1
	}

	return 0
}

func (pi *ProblemInput) CountAtLocation(row, col int) int {
	return pi.CheckVertical(row, col) + pi.CheckVerticalBackwards(row, col) + pi.CheckDiagonalDR(row, col) + pi.CheckDiagonalDL(row, col) + pi.CheckHorizontal(row, col) + pi.CheckHorizontalBackwards(row, col) + pi.CheckDiagonalUR(row, col) + pi.CheckDiagonalUL(row, col)
}

func (pi *ProblemInput) CheckVertical(row, col int) int {
	if pi.E(row, col, 'X') && pi.E(row+1, col, 'M') && pi.E(row+2, col, 'A') && pi.E(row+3, col, 'S') {
		return 1
	}

	return 0
}

func (pi *ProblemInput) CheckVerticalBackwards(row, col int) int {
	if pi.E(row, col, 'X') && pi.E(row-1, col, 'M') && pi.E(row-2, col, 'A') && pi.E(row-3, col, 'S') {
		return 1
	}

	return 0
}

func (pi *ProblemInput) CheckDiagonalUR(row, col int) int {
	if pi.E(row, col, 'X') && pi.E(row-1, col+1, 'M') && pi.E(row-2, col+2, 'A') && pi.E(row-3, col+3, 'S') {
		return 1
	}

	return 0
}

func (pi *ProblemInput) CheckDiagonalUL(row, col int) int {
	if pi.E(row, col, 'X') && pi.E(row-1, col-1, 'M') && pi.E(row-2, col-2, 'A') && pi.E(row-3, col-3, 'S') {
		return 1
	}

	return 0
}

func (pi *ProblemInput) CheckDiagonalDR(row, col int) int {
	if pi.E(row, col, 'X') && pi.E(row+1, col+1, 'M') && pi.E(row+2, col+2, 'A') && pi.E(row+3, col+3, 'S') {
		return 1
	}

	return 0
}

func (pi *ProblemInput) CheckDiagonalDL(row, col int) int {
	if pi.E(row, col, 'X') && pi.E(row+1, col-1, 'M') && pi.E(row+2, col-2, 'A') && pi.E(row+3, col-3, 'S') {
		return 1
	}

	return 0
}

func (pi *ProblemInput) CheckHorizontal(row, col int) int {
	if pi.E(row, col, 'X') && pi.E(row, col+1, 'M') && pi.E(row, col+2, 'A') && pi.E(row, col+3, 'S') {
		return 1
	}

	return 0
}

func (pi *ProblemInput) CheckHorizontalBackwards(row, col int) int {
	if pi.E(row, col, 'X') && pi.E(row, col-1, 'M') && pi.E(row, col-2, 'A') && pi.E(row, col-3, 'S') {
		return 1
	}

	return 0
}

func (pi *ProblemInput) E(row, col int, letter rune) bool {
	// check inbounds
	if row < 0 || row >= len(pi.WordSearch) {
		return false
	}
	if col < 0 || col >= len(pi.WordSearch[row]) {
		return false
	}

	return pi.WordSearch[row][col] == letter
}

func parseInput(input string) ProblemInput {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	result := make([][]rune, len(lines))
	for r, row := range lines {
		result[r] = make([]rune, len(row))
		for c, letter := range row {
			result[r][c] = letter
		}
	}

	return ProblemInput{
		WordSearch: result,
	}
}
