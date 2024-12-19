package main

import (
	_ "embed"
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
	pi := parseInput(input)
	return common.Itoa(pi.CountPossibleTowels())
}

func (s *solver) SolvePart2(input string) string {
	pi := parseInput(input)
	return common.Itoa(pi.NumPossibleWaysToMakeTowels())
}

type Pattern = []rune

type Towel struct {
	Pattern Pattern
	cache   map[string]int
}

func NewTowel(pattern Pattern) *Towel {
	return &Towel{
		Pattern: pattern,
		cache:   make(map[string]int),
	}
}

func Equals(a Pattern, b Pattern) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func (t *Towel) IsPossible(availablePatterns []Pattern) bool {
	// If the towel pattern is empty, it's technically possible
	if len(t.Pattern) == 0 {
		return true
	}

	// Try each available pattern
	for _, pattern := range availablePatterns {
		// Skip patterns that are longer than our target
		if len(pattern) > len(t.Pattern) {
			continue
		}

		// Check if this pattern matches at the start of our towel pattern
		matches := true
		for i := 0; i < len(pattern); i++ {
			if pattern[i] != t.Pattern[i] {
				matches = false
				break
			}
		}

		// If we found a match at the start
		if matches {
			// Create a new towel with the remaining pattern
			remainingPattern := t.Pattern[len(pattern):]
			remainingTowel := &Towel{Pattern: remainingPattern}

			// Recursively check if the remaining pattern is possible
			if remainingTowel.IsPossible(availablePatterns) {
				return true
			}
		}
	}

	return false
}

func (t *Towel) NumPossibleWaysToMakeTowel(availablePatterns []Pattern) int {
	// Check cache first
	key := string(t.Pattern)
	if count, exists := t.cache[key]; exists {
		return count
	}

	// If the towel pattern is empty, it's technically possible
	if len(t.Pattern) == 0 {
		return 1
	}

	// Try each available pattern
	count := 0
	for _, pattern := range availablePatterns {
		// Skip patterns that are longer than our target
		if len(pattern) > len(t.Pattern) {
			continue
		}

		// Check if this pattern matches at the start of our towel pattern
		matches := true
		for i := 0; i < len(pattern); i++ {
			if pattern[i] != t.Pattern[i] {
				matches = false
				break
			}
		}

		// If we found a match at the start
		if matches {
			// Create a new towel with the remaining pattern
			remainingPattern := t.Pattern[len(pattern):]
			remainingTowel := NewTowel(remainingPattern)
			remainingTowel.cache = t.cache // Share the same cache

			count += remainingTowel.NumPossibleWaysToMakeTowel(availablePatterns)
		}
	}

	// Cache the result before returning
	t.cache[key] = count
	return count
}

type ProblemInput struct {
	AvailablePatterns []Pattern

	Towels []Towel
}

func (pi *ProblemInput) CountPossibleTowels() int {
	count := 0
	for _, towel := range pi.Towels {
		if towel.IsPossible(pi.AvailablePatterns) {
			count++
		}
	}
	return count
}

func (pi *ProblemInput) NumPossibleWaysToMakeTowels() int {
	count := 0
	for _, towel := range pi.Towels {
		count += towel.NumPossibleWaysToMakeTowel(pi.AvailablePatterns)
	}
	return count
}

func parseInput(input string) *ProblemInput {
	lines := common.ReadAsLines(input)

	pi := &ProblemInput{
		AvailablePatterns: make([]Pattern, 0),
		Towels:            make([]Towel, 0),
	}

	for i, line := range lines {
		if len(line) == 0 {
			continue
		}

		if i == 0 {
			patternStrs := strings.Split(line, ",")
			for _, patternStr := range patternStrs {
				cleanPatternStr := strings.TrimSpace(patternStr)
				pi.AvailablePatterns = append(pi.AvailablePatterns, []rune(cleanPatternStr))
			}
			continue
		}

		cleanPatternStr := strings.TrimSpace(line)
		pi.Towels = append(pi.Towels, *NewTowel([]rune(cleanPatternStr)))
	}

	return pi
}
