package main

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/jhh3/aoc/common"
)

const (
	DAY  = 3
	YEAR = 2023
)

func main() {
	common.RunFromSolver(&solver{}, YEAR, DAY)
}

//--------------------------------------------------------------------
// Solution
//--------------------------------------------------------------------

type solver struct{}

func (s *solver) SolvePart1(input string) string {
	grid := s.parseInput(input)
	result := 0

	isOnNumber := false
	currentNumber := ""
	currentNumberOkay := false

	for rowIdx, line := range grid {
		for colIdx, char := range line {
			// Check if digit
			if unicode.IsDigit(char) {
				isOnNumber = true
				currentNumber += string(char)
				currentNumberOkay = currentNumberOkay || IsOkay(grid, rowIdx, colIdx)
			} else {
				// Were we on a number?
				if isOnNumber {
					if currentNumberOkay {
						result += common.MustAtoi(currentNumber)
					}

					// reset
					isOnNumber = false
					currentNumber = ""
					currentNumberOkay = false
				}
			}
		}
		if isOnNumber {
			if currentNumberOkay {
				result += common.MustAtoi(currentNumber)
			}

			// reset
			isOnNumber = false
			currentNumber = ""
			currentNumberOkay = false
		}
	}

	return strconv.Itoa(result)
}

func (s *solver) SolvePart2(input string) string {
	return ""
}

//--------------------------------------------------------------------
// Helpers
//--------------------------------------------------------------------

func IsOkay(grid [][]rune, rowIdx, colIdx int) bool {
	for row := rowIdx - 1; row <= rowIdx+1; row++ {
		for col := colIdx - 1; col <= colIdx+1; col++ {
			// Is it in bounds?
			if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[row]) {
				continue
			}
			candidate := grid[row][col]
			if IsSpecialChar(candidate) {
				return true
			}
		}
	}
	return false
}

// = / + & % - # @ $

var specialChars = []rune{'#', '~', '|', '*', '/', '%', '$', '=', '@', '&', '-', '+'}

func IsSpecialChar(char rune) bool {
	// check if in specialChars
	for _, specialChar := range specialChars {
		if char == specialChar {
			return true
		}
	}
	return false
}

//--------------------------------------------------------------------
// Parsing code
//--------------------------------------------------------------------

func (s *solver) parseInput(input string) [][]rune {
	lines := strings.Split(string(input), "\n")
	result := make([][]rune, len(lines)-1)
	for i, line := range lines {
		cleanLine := strings.TrimSpace(line)
		if cleanLine == "" {
			continue
		}
		result[i] = []rune(cleanLine)
	}
	return result
}
