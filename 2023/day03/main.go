package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"unicode"

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
				currentNumberOkay = currentNumberOkay || IsOkay(grid, rowIdx, colIdx, IsSpecialChar)
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
	grid := s.parseInput(input)

	isOnNumber := false
	currentNumber := ""
	gearMap := make(map[string][]int)
	gearKey := ""

	for rowIdx, line := range grid {
		for colIdx, char := range line {
			// Check if digit
			if unicode.IsDigit(char) {
				isOnNumber = true
				currentNumber += string(char)
				if gearKey == "" {
					gearKey = GetOkayKey(grid, rowIdx, colIdx, IsGear)
				}
			} else {
				// Were we on a number?
				if isOnNumber {
					if gearKey != "" {
						val := common.MustAtoi(currentNumber)
						if _, ok := gearMap[gearKey]; !ok {
							gearMap[gearKey] = make([]int, 0)
						}
						gearMap[gearKey] = append(gearMap[gearKey], val)
					}

					// reset
					isOnNumber = false
					gearKey = ""
					currentNumber = ""
				}
			}
		}
		if isOnNumber {
			if gearKey != "" {
				val := common.MustAtoi(currentNumber)
				if _, ok := gearMap[gearKey]; !ok {
					gearMap[gearKey] = make([]int, 0)
				}
				gearMap[gearKey] = append(gearMap[gearKey], val)
			}

			// reset
			isOnNumber = false
			currentNumber = ""
			gearKey = ""
		}
	}

	result := 0
	for _, values := range gearMap {
		if len(values) == 2 {
			result += values[0] * values[1]
		}
	}

	return strconv.Itoa(result)
}

//--------------------------------------------------------------------
// Helpers
//--------------------------------------------------------------------

func IsOkay(grid [][]rune, rowIdx, colIdx int, checkFn func(rune) bool) bool {
	for row := rowIdx - 1; row <= rowIdx+1; row++ {
		for col := colIdx - 1; col <= colIdx+1; col++ {
			// Is it in bounds?
			if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[row]) {
				continue
			}
			candidate := grid[row][col]
			if checkFn(candidate) {
				return true
			}
		}
	}
	return false
}

func GetOkayKey(grid [][]rune, rowIdx, colIdx int, checkFn func(rune) bool) string {
	for row := rowIdx - 1; row <= rowIdx+1; row++ {
		for col := colIdx - 1; col <= colIdx+1; col++ {
			// Is it in bounds?
			if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[row]) {
				continue
			}
			candidate := grid[row][col]
			if checkFn(candidate) {
				return fmt.Sprintf("%d,%d", row, col)
			}
		}
	}
	return ""
}

// = / + & % - # @ $

func IsGear(char rune) bool {
	return char == '*'
}

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
	input = strings.TrimRight(input, "\n")
	lines := strings.Split(string(input), "\n")

	result := make([][]rune, len(lines))
	for i, line := range lines {
		cleanLine := strings.TrimSpace(line)
		if cleanLine == "" {
			continue
		}
		result[i] = []rune(cleanLine)
	}
	return result
}
